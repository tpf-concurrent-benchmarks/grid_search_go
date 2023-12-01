package main

import (
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/nats-io/nats.go"
	"log"
	"manager/src/interval"
	"manager/src/message_processor"
	"manager/src/utils"
	common "shared"
	"shared/config"
	"sync"
	"time"
)

func main() {
	managerConfig := config.GetConfig()
	connString := config.CreateConnectionAddress(managerConfig.Host, managerConfig.Port)

	metricsAddr := config.CreateMetricAddress(managerConfig.Metrics.Host, managerConfig.Metrics.Port)
	statsdClient, err := statsd.NewClient(metricsAddr, "manager")

	startTime := time.Now()

	natsConnection, err := nats.Connect(connString)
	encodedConn, _ := nats.NewEncodedConn(natsConnection, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}

	defer encodedConn.Close()

	data := config.GetDataFromJson("./src/resources/data.json")
	intervals := interval.NewIntervalFromArray(data.Data)
	partition := interval.NewPartition(intervals, len(intervals), data.MaxItemsPerBatch)
	workSent := partition.GetNPartitions()
	m := &sync.Mutex{}
	c := sync.NewCond(m)
	c.L.Lock()

	messageProcessor := message_processor.NewMessageProcessor(data.Agg)

	subscribeForResults(encodedConn, managerConfig.Queues.Input, messageProcessor, workSent, c)

	sendWork(partition, data.Agg, encodedConn, managerConfig.Queues.Output)

	c.Wait()
	sendEndMessage(natsConnection)
	c.L.Unlock()

	messageProcessor.SaveResults()
	endTime := time.Now()
	elapseTime := endTime.Sub(startTime).Milliseconds()

	err = statsdClient.Gauge("completion_time", elapseTime, 1.0)
	if err != nil {
		log.Fatalf("Error sending metric to statsd: %s", err)
	}

}

func subscribeForResults(encodedConn *nats.EncodedConn, inputQueue string, messageProcessor *message_processor.MessageProcessor, workSent uint64, c *sync.Cond) {
	resultsReceived := uint64(0)
	_, _ = encodedConn.Subscribe(inputQueue, func(msg *nats.Msg) {
		message := utils.ParseMessage(msg)
		messageProcessor.ProcessMessage(message)
		resultsReceived++
		if resultsReceived == workSent {
			c.L.Unlock()
			c.Signal()
		}
	})
}

func sendEndMessage(conn *nats.Conn) {
	err := conn.Publish(common.EndWorkQueue, []byte(common.EndWorkMessage))
	if err != nil {
		log.Fatalf("Error publishing to queue: %s", err)
	}
}

func sendWork(partition *interval.Partition, aggregation string, encodedConn *nats.EncodedConn, outputQueue string) {
	for partition.Available() {
		partitionData := partition.Next()
		workMessage := utils.CreateWorkMessageFrom(partitionData, aggregation)
		err := encodedConn.Publish(outputQueue, workMessage)
		if err != nil {
			log.Fatalf("Error publishing to queue: %s", err)
		}
	}
}

package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"manager/src/interval"
	"manager/src/message_processor"
	"manager/src/utils"
	common "shared"
	"shared/config"
	"time"
)

func main() {
	managerConfig := config.GetConfig()
	connString := config.CreateConnectionAddress(managerConfig.Host, managerConfig.Port)

	metricsAddr := config.CreateMetricAddress(managerConfig.Metrics.Host, managerConfig.Metrics.Port)
	statsdClient := utils.CreateStatsClient(metricsAddr, utils.GetNodeID())

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
	ch := make(chan bool)

	messageProcessor := message_processor.NewMessageProcessor(data.Agg)

	// For sync purposes
	time.Sleep(5 * time.Second)

	subscribeForResults(encodedConn, managerConfig.Queues.Input, messageProcessor, workSent, ch)

	sendWork(partition, data.Agg, encodedConn, managerConfig.Queues.Output)

	if <-ch {
		log.Println("All results received, sending end message")
		sendEndMessage(natsConnection)
		messageProcessor.SaveResults()
		endTime := time.Now()
		elapseTime := endTime.Sub(startTime).Milliseconds()
		err = statsdClient.Gauge("completion_time", elapseTime, 1.0)
		if err != nil {
			log.Fatalf("Error sending metric to statsd: %s", err)
		}
		close(ch)
	}

}

func subscribeForResults(encodedConn *nats.EncodedConn, inputQueue string, messageProcessor *message_processor.MessageProcessor, workSent uint64, ch chan bool) {
	resultsReceived := uint64(0)
	log.Println("Subscribing for results, work sent: ", workSent)
	_, _ = encodedConn.Subscribe(inputQueue, func(msg *nats.Msg) {
		message := utils.ParseMessage(msg.Data)
		messageProcessor.ProcessMessage(message)
		resultsReceived++
		if resultsReceived == workSent {
			ch <- true
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

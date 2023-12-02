package main

import (
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	common "shared"
	"shared/config"
	"shared/dto"
	"time"
	"worker/src/grid_search"
)

func main() {
	workerConfig := config.GetConfig()
	connString := config.CreateConnectionAddress(workerConfig.Host, workerConfig.Port)
	natsConn, err := nats.Connect(connString)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}
	encodedConn, _ := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)

	metricsAddr := config.CreateMetricAddress(workerConfig.Metrics.Host, workerConfig.Metrics.Port)
	statsdClient := CreateStatsClient(metricsAddr, GetNodeID())
	ch := make(chan bool)

	waitForEnd(encodedConn, ch)

	subscribeForWork(encodedConn, workerConfig, statsdClient)

	if <-ch {
		encodedConn.Close()
		close(ch)
	}
}

func subscribeForWork(encodedConn *nats.EncodedConn, workerConfig config.Config, statsdClient statsd.Statter) {
	_, err := encodedConn.QueueSubscribe(workerConfig.Queues.Input, "workers_group", func(message *dto.WorkMessage) {
		startTime := time.Now()
		aggregation, gridSearch := gridSearchFrom(message)
		gridSearch.Search(grid_search.GriewankFunc)
		sendResult(aggregation, encodedConn, workerConfig.Queues.Output, gridSearch)

		endTime := time.Now()
		elapseTime := endTime.Sub(startTime).Milliseconds()

		err := statsdClient.Timing("work_time", elapseTime, 1.0)
		if err != nil {
			log.Fatalf("Error sending timing metric to statsd: %s", err)
		}
		err = statsdClient.Inc("results_produced", 1, 1.0)
		if err != nil {
			log.Fatalf("Error sending increment metric to statsd: %s", err)
		}
	})
	if err != nil {
		log.Fatalf("Error subscribing to queue: %s", err)
	}
}

func CreateStatsClient(metricsAddr, prefix string) statsd.Statter {
	clientConfig := &statsd.ClientConfig{
		Address: metricsAddr,
		Prefix:  prefix,
	}

	statsdClient, err := statsd.NewClientWithConfig(clientConfig)
	if err != nil {
		log.Fatalf("Error creating statsd client: %s", err)
	}
	return statsdClient
}

func sendResult(aggregation string, encodedConn *nats.EncodedConn, outputQueue string, gridSearch *grid_search.GridSearch) {
	switch aggregation {
	case "MAX":
		_ = encodedConn.Publish(outputQueue, dto.MaxResultsDTO{Value: gridSearch.GetResult(), Parameters: gridSearch.GetInput()})
	case "MIN":
		_ = encodedConn.Publish(outputQueue, dto.MinResultsDTO{Value: gridSearch.GetResult(), Parameters: gridSearch.GetInput()})
	case "AVG":
		_ = encodedConn.Publish(outputQueue, dto.AvgResultsDTO{Value: gridSearch.GetResult(), ParametersAmount: gridSearch.GetTotalInputs()})
	default:
		log.Fatalf("Invalid aggregation type: %s", aggregation)
	}
}

func gridSearchFrom(message *dto.WorkMessage) (string, *grid_search.GridSearch) {
	aggregation := message.Agg
	parameters := message.Data
	start, end, step := [grid_search.Size]float64{}, [grid_search.Size]float64{}, [grid_search.Size]float64{}
	for i := 0; i < len(parameters); i++ {
		start[i] = parameters[i][0]
		end[i] = parameters[i][1]
		step[i] = parameters[i][2]
	}
	params := grid_search.NewParams(start, end, step)
	gridSearch := grid_search.NewGridSearch(params, aggregation)
	return aggregation, gridSearch
}

func waitForEnd(encodedConn *nats.EncodedConn, ch chan bool) {
	_, err := encodedConn.Subscribe(common.EndWorkQueue, func(msg *nats.Msg) {
		message := string(msg.Data)
		if message == common.EndWorkMessage {
			log.Println("Received end message")
			ch <- true
		}
	})
	if err != nil {
		log.Fatalf("Error subscribing to end queue: %s", err)
	}
}

func GetNodeID() string {
	if os.Getenv("LOCAL") == "" {
		return os.Getenv("NODE_ID")
	}
	return "manager"
}

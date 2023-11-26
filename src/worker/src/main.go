package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"shared/config"
	"shared/dto"
	"worker/src/grid_search"
)

func main() {
	workerConfig := config.GetConfig()
	connString := config.CreateConnectionString(workerConfig.Host, workerConfig.Port)
	natsConn, err := nats.Connect(connString)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}
	encodedConn, _ := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)
	defer encodedConn.Close()

	_, err = encodedConn.QueueSubscribe(workerConfig.Queues.Input, "workers_group", func(message *dto.WorkMessage) {
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
		gridSearch.Search(grid_search.GriewankFunc)

		switch aggregation {
		case "MAX":
			_ = encodedConn.Publish(workerConfig.Queues.Output, dto.MaxResultsDTO{Value: gridSearch.GetResult(), Parameters: gridSearch.GetInput()})
		case "MIN":
			_ = encodedConn.Publish(workerConfig.Queues.Output, dto.MinResultsDTO{Value: gridSearch.GetResult(), Parameters: gridSearch.GetInput()})
		case "AVG":
			_ = encodedConn.Publish(workerConfig.Queues.Output, dto.AvgResultsDTO{Value: gridSearch.GetResult(), ParametersAmount: gridSearch.GetTotalInputs()})
		default:
			log.Fatalf("Invalid aggregation type: %s", aggregation)
		}
	})
	if err != nil {
		log.Fatalf("Error subscribing to queue: %s", err)
	}
	select {}
}

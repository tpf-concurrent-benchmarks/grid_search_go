package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"manager/src/interval"
	"manager/src/utils"
	"math"
	"shared/config"
	"shared/dto"
)

func main() {
	managerConfig := config.GetConfig()
	connString := config.CreateConnectionString(managerConfig.Host, managerConfig.Port)

	natsConnection, err := nats.Connect(connString)
	encodedConn, _ := nats.NewEncodedConn(natsConnection, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}
	defer encodedConn.Close()

	data := config.GetDataFromJson("./src/resources/data.json")
	intervals := interval.NewIntervalFromArray(data.Data)
	partition := interval.NewPartition(intervals, len(intervals), data.MaxItemsPerBatch)

	for partition.Available() {
		partitionData := partition.Next()
		workMessage := utils.CreateWorkMessageFrom(partitionData, data.Agg)
		err := encodedConn.Publish(managerConfig.Queues.Output, workMessage)
		if err != nil {
			log.Fatalf("Error publishing to queue: %s", err)
		}
	}

	if data.Agg == "MAX" {
		currentMaxValue := math.Inf(-1)
		var currentMaxParameters [3]float64
		subscribe, _ := encodedConn.Subscribe(managerConfig.Queues.Input, func(message *dto.MaxResultsDTO) {
			if message.Value > currentMaxValue {
				currentMaxValue = message.Value
				currentMaxParameters = message.Parameters
			}
		})
		// TODO this value should be the number of expected results (it equals to the number of work messages sent)
		_ = subscribe.AutoUnsubscribe(100)
		log.Println("Max value:", currentMaxValue, "Max parameters:", currentMaxParameters)
	} else if data.Agg == "MIN" {
		currentMin := math.Inf(1)
		var currentMinParameters [3]float64
		subscribe, _ := encodedConn.Subscribe(managerConfig.Queues.Input, func(message *dto.MinResultsDTO) {
			if message.Value < currentMin {
				currentMin = message.Value
				currentMinParameters = message.Parameters
			}
		})
		// TODO this value should be the number of expected results (it equals to the number of work messages sent)
		_ = subscribe.AutoUnsubscribe(100)
		log.Println("Min value:", currentMin, "Min parameters:", currentMinParameters)
	} else if data.Agg == "AVG" {
		currentAverage := 1.0
		totalParameters := 1.0
		subscribe, _ := encodedConn.Subscribe(managerConfig.Queues.Input, func(message *dto.AvgResultsDTO) {
			currentAverage += (currentAverage*totalParameters + message.Value*float64(message.ParametersAmount)) / (totalParameters + float64(message.ParametersAmount))
			totalParameters += float64(message.ParametersAmount)
		})
		// TODO this value should be the number of expected results (it equals to the number of work messages sent)
		_ = subscribe.AutoUnsubscribe(100)
		log.Println("Average value: ", currentAverage, "Total parameters: ", totalParameters)
	}

}

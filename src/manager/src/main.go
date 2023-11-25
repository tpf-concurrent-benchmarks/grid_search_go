package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"manager/src/interval"
	"manager/src/utils"
	"shared/config"
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
	intervals := interval.IntervalsFromJson(data)
	partition := interval.NewPartition(intervals, len(intervals), data.MaxItemsPerBatch)

	for partition.Available() {
		partitionData := partition.Next()
		workMessage := utils.CreateWorkMessageFrom(partitionData, data.Agg)
		err := encodedConn.Publish(managerConfig.Queues.Output, workMessage)
		if err != nil {
			log.Fatalf("Error publishing to queue: %s", err)
		}
	}

}

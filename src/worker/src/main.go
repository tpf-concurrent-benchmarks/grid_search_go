package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"shared/config"
	"worker/src/processing"
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

	_, err = encodedConn.QueueSubscribe(workerConfig.Queues.Input, "workers_group", processing.ProcessMessage)
	if err != nil {
		log.Fatalf("Error subscribing to queue: %s", err)
	}
	select {}
}

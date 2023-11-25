package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"shared/config"
)

func main() {
	workerConfig := config.GetConfig()
	connString := config.CreateConnectionString(workerConfig.Host, workerConfig.Port)
	nc, _ := nats.Connect(connString)
	defer nc.Close()

	_, err := nc.QueueSubscribe(workerConfig.Queues.Input, "workers_group", func(message *nats.Msg) {
		println(string(message.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to queue: %s", err)
	}
	select {}
}

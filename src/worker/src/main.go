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

	_, err := nc.QueueSubscribe("foo", "job_workers", func(message *nats.Msg) {
		println(string(message.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to queue: %s", err)
	}
	select {}
}

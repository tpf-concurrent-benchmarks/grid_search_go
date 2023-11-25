package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	nc, _ := nats.Connect("nats://localhost:4222")

	_, err := nc.QueueSubscribe("foo", "job_workers", func(message *nats.Msg) {
		// Print the message
		println(string(message.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to queue: %s", err)
	}
	nc.Close()
}

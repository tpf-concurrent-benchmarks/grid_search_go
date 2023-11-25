package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	nc.Close()
}

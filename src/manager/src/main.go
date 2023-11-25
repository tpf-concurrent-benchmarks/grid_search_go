package main

import (
	"github.com/nats-io/nats.go"
	"log"
	_ "manager/src/interval"
	"shared/config"
)

func main() {
	_ = config.GetConfig()
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}

	nc.Publish("foo", []byte("Hello World"))

	nc.Close()
}

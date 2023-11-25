package main

import (
	"github.com/nats-io/nats.go"
	"log"
	_ "manager/src/interval"
	"shared/config"
)

func main() {
	managerConfig := config.GetConfig()
	connString := config.CreateConnectionString(managerConfig.Host, managerConfig.Port)
	nc, err := nats.Connect(connString)
	data := config.GetDataFromJson("./src/resources/data.json")
	defer nc.Close()

	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}

	nc.Publish(managerConfig.Queues.Output, []byte("Hello World"))

}

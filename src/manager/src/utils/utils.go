package utils

import (
	"encoding/json"
	"github.com/cactus/go-statsd-client/v5/statsd"
	"log"
	"manager/src/interval"
	"os"
	"shared/dto"
)

func CreateWorkMessageFrom(intervals []interval.Interval, aggregation string) dto.WorkMessage {
	intervalList := make([][3]float64, len(intervals))
	for i, auxInterval := range intervals {
		intervalList[i] = auxInterval.GetInterval()
	}
	message := dto.WorkMessage{
		Data: intervalList,
		Agg:  aggregation,
	}
	return message
}

func GetNodeID() string {
	if os.Getenv("LOCAL") == "" {
		return os.Getenv("NODE_ID")
	}
	return "manager"
}

func ParseMessage(message []byte) map[string]json.RawMessage {
	var messageMap map[string]json.RawMessage
	err := json.Unmarshal(message, &messageMap)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return messageMap
}

func CreateStatsClient(metricsAddr, prefix string) statsd.Statter {
	clientConfig := &statsd.ClientConfig{
		Address: metricsAddr,
		Prefix:  prefix,
	}

	statsdClient, err := statsd.NewClientWithConfig(clientConfig)
	if err != nil {
		log.Fatalf("Error creating statsd client: %s", err)
	}
	return statsdClient
}

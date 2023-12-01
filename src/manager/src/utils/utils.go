package utils

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"manager/src/interval"
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

func ParseMessage(message *nats.Msg) map[string]interface{} {
	messageMap := make(map[string]interface{})
	err := json.Unmarshal(message.Data, &messageMap)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return messageMap
}

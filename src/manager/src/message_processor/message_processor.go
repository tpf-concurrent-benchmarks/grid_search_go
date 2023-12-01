package message_processor

import (
	"encoding/json"
	"log"
	"os"
)

type MessageProcessor struct {
	value           float64
	parameters      [3]float64
	currentAverage  float64
	totalParameters float64
	aggregation     string
}

func NewMessageProcessor(aggregation string) *MessageProcessor {
	return &MessageProcessor{
		aggregation:     aggregation,
		currentAverage:  0.0,
		totalParameters: 0.0,
	}
}

func (mp *MessageProcessor) ProcessMessage(message map[string]interface{}) {
	value := message["value"].(float64)
	parameters := message["parameters"].([3]float64)

	if mp.aggregation == "MAX" {
		if value > mp.value {
			mp.value = value
			mp.parameters = parameters
		}
	} else if mp.aggregation == "MIN" {
		if value < mp.value {
			mp.value = value
			mp.parameters = parameters
		}
	} else {
		paramsAmount := message["paramsAmount"].(uint64)
		mp.currentAverage = (mp.currentAverage*mp.totalParameters + message["value"].(float64)*float64(paramsAmount)) / (mp.totalParameters + float64(paramsAmount))
		mp.totalParameters += float64(paramsAmount)
	}
}

func (mp *MessageProcessor) SaveResults() {
	resultsJSON := mp.resultsToJson()
	log.Println("Saving results: ", string(resultsJSON))

	err := os.WriteFile("results.json", resultsJSON, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func (mp *MessageProcessor) resultsToJson() []byte {
	results := map[string]interface{}{
		"value":      mp.value,
		"parameters": mp.parameters,
		"agg":        mp.aggregation,
	}

	resultsJSON, _ := json.MarshalIndent(results, "", "  ")
	return resultsJSON
}

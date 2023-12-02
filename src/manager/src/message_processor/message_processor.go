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

func (mp *MessageProcessor) ProcessMessage(message map[string]json.RawMessage) {
	var value float64
	var parameters [3]float64
	err := json.Unmarshal(message["value"], &value)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	err = json.Unmarshal(message["parameters"], &parameters)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

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
		var paramsAmount uint64
		err = json.Unmarshal(message["paramsAmount"], &paramsAmount)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
		mp.currentAverage = (mp.currentAverage*mp.totalParameters + value*float64(paramsAmount)) / (mp.totalParameters + float64(paramsAmount))
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

package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host    string  `json:"host"`
	Port    int     `json:"port"`
	Queues  Queues  `json:"queues"`
	Metrics Metrics `json:"metrics"`
}

type Queues struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Metrics struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func readConfig(configPath string) Config {
	var config Config

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Unable to unmarshal JSON: %v", err)
	}

	return config
}

func GetConfig() Config {
	if os.Getenv("ENV") == "local" {
		return readConfig("./src/resources/config_local.json")
	}
	return readConfig("./src/resources/config.json")
}

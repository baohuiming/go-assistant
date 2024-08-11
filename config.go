package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Weather struct {
		APIKey    string  `yaml:"api_key"`
		Latitude  float64 `yaml:"latitude"`
		Longitude float64 `yaml:"longitude"`
	} `yaml:"weather"`

	DLNA struct {
		DeviceName string `yaml:"device_name"`
	} `yaml:"dlna"`

	LLM struct {
		ApiKey  string `yaml:"api_key"`
		Prompt  string `yaml:"prompt"`
		BaseUrl string `yaml:"base_url"`
		Model   string `yaml:"model"`
	} `yaml:"llm"`
}

func ReadConfig() *Config {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML data: %v", err)
	}

	return &config
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Settings contains the config.json information for configuring the listening port, monitored application details, etc
type Settings struct {
	Target struct {
		HealthEndpoint string `json:"health_endpoint"` // "health_endpoint": "http://localhost/health",
		SmokeEndpoint  string `json:"smoke_endpoint"`  // "smoke_endpoint": "http://localhost/health",
	} `json:"targets"`
	Service struct {
		HTTPPort       string `json:"http_port"`       // port to listen on for web interface (5704),
		SmokeInterval  int    `json:"smoke_interval"`  // how frequently to poll endpoint for smoke status
		HealthInterval int    `json:"health_interval"` // how frequently to poll endpoint for health status
	} `json:"service"`
}

func (s *Settings) parseSettingsFile() {
	confFile := "../../init/config.json"
	if len(os.Args) > 1 {
		confFile = os.Args[1]
	}

	file, err := os.Open(confFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileContent, err := os.Open(confFile)
	if err != nil {
		fmt.Println("Could not open config file", err.Error())
	}

	jsonParser := json.NewDecoder(fileContent)
	if err = jsonParser.Decode(&s); err != nil {
		fmt.Println("Could not load config file. Check JSON formatting.", err.Error())
	}
}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Settings contains the config.json information for configuring the listening port, monitored application details, etc
type Settings struct {
	Targets []Target `json:"targets"`
	Service struct {
		HTTPPort string `json:"http_port"` // port to listen on for web interface (5704),
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

	// Populate global STATUS with targets from config file
	s.populateTargets()
}

func (s *Settings) populateTargets() {
	for i := range s.Targets {
		s.Targets[i].ID = i
		STATUS.Targets = append(STATUS.Targets, s.Targets[i])
	}
}

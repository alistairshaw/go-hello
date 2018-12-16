package config

import (
	"encoding/json"
	"os"
)

//SMTPConfiguration contains SMTP details of server to use
type SMTPConfiguration struct {
	SMTPHost         string
	SMTPUser         string
	SMTPEmailAddress string
	SMTPPassword     string
	SMTPPort         int
}

//LogConfiguration contains configuration for logging
type LogConfiguration struct {
	LogRoutesToConsole bool
}

//SMTP retrieves smtp config values from the config file
func SMTP() (config SMTPConfiguration, err error) {
	configuration := SMTPConfiguration{}
	filename := "config/config.json"
	file, err := os.Open(filename)
	if err != nil {
		return configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}

//LogRoutesToConsole determines if this option is enabled
func LogRoutesToConsole() bool {
	configuration := LogConfiguration{}
	filename := "config/config.json"
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return false
	}

	return configuration.LogRoutesToConsole
}

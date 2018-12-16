package config

import (
	"encoding/json"
	"os"
)

//SMTPConfiguration : SMTP details of server to use
type SMTPConfiguration struct {
	SMTPHost         string
	SMTPUser         string
	SMTPEmailAddress string
	SMTPPassword     string
	SMTPPort         int
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

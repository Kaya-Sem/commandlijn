package cmd

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	DeLijnAPIKey string  `yaml:"delijn_api_key"`
	Aliases      []Alias `yaml:"aliases"`
}

type Alias struct {
	Name     string          `yaml:"name"`
	Provider TransitProvider `yaml:"provider"`
	ID       []string        `yaml:"ID"`
}

func GetConfig() ConfigStruct {
	data, err := os.ReadFile(getConfigFilePath())
	if err != nil {
		log.Printf("Error getting config file at %s: %v", getConfigFilePath(), err)
		os.Exit(ErrFileRead)
	}

	var config ConfigStruct

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("Error parsing config: %v", err)
		os.Exit(ErrUnmarshal)
	}

	return config
}

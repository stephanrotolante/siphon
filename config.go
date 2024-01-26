package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type SiphonConfig struct {
	proxy struct {
		port int    `yaml:"port"`
		host string `yaml:"host"`
	} `yaml:"proxy"`
}

func ReadConfig(cliFlags *CliFlags) SiphonConfig {
	yamlFile, err := os.ReadFile(cliFlags.configPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Parse YAML
	var siphonConfig SiphonConfig
	err = yaml.Unmarshal(yamlFile, &siphonConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	return siphonConfig
}

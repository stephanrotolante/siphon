package config

import (
	"log"
	"os"

	"github.com/stephanrotolante/siphon/cli"
	"gopkg.in/yaml.v3"
)

type SiphonConfig struct {
	Proxy struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"proxy"`

	Https struct {
		Enabled  bool   `yaml:"enabled"`
		CertPath string `yaml:"cert_path"`
		KeyPath  string `yaml:"key_path"`
	} `yaml:"https"`

	Target struct {
		DefaultHost string `yaml:"default_host"`
		TargetHost  string `yaml:"target_host"`
		Probability int    `yaml:"probability"`
	} `yaml:"target"`

	AllowedVerbs []string `yaml:"allowed_verbs"`
}

func createDefaultConfig() SiphonConfig {
	var config = SiphonConfig{}

	// default proxy
	config.Proxy.Host = ""
	config.Proxy.Port = 3250

	// default https
	config.Https.Enabled = false
	config.Https.CertPath = ""
	config.Https.KeyPath = ""

	// default target
	config.Target.Probability = 100

	return config
}

func ReadConfig(cliFlags *cli.CliFlags) SiphonConfig {
	yamlFile, err := os.ReadFile(cliFlags.ConfigPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var siphonConfig = createDefaultConfig()

	err = yaml.Unmarshal(yamlFile, &siphonConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	return siphonConfig
}

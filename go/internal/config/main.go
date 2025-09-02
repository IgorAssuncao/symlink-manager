package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tools []map[string]Tool `yaml:"tools"`
}

type Tool struct {
	Path   string `yaml:"path"`
	Target string `yaml:"target"`
}

func (c *Config) GetConfig() *Config {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("%v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)

	return c
}

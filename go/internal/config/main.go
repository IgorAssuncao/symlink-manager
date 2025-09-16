package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tools map[string]Tool `yaml:"tools"`
}

type Tool struct {
	Path   string `yaml:"path"`
	Target string `yaml:"target"`
}

func NewConfig(configPath string) *Config {
	c := &Config{}
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("%v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)

	return c
}

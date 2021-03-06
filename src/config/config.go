package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Package vars and default values
var (
	configCache *Config

	configFilePath = "./config.yaml"
	configDefault  = &Config{
		DebugMode:     false,
	}
)

// Config read from yaml file
type Config struct {
	DebugMode                    bool     `yaml:"debugMode"`
	SlackToken                   string   `yaml:"slackToken"`
}

// SetFilePath to... set config filepath :)
func SetFilePath(path string) {
	configFilePath = path
}

// Get returns current config
func Get() *Config {

	// Load config the first time
	if configCache == nil {
		conf, err := loadConfigFromFile(configFilePath)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		configCache = conf
	}

	return configCache
}

// open and decode yaml file
func loadConfigFromFile(filePath string) (*Config, error) {
	conf := configDefault

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileData, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

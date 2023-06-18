package main

import (
	"fmt"

	"github.com/koddr/gosl"
)

// newConfig provides a new config instance.
func newConfig() (*Config, error) {
	// Check, if the file name is too short.
	if configFilePath == "" {
		return nil, fmt.Errorf("invalid format of config file, see: %s", WikiPageURL)
	}

	// Create a new config instance.
	config := &Config{}

	// Load config from path or HTTP by the given file format.
	_, err := gosl.ParseFileWithEnvToStruct(configFilePath, envPrefix, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

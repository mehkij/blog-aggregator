package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	var config Config

	e := json.Unmarshal(jsonData, &config)

	return config, e
}

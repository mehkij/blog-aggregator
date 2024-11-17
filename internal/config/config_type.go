package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	var filepath string

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filepath = home + configFileName

	return filepath, nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	// 0644: Owner can read/write; others can read
	e := os.WriteFile(filepath, data, 0644)

	return e
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	err := write(*c)

	return err
}

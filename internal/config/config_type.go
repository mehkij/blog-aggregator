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

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
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

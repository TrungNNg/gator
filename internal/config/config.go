package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBUrl    string `json:"db_url"`
	UserName string `json:"current_user_name,omitempty"`
}

func (c Config) SetUser(newName string) error {
	file, err := os.Create(getConfigFilePath())
	if err != nil {
		fmt.Println(err)
		return err
	}
	encoder := json.NewEncoder(file)
	c.UserName = newName
	err = encoder.Encode(c)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Read() Config {
	file, err := os.Open(getConfigFilePath())
	if err != nil {
		fmt.Println(err)
		return Config{}
	}
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}
	return cfg
}

// HELPER FUNCTIONS

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return homeDir + "/.gatorconfig.json"
}

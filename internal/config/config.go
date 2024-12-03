package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type AuthService struct {
	Url string `yaml:"url"`
}

type Config struct {
	Http struct {
		Port string `yaml:"port"`
	} `yaml:"http"`
	AuthService AuthService `yaml:"auth_service"`
}

func Init(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling yaml: %v", err)
	}

	return &config, nil
}

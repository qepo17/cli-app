package main

import "github.com/kelseyhightower/envconfig"

type Config struct {
	OpenRouterAPIKey string `envconfig:"OPEN_ROUTER_API_KEY" default:""`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

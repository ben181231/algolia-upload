package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config represents a configuration of the script
type Config struct {
	AppID     string `envconfig:"app_id" required:"true"`
	AdminKey  string `envconfig:"admin_key" required:"true"`
	IndexName string `envconfig:"index_name" required:"true"`
}

func ReadFromEnv() (*Config, error) {
	conf := Config{}
	if err := envconfig.Process("algolia", &conf); err != nil {
		return nil, fmt.Errorf("failed to read config from environment: %w", err)
	}

	return &conf, nil
}

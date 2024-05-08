package models

import (
	"errors"
	"os"
)

type Config struct {
	APIKey string
}

func GetConfig() (*Config, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		return nil, errors.New("GOOGLE_API_KEY environment variable not set")
	}

	return &Config{
		APIKey: apiKey,
	}, nil
}

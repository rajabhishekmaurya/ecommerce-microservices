package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
	}

	if cfg.Port == "" {
		cfg.Port = "8083"
	}

	return cfg
}

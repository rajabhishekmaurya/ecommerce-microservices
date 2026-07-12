package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string
	Port      string
	JWTSecret string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName:   os.Getenv("APP_NAME"),
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.Port == "" {
		cfg.Port = "8081"
	}

	return cfg
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string
	Port      string
	JWTSecret string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName:   os.Getenv("APP_NAME"),
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	if cfg.Port == "" {
		cfg.Port = "8081"
	}

	return cfg
}

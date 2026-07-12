package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	KafkaBroker string
	KafkaTopic string
}

func Load() *Config {

	_ = godotenv.Load()

	cfg := &Config{
		AppName:     os.Getenv("APP_NAME"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		KafkaTopic:  os.Getenv("KAFKA_TOPIC"),
	}

	if cfg.AppName == "" {
		cfg.AppName = "Notification Service"
	}

	if cfg.KafkaBroker == "" {
		cfg.KafkaBroker = "localhost:9092"
	}

	if cfg.KafkaTopic == "" {
		cfg.KafkaTopic = "payment-events"
	}

	return cfg
}

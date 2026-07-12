package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string

	AuthServiceURL         string
	UserServiceURL         string
	OrderServiceURL        string
	PaymentServiceURL      string
	NotificationServiceURL string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),

		AuthServiceURL:         os.Getenv("AUTH_SERVICE_URL"),
		UserServiceURL:         os.Getenv("USER_SERVICE_URL"),
		OrderServiceURL:        os.Getenv("ORDER_SERVICE_URL"),
		PaymentServiceURL:      os.Getenv("PAYMENT_SERVICE_URL"),
		NotificationServiceURL: os.Getenv("NOTIFICATION_SERVICE_URL"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	log.Println("Configuration Loaded")

	return cfg
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string

	PaymentServiceAddr string
}
func Load() *Config {
    _ = godotenv.Load()

    cfg := &Config{
        AppName: os.Getenv("APP_NAME"),
        Port: os.Getenv("PORT"),

        PaymentServiceAddr: os.Getenv("PAYMENT_SERVICE_ADDR"),
    }

    if cfg.Port == "" {
        cfg.Port = "8083"
    }

    if cfg.PaymentServiceAddr == "" {
        cfg.PaymentServiceAddr = "localhost:50051"
    }

    return cfg
}

package server

import (
	"log"

	"github.com/rajabhishekmaurya/ecommerce-microservices/notification-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/notification-service/internal/service"
)

type Server struct {
	cfg      *config.Config
	consumer *service.KafkaConsumer
}

func New() *Server {

	cfg := config.Load()

	return &Server{
		cfg:      cfg,
		consumer: service.NewKafkaConsumer(),
	}
}

func (s *Server) Start() error {

	log.Printf("%s started", s.cfg.AppName)

	return s.consumer.Start()
}

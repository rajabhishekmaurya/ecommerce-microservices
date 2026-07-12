package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rajabhishekmaurya/ecommerce-microservices/order-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/order-service/internal/handler"
	"github.com/rajabhishekmaurya/ecommerce-microservices/order-service/internal/service"
)

type Server struct {
	cfg *config.Config
	e   *echo.Echo
}

func New() (*Server, error) {

	cfg := config.Load()

	orderService, err := service.NewOrderService()
	if err != nil {
		return nil, err
	}

	orderHandler := handler.NewOrderHandler(orderService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", orderHandler.Health)
	e.POST("/orders", orderHandler.CreateOrder)

	return &Server{
		cfg: cfg,
		e:   e,
	}, nil
}

func (s *Server) Start() error {

	log.Printf("%s started on :%s", s.cfg.AppName, s.cfg.Port)

	return s.e.Start(":" + s.cfg.Port)
}

func (s *Server) Shutdown() error {
	return s.e.Shutdown(nil)
}

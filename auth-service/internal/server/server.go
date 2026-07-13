package server

import (
	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/handler"
	commonmiddleware "github.com/rajabhishekmaurya/ecommerce-microservices/common/middleware"
	"github.com/rajabhishekmaurya/ecommerce-microservices/common/monitoring"
)

type Server struct {
	echo *echo.Echo
	cfg  *config.Config
}

func New() *Server {

	cfg := config.Load()

	e := echo.New()

	e.HideBanner = true

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	e.Use(commonmiddleware.PrometheusMiddleware())

	e.GET("/metrics", monitoring.Handler())

	handler.RegisterRoutes(e, cfg)

	return &Server{
		echo: e,
		cfg:  cfg,
	}
}

func (s *Server) Start() error {
	log.Printf("%s started on :%s", s.cfg.AppName, s.cfg.Port)
	return s.echo.Start(":" + s.cfg.Port)
}

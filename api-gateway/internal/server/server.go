package server

import (
	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/rajabhishekmaurya/ecommerce-microservices/api-gateway/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/api-gateway/internal/handler"
	customMiddleware "github.com/rajabhishekmaurya/ecommerce-microservices/api-gateway/internal/middleware"
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
	e.Use(customMiddleware.RequestID)

	handler.Register(e, cfg)

	return &Server{
		echo: e,
		cfg:  cfg,
	}
}

func (s *Server) Start() error {

	log.Printf("%s listening on %s", s.cfg.AppName, s.cfg.Port)

	return s.echo.Start(":" + s.cfg.Port)

}

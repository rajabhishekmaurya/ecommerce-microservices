package server

import (
	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/handler"
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

	handler.RegisterRoutes(e)

	return &Server{
		echo: e,
		cfg:  cfg,
	}
}

func (s *Server) Start() error {

	log.Printf("%s started on :%s", s.cfg.AppName, s.cfg.Port)

	return s.echo.Start(":" + s.cfg.Port)
}

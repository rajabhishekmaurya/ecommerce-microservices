package server

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/handler"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/repository"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/service"
)

type Server struct {
	echo *echo.Echo
	cfg  *config.Config
	db   *sql.DB
}

func New() *Server {

	cfg := config.Load()

	db, err := config.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(db)

	authService := service.NewAuthService(repo, cfg)

	authHandler := handler.NewAuthHandler(authService)

	e := echo.New()

	e.HideBanner = true

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	handler.RegisterRoutes(e, authHandler)

	return &Server{
		echo: e,
		cfg:  cfg,
		db:   db,
	}
}

func (s *Server) Start() error {

	log.Printf("%s started on :%s", s.cfg.AppName, s.cfg.Port)

	return s.echo.Start(":" + s.cfg.Port)
}

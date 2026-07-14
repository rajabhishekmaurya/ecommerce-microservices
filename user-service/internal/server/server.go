package server

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/handler"
	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/repository"
	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/service"
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

	userService := service.NewUserService(repo)

	userHandler := handler.NewUserHandler(userService)

	e := echo.New()

	e.HideBanner = true

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	handler.RegisterRoutes(e, userHandler)

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

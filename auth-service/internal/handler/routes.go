package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/config"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config) {

	auth := NewAuthHandler(cfg)

	e.GET("/health", auth.Health)

	e.POST("/register", auth.Register)

	e.POST("/login", auth.Login)

}

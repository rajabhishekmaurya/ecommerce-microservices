package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/rajabhishekmaurya/ecommerce-microservices/api-gateway/internal/config"
)

func Register(e *echo.Echo, cfg *config.Config) {

	e.GET("/health", Health)

	auth := e.Group("/auth")
	auth.Any("/*", ReverseProxy(cfg.AuthServiceURL, "/auth"))

	users := e.Group("/users")
	users.Any("/*", ReverseProxy(cfg.UserServiceURL, "/users"))

}

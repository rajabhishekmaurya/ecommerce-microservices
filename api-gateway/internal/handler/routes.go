package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/rajabhishekmaurya/ecommerce-microservices/api-gateway/internal/config"
)

func Register(e *echo.Echo, cfg *config.Config) {

	e.GET("/health", Health)

	auth := e.Group("/auth")
	auth.Any("/*", ReverseProxy(cfg.AuthServiceURL))

	users := e.Group("/users")
	users.Any("/*", ReverseProxy(cfg.UserServiceURL))

	orders := e.Group("/orders")
	orders.Any("", ReverseProxy(cfg.OrderServiceURL))
	orders.Any("/*", ReverseProxy(cfg.OrderServiceURL))
}

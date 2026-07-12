package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {

	health := NewHealthHandler()

	e.GET("/", health.Health)
	e.GET("/health", health.Health)
}

package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(
	e *echo.Echo,
	userHandler *UserHandler,
) {

	e.GET("/health", userHandler.Health)

	e.POST("/users", userHandler.Register)
}

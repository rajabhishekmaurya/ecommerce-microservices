package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, authHandler *AuthHandler) {

	e.GET("/health", authHandler.Health)

	e.POST("/login", authHandler.Login)
}

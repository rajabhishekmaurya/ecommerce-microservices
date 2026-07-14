package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, auth *AuthHandler) {

	e.GET("/health", auth.Health)

	e.POST("/register", auth.Register)

	e.POST("/login", auth.Login)

}

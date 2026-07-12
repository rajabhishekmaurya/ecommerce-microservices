package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {

	e.GET("/health", Health)

	e.POST("/users", CreateUser)
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)
}

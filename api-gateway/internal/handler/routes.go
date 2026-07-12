package handler

import "github.com/labstack/echo/v4"

func Register(e *echo.Echo) {

	e.GET("/", Health)
	e.GET("/health", Health)

}

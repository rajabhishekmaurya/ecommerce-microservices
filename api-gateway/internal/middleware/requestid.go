package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		requestID := uuid.New().String()

		c.Response().Header().Set("X-Request-ID", requestID)

		return next(c)
	}
}

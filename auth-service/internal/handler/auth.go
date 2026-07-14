package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Health(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"service": "auth-service",
		"status":  "UP",
	})
}

func (h *AuthHandler) Login(c echo.Context) error {

	var req service.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	resp, err := h.service.Login(
		c.Request().Context(),
		&req,
	)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/model"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/service"
)

type AuthHandler struct {
	cfg     *config.Config
	service *service.AuthService
}

func NewAuthHandler(
	cfg *config.Config,
	service *service.AuthService,
) *AuthHandler {

	return &AuthHandler{
		cfg:     cfg,
		service: service,
	}
}

func (h *AuthHandler) Health(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"service": "auth-service",
		"status":  "UP",
	})

}

func (h *AuthHandler) Register(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User Registered Successfully",
	})

}
func (h *AuthHandler) Login(c echo.Context) error {

	req := new(model.LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.Login(req)
	if err != nil {

		if errors.Is(err, service.ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

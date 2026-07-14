package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/model"
	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Health(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"service": "user-service",
		"status":  "UP",
	})
}

func (h *UserHandler) Register(c echo.Context) error {

	var user model.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.Register(c.Request().Context(), &user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	user.Password = "*****"

	return c.JSON(http.StatusCreated, user)
}

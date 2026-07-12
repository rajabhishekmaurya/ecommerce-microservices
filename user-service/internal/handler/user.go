package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/model"
	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/service"
)

func Health(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"service": "user-service",
		"status":  "UP",
	})
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, service.GetUsers())
}

func CreateUser(c echo.Context) error {

	var user model.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, service.CreateUser(user))
}

func GetUser(c echo.Context) error {

	user, ok := service.GetUser(c.Param("id"))

	if !ok {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {

	var user model.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if !service.UpdateUser(c.Param("id"), user) {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, "Updated Successfully")
}

func DeleteUser(c echo.Context) error {

	if !service.DeleteUser(c.Param("id")) {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, "Deleted Successfully")
}

package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/config"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Dummy validation
	if req.Username != "admin" || req.Password != "admin123" {

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid Username or Password",
		})

	}

	claims := jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(h.cfg.JWTSecret))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})

}

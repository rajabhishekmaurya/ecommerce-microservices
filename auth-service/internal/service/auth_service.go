package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/model"
)

type AuthService struct {
	cfg *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		cfg: cfg,
	}
}

func (s *AuthService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {

	// Dummy validation
	if req.Username != "admin" || req.Password != "admin123" {
		return nil, ErrInvalidCredentials
	}

	claims := jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token: tokenString,
	}, nil
}

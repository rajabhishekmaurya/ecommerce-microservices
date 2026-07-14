package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.UserRepository
	cfg  *config.Config
}
type LoginRequest struct {
	Username string
	Password string
}
type LoginResponse struct {
	Token string
}

func NewAuthService(repo repository.UserRepository, cfg *config.Config) *AuthService {

	return &AuthService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {

	user, err := s.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte(s.cfg.JWTSecret),
	)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: tokenString,
	}, nil
}

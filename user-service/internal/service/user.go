package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/model"
	"github.com/rajabhishekmaurya/ecommerce-microservices/user-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, user *model.User) error {

	// Check if username already exists
	existingUser, err := s.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// Save user
	return s.repo.Create(ctx, user)
}

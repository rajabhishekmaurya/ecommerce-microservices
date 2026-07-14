package repository

import (
	"context"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/model"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

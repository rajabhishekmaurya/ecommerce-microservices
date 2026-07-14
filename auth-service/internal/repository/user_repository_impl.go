package repository

import (
	"context"
	"database/sql"

	"github.com/rajabhishekmaurya/ecommerce-microservices/auth-service/internal/model"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByUsername(ctx context.Context,username string) (*model.User, error) {

	query := `
	SELECT
		id,
		username,
		email,
		password,
		created_at,
		updated_at
	FROM users
	WHERE username = ?
	`

	user := &model.User{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

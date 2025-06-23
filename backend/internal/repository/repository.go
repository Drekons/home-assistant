package repository

import (
	"context"
	"github.com/Drekons/home-assistant/backend/internal/database"
	"github.com/Drekons/home-assistant/backend/internal/interfaces"
	"github.com/pkg/errors"
)

type Repository struct {
	users interfaces.UserRepository
}

func NewRepository(ctx context.Context, db *database.DB) (interfaces.Repository, error) {
	users, err := NewUserRepository(ctx, db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user repository")
	}

	return &Repository{users}, nil
}

func (r *Repository) Users() interfaces.UserRepository {
	return r.users
}

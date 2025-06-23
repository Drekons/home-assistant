package interfaces

import (
	"context"
	"github.com/Drekons/home-assistant/backend/internal/entity"
)

type Repository interface {
	Users() UserRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

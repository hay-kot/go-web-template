package repo

import (
	"context"

	"github.com/hay-kot/git-web-template/backend/internal/dtos"
)

type UserRepository interface {
	// GetOneId returns a user by id
	GetOneId(id int, ctx context.Context) (dtos.UserOut, error)
	// GetOneEmail returns a user by email
	GetOneEmail(email string, ctx context.Context) (dtos.UserOut, error)
	// GetAll returns all users
	GetAll(ctx context.Context) ([]dtos.UserOut, error)
	// Create creates a new user
	Create(user *dtos.UserCreate, ctx context.Context) (dtos.UserOut, error)
	// Update updates a user
	Update(user *dtos.UserCreate, ctx context.Context) error
	// Delete deletes a user
	Delete(id int, ctx context.Context) error

	DeleteAll(ctx context.Context) error
}

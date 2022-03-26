package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/types"
)

type UserRepository interface {
	// GetOneId returns a user by id
	GetOneId(id uuid.UUID, ctx context.Context) (types.UserOut, error)
	// GetOneEmail returns a user by email
	GetOneEmail(email string, ctx context.Context) (types.UserOut, error)
	// GetAll returns all users
	GetAll(ctx context.Context) ([]types.UserOut, error)
	// Get Super Users
	GetSuperusers(ctx context.Context) ([]types.UserOut, error)
	// Create creates a new user
	Create(user *types.UserCreate, ctx context.Context) (types.UserOut, error)
	// Update updates a user
	Update(user *types.UserCreate, ctx context.Context) error
	// Delete deletes a user
	Delete(id uuid.UUID, ctx context.Context) error

	DeleteAll(ctx context.Context) error
}

package repo

import "context"

type UserRepository interface {
	// GetOneId returns a user by id
	GetOneId(id int, ctx context.Context) (UserOut, error)
	// GetOneEmail returns a user by email
	GetOneEmail(email string, ctx context.Context) (UserOut, error)
	// GetAll returns all users
	GetAll(ctx context.Context) ([]UserOut, error)
	// Create creates a new user
	Create(user *UserCreate, ctx context.Context) error
	// Update updates a user
	Update(user *UserCreate, ctx context.Context) error
	// Delete deletes a user
	Delete(id int, ctx context.Context) error
}

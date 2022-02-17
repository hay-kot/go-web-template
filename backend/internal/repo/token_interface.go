package repo

import (
	"context"

	"github.com/hay-kot/git-web-template/backend/internal/dtos"
)

type TokenRepository interface {
	// GetUserFromToken get's a user from a token
	GetUserFromToken(token []byte, ctx context.Context) (dtos.UserOut, error)
	// Creates a token for a user
	CreateToken(createToken dtos.UserAuthTokenCreate, ctx context.Context) (dtos.UserAuthToken, error)
	// DeleteToken remove a single token from the database - equivalent to revoke or logout
	DeleteToken(token []byte, ctx context.Context) error
	// PurgeExpiredTokens removes all expired tokens from the database
	PurgeExpiredTokens(ctx context.Context) (int, error)
	// DeleteAll removes all tokens from the database
	DeleteAll(ctx context.Context) (int, error)
}

package repo

import (
	"context"
	"time"

	"github.com/hay-kot/git-web-template/backend/ent"
	"github.com/hay-kot/git-web-template/backend/ent/authtokens"
	"github.com/hay-kot/git-web-template/backend/internal/dtos"
	"github.com/hay-kot/git-web-template/backend/internal/mapper"
)

type EntTokenRepository struct {
	db *ent.Client
}

// GetUserFromToken get's a user from a token
func (r *EntTokenRepository) GetUserFromToken(token []byte, ctx context.Context) (dtos.UserOut, error) {
	dbToken, err := r.db.AuthTokens.Query().
		Where(authtokens.Token(token)).
		Where(authtokens.ExpiresAtGTE(time.Now())).
		WithUser().
		Only(ctx)

	if err != nil {
		return dtos.UserOut{}, err
	}

	return mapper.UserOutFromModel(*dbToken.Edges.User), nil
}

// Creates a token for a user
func (r *EntTokenRepository) CreateToken(createToken dtos.UserAuthTokenCreate, ctx context.Context) (dtos.UserAuthToken, error) {
	tokenOut := dtos.UserAuthToken{}

	dbToken, err := r.db.AuthTokens.Create().
		SetToken(createToken.TokenHash).
		SetUserID(createToken.UserId).
		SetExpiresAt(createToken.ExpiresAt).
		Save(ctx)

	if err != nil {
		return tokenOut, err
	}

	tokenOut.TokenHash = dbToken.Token
	tokenOut.UserId = createToken.UserId
	tokenOut.CreatedAt = dbToken.CreatedAt
	tokenOut.ExpiresAt = dbToken.ExpiresAt

	return tokenOut, nil
}

// DeleteToken remove a single token from the database - equivalent to revoke or logout
func (r *EntTokenRepository) DeleteToken(token []byte, ctx context.Context) error {
	_, err := r.db.AuthTokens.Delete().Where(authtokens.Token(token)).Exec(ctx)
	return err
}

// PurgeExpiredTokens removes all expired tokens from the database
func (r *EntTokenRepository) PurgeExpiredTokens(ctx context.Context) (int, error) {
	tokensDeleted, err := r.db.AuthTokens.Delete().Where(authtokens.ExpiresAtLTE(time.Now())).Exec(ctx)

	if err != nil {
		return 0, err
	}

	return tokensDeleted, nil
}

func (r *EntTokenRepository) DeleteAll(ctx context.Context) (int, error) {
	amount, err := r.db.AuthTokens.Delete().Exec(ctx)
	return amount, err
}

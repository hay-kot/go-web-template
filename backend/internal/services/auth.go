package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
)

var (
	oneWeek      = time.Hour * 24 * 7
	InvalidLogin = errors.New("invalid username or password")
	InvalidToken = errors.New("invalid token")
)

type AuthService struct {
	repos *repo.AllRepos
}

func (as *AuthService) createToken(ctx context.Context, userId uuid.UUID) (types.UserAuthTokenDetail, error) {
	newToken := hasher.GenerateToken()

	created, err := as.repos.AuthTokens.CreateToken(types.UserAuthTokenCreate{
		UserId:    userId,
		TokenHash: newToken.Hash,
		ExpiresAt: time.Now().Add(oneWeek),
	}, ctx)

	return types.UserAuthTokenDetail{Raw: newToken.Raw, ExpiresAt: created.ExpiresAt}, err
}

func (as *AuthService) Login(ctx context.Context, username, password string) (types.UserAuthTokenDetail, error) {
	usr, err := as.repos.Users.GetOneEmail(username, ctx)

	if err != nil || !hasher.CheckPasswordHash(password, usr.Password) {
		return types.UserAuthTokenDetail{}, InvalidLogin
	}

	return as.createToken(ctx, usr.Id)
}

func (as *AuthService) Logout(ctx context.Context, token string) error {
	hash := hasher.HashToken(token)
	err := as.repos.AuthTokens.DeleteToken(hash, ctx)
	return err
}

func (as *AuthService) RenewToken(ctx context.Context, token string) (types.UserAuthTokenDetail, error) {
	hash := hasher.HashToken(token)

	dbToken, err := as.repos.AuthTokens.GetUserFromToken(hash, ctx)

	if err != nil {
		return types.UserAuthTokenDetail{}, InvalidToken
	}

	newToken, _ := as.createToken(ctx, dbToken.Id)

	return newToken, nil
}
package repo

import (
	"context"
	"testing"
	"time"

	"github.com/hay-kot/git-web-template/backend/internal/dtos"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/stretchr/testify/assert"
)

func Test_EntAuthTokenRepo_CreateToken(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	user := UserFactory()

	userOut, _ := testRepos.Users.Create(&user, ctx)

	expiresAt := time.Now().Add(time.Hour)

	generatedToken := hasher.GenerateToken()

	token, err := testRepos.AuthTokens.CreateToken(dtos.UserAuthTokenCreate{
		TokenHash: generatedToken.Hash,
		ExpiresAt: expiresAt,
		UserId:    userOut.Id,
	}, ctx)

	assert.NoError(err)
	assert.Equal(userOut.Id, token.UserId)
	assert.Equal(expiresAt, token.ExpiresAt)

	// Cleanup
	err = testRepos.Users.Delete(userOut.Id, ctx)
	_, err = testRepos.AuthTokens.DeleteAll(ctx)
}

func Test_EntAuthTokenRepo_GetUserByToken(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	user := UserFactory()
	userOut, _ := testRepos.Users.Create(&user, ctx)

	expiresAt := time.Now().Add(time.Hour)
	generatedToken := hasher.GenerateToken()

	token, err := testRepos.AuthTokens.CreateToken(dtos.UserAuthTokenCreate{
		TokenHash: generatedToken.Hash,
		ExpiresAt: expiresAt,
		UserId:    userOut.Id,
	}, ctx)

	// Get User from token
	foundUser, err := testRepos.AuthTokens.GetUserFromToken(token.TokenHash, ctx)

	assert.NoError(err)
	assert.Equal(userOut.Id, foundUser.Id)
	assert.Equal(userOut.Name, foundUser.Name)
	assert.Equal(userOut.Email, foundUser.Email)

	// Cleanup
	err = testRepos.Users.Delete(userOut.Id, ctx)
	_, err = testRepos.AuthTokens.DeleteAll(ctx)
}

func Test_EntAuthTokenRepo_PurgeExpiredTokens(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	user := UserFactory()
	userOut, _ := testRepos.Users.Create(&user, ctx)

	createdTokens := []dtos.UserAuthToken{}

	for i := 0; i < 5; i++ {
		expiresAt := time.Now()
		generatedToken := hasher.GenerateToken()

		createdToken, err := testRepos.AuthTokens.CreateToken(dtos.UserAuthTokenCreate{
			TokenHash: generatedToken.Hash,
			ExpiresAt: expiresAt,
			UserId:    userOut.Id,
		}, ctx)

		assert.NoError(err)
		assert.NotNil(createdToken)

		createdTokens = append(createdTokens, createdToken)

	}

	// Purge expired tokens
	tokensDeleted, err := testRepos.AuthTokens.PurgeExpiredTokens(ctx)

	assert.NoError(err)
	assert.Equal(5, tokensDeleted)

	// Check if tokens are deleted
	for _, token := range createdTokens {
		_, err := testRepos.AuthTokens.GetUserFromToken(token.TokenHash, ctx)
		assert.Error(err)
	}

	// Cleanup
	err = testRepos.Users.Delete(userOut.Id, ctx)
	_, err = testRepos.AuthTokens.DeleteAll(ctx)
}

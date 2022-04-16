package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_SetAuthContext(t *testing.T) {
	user := &types.UserOut{
		ID: uuid.New(),
	}

	token := uuid.New().String()

	ctx := SetUserContext(context.Background(), user, token)

	ctxUser := GetUserContext(ctx)

	assert.NotNil(t, ctxUser)
	assert.Equal(t, user.ID, ctxUser.ID)

	ctxUserToken := GetUserTokenFromContext(ctx)
	assert.NotEmpty(t, ctxUserToken)
}

func Test_SetAuthContext_Nulls(t *testing.T) {
	ctx := SetUserContext(context.Background(), nil, "")

	ctxUser := GetUserContext(ctx)

	assert.Nil(t, ctxUser)

	ctxUserToken := GetUserTokenFromContext(ctx)
	assert.Empty(t, ctxUserToken)
}

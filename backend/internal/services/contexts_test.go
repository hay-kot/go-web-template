package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/dtos"
	"github.com/stretchr/testify/assert"
)

func Test_SetAuthContext(t *testing.T) {
	user := &dtos.UserOut{
		Id: uuid.New(),
	}

	token := uuid.New().String()

	ctx := SetAuthContext(context.Background(), user, token)

	ctxUser := UserFromContext(ctx)

	assert.NotNil(t, ctxUser)
	assert.Equal(t, user.Id, ctxUser.Id)

	ctxUserToken := UserTokenFromContext(ctx)
	assert.NotEmpty(t, ctxUserToken)
}

func Test_SetAuthContext_Nulls(t *testing.T) {
	ctx := SetAuthContext(context.Background(), nil, "")

	ctxUser := UserFromContext(ctx)

	assert.Nil(t, ctxUser)

	ctxUserToken := UserTokenFromContext(ctx)
	assert.Empty(t, ctxUserToken)
}

package services

import (
	"context"

	"github.com/hay-kot/git-web-template/backend/internal/dtos"
)

type contextKeys struct {
	name string
}

var (
	ContextUser      = &contextKeys{name: "user"}
	ContextUserToken = &contextKeys{name: "UserToken"}
)

func UserFromContext(ctx context.Context) *dtos.UserOut {
	if val := ctx.Value(ContextUser); val != nil {
		return val.(*dtos.UserOut)
	}
	return nil
}

func UserTokenFromContext(ctx context.Context) string {
	if val := ctx.Value(ContextUserToken); val != nil {
		return val.(string)
	}
	return ""
}

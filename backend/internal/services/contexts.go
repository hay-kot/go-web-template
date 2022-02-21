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

// SetAuthContext is a helper function that sets the ContextUser and ContextUserToken
// values within the context of a web request (or any context).
func SetAuthContext(ctx context.Context, user *dtos.UserOut, token string) context.Context {
	ctx = context.WithValue(ctx, ContextUser, user)
	ctx = context.WithValue(ctx, ContextUserToken, token)
	return ctx
}

// UserFromContext is a helper function that returns the user from the context.
func UserFromContext(ctx context.Context) *dtos.UserOut {
	if val := ctx.Value(ContextUser); val != nil {
		return val.(*dtos.UserOut)
	}
	return nil
}

// UserTokenFromContext is a helper function that returns the user token from the context.
func UserTokenFromContext(ctx context.Context) string {
	if val := ctx.Value(ContextUserToken); val != nil {
		return val.(string)
	}
	return ""
}

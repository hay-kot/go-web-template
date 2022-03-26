package services

import (
	"context"

	"github.com/hay-kot/git-web-template/backend/internal/dtos"
)

type contextKeys struct {
	name string
}

var (
	ContextUser      = &contextKeys{name: "User"}
	ContextUserToken = &contextKeys{name: "UserToken"}
)

// SetUserContext is a helper function that sets the ContextUser and ContextUserToken
// values within the context of a web request (or any context).
func SetUserContext(ctx context.Context, user *dtos.UserOut, token string) context.Context {
	ctx = context.WithValue(ctx, ContextUser, user)
	ctx = context.WithValue(ctx, ContextUserToken, token)
	return ctx
}

// GetUserContext is a helper function that returns the user from the context.
func GetUserContext(ctx context.Context) *dtos.UserOut {
	if val := ctx.Value(ContextUser); val != nil {
		return val.(*dtos.UserOut)
	}
	return nil
}

// GetUserTokenFromContext is a helper function that returns the user token from the context.
func GetUserTokenFromContext(ctx context.Context) string {
	if val := ctx.Value(ContextUserToken); val != nil {
		return val.(string)
	}
	return ""
}

package v1

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/internal/repo"
	"github.com/hay-kot/git-web-template/pkgs/logger"
)

type Handlersv1 struct {
	log   logger.SharedLogger
	jwt   *jwtauth.JWTAuth
	repos *repo.AllRepos
}

func NewHandlerV1(prefix string, repos *repo.AllRepos, jwt *jwtauth.JWTAuth, log logger.SharedLogger) (func(s string) string, *Handlersv1) {
	h := &Handlersv1{
		log:   log,
		jwt:   jwt,
		repos: repos,
	}

	v1Base := prefix + "/v1"
	prefixFunc := func(s string) string {
		return v1Base + s
	}

	return prefixFunc, h
}

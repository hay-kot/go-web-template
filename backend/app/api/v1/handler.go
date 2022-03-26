package v1

import (
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
)

type Handlersv1 struct {
	log      *logger.Logger
	repos    *repo.AllRepos
	services *services.AllServices
}

func NewHandlerV1(prefix string, repos *repo.AllRepos, log *logger.Logger, auth *services.AllServices) (func(s string) string, *Handlersv1) {
	h := &Handlersv1{
		log:      log,
		repos:    repos,
		services: auth,
	}

	v1Base := prefix + "/v1"
	prefixFunc := func(s string) string {
		return v1Base + s
	}

	return prefixFunc, h
}

package main

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/ent"
	"github.com/hay-kot/git-web-template/internal/config"
	"github.com/hay-kot/git-web-template/pkgs/logger"
)

type app struct {
	Conf   *config.Config
	logger logger.SharedLogger
	jwt    *jwtauth.JWTAuth
	db     *ent.Client
}

func NewApp(conf *config.Config) *app {
	s := &app{
		Conf: conf,
	}

	s.jwt = jwtauth.New("HS256", []byte("secret"), nil)

	return s
}

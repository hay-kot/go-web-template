package main

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/ent"
	"github.com/hay-kot/git-web-template/internal/config"
	"github.com/hay-kot/git-web-template/pkgs/logger"
	"github.com/hay-kot/git-web-template/pkgs/mailer"
	"github.com/hay-kot/git-web-template/pkgs/server"
)

type app struct {
	conf   *config.Config
	logger *logger.Logger
	mailer mailer.Mailer
	jwt    *jwtauth.JWTAuth
	db     *ent.Client
	server *server.Server
}

func NewApp(conf *config.Config) *app {
	s := &app{
		conf: conf,
	}

	s.jwt = jwtauth.New("HS256", []byte("secret"), nil)

	s.mailer = mailer.Mailer{
		Host:     s.conf.Mailer.Host,
		Port:     s.conf.Mailer.Port,
		Username: s.conf.Mailer.Username,
		Password: s.conf.Mailer.Password,
		From:     s.conf.Mailer.From,
	}

	return s
}

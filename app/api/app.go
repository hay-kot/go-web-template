package main

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/ent"
	"github.com/hay-kot/git-web-template/internal/config"
	"github.com/hay-kot/git-web-template/pkgs/logger"
	"github.com/hay-kot/git-web-template/pkgs/mailer"
)

type app struct {
	Conf   *config.Config
	logger logger.SharedLogger
	mailer mailer.Mailer
	jwt    *jwtauth.JWTAuth
	db     *ent.Client
}

func NewApp(conf *config.Config) *app {
	s := &app{
		Conf: conf,
	}

	s.jwt = jwtauth.New("HS256", []byte("secret"), nil)

	s.mailer = mailer.Mailer{
		Host:     s.Conf.Mailer.Host,
		Port:     s.Conf.Mailer.Port,
		Username: s.Conf.Mailer.Username,
		Password: s.Conf.Mailer.Password,
		From:     s.Conf.Mailer.From,
	}

	return s
}

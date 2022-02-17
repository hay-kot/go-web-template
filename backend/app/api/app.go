package main

import (
	"github.com/hay-kot/git-web-template/backend/ent"
	"github.com/hay-kot/git-web-template/backend/internal/config"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/mailer"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

type app struct {
	conf   *config.Config
	logger *logger.Logger
	mailer mailer.Mailer
	db     *ent.Client
	server *server.Server
	repos  *repo.AllRepos
}

func NewApp(conf *config.Config) *app {
	s := &app{
		conf: conf,
	}

	s.mailer = mailer.Mailer{
		Host:     s.conf.Mailer.Host,
		Port:     s.conf.Mailer.Port,
		Username: s.conf.Mailer.Username,
		Password: s.conf.Mailer.Password,
		From:     s.conf.Mailer.From,
	}

	return s
}

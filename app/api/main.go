package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hay-kot/git-web-template/ent"
	"github.com/hay-kot/git-web-template/internal/config"
	"github.com/hay-kot/git-web-template/internal/repo"
	"github.com/hay-kot/git-web-template/pkgs/logger"
	"github.com/hay-kot/git-web-template/pkgs/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfgFile := "config.yml"

	cfg, err := config.NewConfig(cfgFile)
	if err != nil {
		panic(err)
	}
	if err := run(cfg); err != nil {
		panic(err)
	}
}

func run(cfg *config.Config) error {
	app := NewApp(cfg)

	// =========================================================================
	// Setup Logger

	var wrt io.Writer
	wrt = os.Stdout
	if app.Conf.Log.File != "" {
		f, err := os.OpenFile(app.Conf.Log.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)
		wrt = io.MultiWriter(wrt, f)
	}

	app.logger = logger.New(wrt, logger.LevelDebug)

	// =========================================================================
	// Initialize Database & Repos

	c, err := ent.Open(cfg.Database.GetDriver(), cfg.Database.GetUrl())
	if err != nil {
		app.logger.Fatal(err, logger.Props{
			"details":  "failed to connect to database",
			"database": cfg.Database.GetDriver(),
			"url":      cfg.Database.GetUrl(),
		})
	}
	defer func(c *ent.Client) {
		_ = c.Close()
	}(c)
	if err := c.Schema.Create(context.Background()); err != nil {
		app.logger.Fatal(err, logger.Props{
			"details": "failed to create schema",
		})
	}

	app.db = c

	repos := &repo.AllRepos{
		Users: repo.NewUserRepositoryEnt(app.db),
	}

	// =========================================================================
	// Start Server

	app.Conf.Print()

	server := server.Server{
		Port: cfg.Web.Port,
		Host: cfg.Web.Host,
	}

	routes := app.newRouter(repos)
	app.LogRoutes(routes)

	app.logger.Info("Starting HTTP Server", logger.Props{
		"host": server.Host,
		"port": server.Port,
	})

	return server.Start(routes)
}

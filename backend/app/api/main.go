package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hay-kot/git-web-template/backend/ent"
	"github.com/hay-kot/git-web-template/backend/internal/config"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
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
	if app.conf.Log.File != "" {
		f, err := os.OpenFile(app.conf.Log.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
	app.repos = repo.EntAllRepos(c)

	// =========================================================================
	// Start Server

	app.conf.Print()

	app.server = server.NewServer(app.conf.Web.Host, app.conf.Web.Port)

	routes := app.newRouter(app.repos)
	app.LogRoutes(routes)

	app.EnsureAdministrator()
	app.SeedDatabase(app.repos)

	app.logger.Info("Starting HTTP Server", logger.Props{
		"host": app.server.Host,
		"port": app.server.Port,
	})

	return app.server.Start(routes)
}

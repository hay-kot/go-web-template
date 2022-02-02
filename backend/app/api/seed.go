package main

import (
	"context"

	"github.com/hay-kot/git-web-template/backend/internal/dtos"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
)

func (a *app) SeedDatabase(repos *repo.AllRepos) {
	if !a.conf.Seed.Enabled {
		return
	}

	for _, user := range a.conf.Seed.Users {

		hashedPw, err := hasher.HashPassword(user.Password)

		if err != nil {
			a.logger.Error(err, logger.Props{
				"details": "failed to hash password",
				"user":    user.Name,
			})
		}

		// Check if User Exists
		usr, err := repos.Users.GetOneEmail(user.Email, context.Background())

		if err != nil {
			a.logger.Error(err, logger.Props{
				"details": "failed to get user during seed",
				"user":    user.Name,
			})
		}

		if usr.Id != 0 {
			a.logger.Info("seed user already exists", logger.Props{
				"user": user.Name,
			})
			continue
		}

		_, err = repos.Users.Create(&dtos.UserCreate{
			Name:        user.Name,
			Email:       user.Email,
			IsSuperuser: user.IsSuperuser,
			Password:    hashedPw,
		}, context.Background())

		if err != nil {
			a.logger.Error(err, logger.Props{
				"details": "failed to create seed user",
				"name":    user.Name,
			})
		}

		a.logger.Info("creating seed user", logger.Props{
			"name": user.Name,
		})
	}
}

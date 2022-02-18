package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/dtos"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
)

const (
	DefaultName     = "Admin"
	DefaultEmail    = "admin@admin.com"
	DefaultPassword = "admin"
)

// EnsureAdministartor ensures that there is at least one superuser in the database
// if one isn't found a default is generate using the default credentials
func (a *app) EnsureAdministrator() {
	superusers, err := a.repos.Users.GetSuperusers(context.Background())

	if err != nil {
		a.logger.Error(err, nil)
	}

	if len(superusers) > 0 {
		return
	}

	pw, _ := hasher.HashPassword(DefaultPassword)

	newSuperUser := dtos.UserCreate{
		Name:        DefaultName,
		Email:       DefaultEmail,
		IsSuperuser: true,
		Password:    pw,
	}

	a.logger.Info("creating default superuser", logger.Props{
		"name":  newSuperUser.Name,
		"email": newSuperUser.Email,
	})

	_, err = a.repos.Users.Create(&newSuperUser, context.Background())

	if err != nil {
		a.logger.Fatal(err, nil)
	}

}

func (a *app) SeedDatabase(repos *repo.AllRepos) {
	if !a.conf.Seed.Enabled {
		return
	}

	for _, user := range a.conf.Seed.Users {

		// Check if User Exists
		usr, _ := repos.Users.GetOneEmail(user.Email, context.Background())

		if usr.Id != uuid.Nil {
			a.logger.Info("seed user already exists", logger.Props{
				"user": user.Name,
			})
			continue
		}

		hashedPw, err := hasher.HashPassword(user.Password)

		if err != nil {
			a.logger.Error(err, logger.Props{
				"details": "failed to hash password",
				"user":    user.Name,
			})
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

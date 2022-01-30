package mocks

import (
	"context"

	"github.com/hay-kot/git-web-template/ent"
	"github.com/hay-kot/git-web-template/internal/repo"
	_ "github.com/mattn/go-sqlite3"
)

func GetEntRepos() (*repo.AllRepos, func() error) {
	c, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}

	if err := c.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	return &repo.AllRepos{Users: repo.NewUserRepositoryEnt(c)}, c.Close

}

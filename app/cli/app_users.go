package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/hay-kot/git-web-template/internal/repo"
	"github.com/hay-kot/git-web-template/pkgs/hasher"
	"github.com/urfave/cli/v2"
)

func (a *app) UserCreate(c *cli.Context) error {
	// Get Flags
	name := c.String("name")
	password := c.String("password")
	email := c.String("email")
	isSuper := c.Bool("is-super")

	fmt.Println("Creating Superuser")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Email: %s\n", email)

	pwHash, err := hasher.HashPassword(password)
	if err != nil {
		return err
	}

	usr := &repo.UserCreate{
		Name:        name,
		Email:       email,
		Password:    pwHash,
		IsSuperuser: isSuper,
	}

	_, err = a.repos.Users.Create(usr, context.Background())

	if err == nil {
		fmt.Println("Super user created")
	}
	return err
}

func (a *app) UserDelete(c *cli.Context) error {
	// Get Flags
	id := c.Int("id")

	fmt.Printf("Deleting user with id: %d\n", id)

	// Confirm Action
	fmt.Printf("Are you sure you want to delete this user? (y/n) ")
	var answer string
	_, err := fmt.Scanln(&answer)
	if answer != "y" || err != nil {
		fmt.Println("Aborting")
		return nil
	}

	err = a.repos.Users.Delete(id, context.Background())

	if err == nil {
		fmt.Printf("%v User(s) deleted (id=%v)\n", 1, id)
	}
	return err
}

func (a *app) UserList(c *cli.Context) error {
	fmt.Println("Superuser List")

	users, err := a.repos.Users.GetAll(context.Background())

	if err != nil {
		return err
	}

	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer func(tabWriter *tabwriter.Writer) {
		_ = tabWriter.Flush()
	}(tabWriter)

	_, err = fmt.Fprintln(tabWriter, "Id\tName\tEmail\tIsSuper")

	if err != nil {
		return err
	}

	for _, u := range users {
		_, _ = fmt.Fprintf(tabWriter, "%v\t%s\t%s\t%v\n", u.Id, u.Name, u.Email, u.IsSuperuser)
	}

	return nil
}

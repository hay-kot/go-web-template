package main

import (
	"context"
	"fmt"
	"github.com/hay-kot/git-web-template/ent/user"
	"github.com/hay-kot/git-web-template/pkgs/hasher"
	"github.com/urfave/cli/v2"
	"os"
	"text/tabwriter"
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

	usr, err := a.db.User.
		Create().
		SetName(name).
		SetEmail(email).
		SetPassword(pwHash).
		SetIsSuperuser(isSuper).
		Save(context.Background())

	if err == nil {
		fmt.Printf("Super user created: %v\n", usr.ID)
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

	numDeleted, err := a.db.User.
		Delete().
		Where(user.ID(id)).Exec(context.Background())

	if err == nil {
		fmt.Printf("%v User(s) deleted (id=%v)\n", numDeleted, id)
	}
	return err
}

func (a *app) UserList(c *cli.Context) error {
	fmt.Println("Superuser List")
	users, err := a.db.User.Query().Where(user.IsSuperuser(true)).All(context.Background())
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
		_, _ = fmt.Fprintf(tabWriter, "%v\t%s\t%s\t%v\n", u.ID, u.Name, u.Email, u.IsSuperuser)
	}

	return nil
}

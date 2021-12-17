package repo

import (
	"errors"

	"github.com/hay-kot/git-web-template/ent"
)

var (
	ErrNameEmpty  = errors.New("name is empty")
	ErrEmailEmpty = errors.New("email is empty")
)

type UserCreate struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsSuperuser bool   `json:"isSuperuser"`
}

func (u *UserCreate) Validate() error {
	if u.Name == "" {
		return ErrNameEmpty
	}
	if u.Email == "" {
		return ErrEmailEmpty
	}
	return nil
}

type UserOut struct {
	ID       int    `json:"id"`
	Password string `json:"-"`
	UserCreate
}

func entToUserOut(usr *UserOut, entUsr *ent.User) {
	usr.ID = entUsr.ID
	usr.Password = entUsr.Password
	usr.UserCreate = UserCreate{
		Name:        entUsr.Name,
		Email:       entUsr.Email,
		IsSuperuser: entUsr.IsSuperuser,
	}
}

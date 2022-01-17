package repo

import (
	"errors"
)

var (
	ErrNameEmpty  = errors.New("name is empty")
	ErrEmailEmpty = errors.New("email is empty")
)

type UserCreate struct {
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
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
	Id          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password"`
	IsSuperuser bool   `json:"isSuperuser"`
}

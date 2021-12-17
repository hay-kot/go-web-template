package repo

import (
	"context"

	"github.com/hay-kot/git-web-template/ent"
	"github.com/hay-kot/git-web-template/ent/user"
)

type EntUserRepository struct {
	db *ent.Client
}

// NewEntUserRepository returns a new instance of the EntUserRepository that relies on the given *ent.Client.
func NewUserRepositoryEnt(db *ent.Client) *EntUserRepository {
	return &EntUserRepository{db: db}
}

func (e *EntUserRepository) GetOneId(id int, ctx context.Context) (UserOut, error) {
	usr, err := e.db.User.Query().Where(user.ID(id)).Only(ctx)

	usrOut := UserOut{}

	if err != nil {
		return usrOut, err
	}

	entToUserOut(&usrOut, usr)

	return usrOut, nil
}

func (e *EntUserRepository) GetOneEmail(email string, ctx context.Context) (UserOut, error) {
	usr, err := e.db.User.Query().Where(user.Email(email)).Only(ctx)

	usrOut := UserOut{}

	if err != nil {
		return usrOut, err
	}

	entToUserOut(&usrOut, usr)

	return usrOut, nil
}

func (e *EntUserRepository) GetAll(ctx context.Context) ([]UserOut, error) {
	users, err := e.db.User.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	var usrs []UserOut

	for _, usr := range users {
		usrs = append(usrs, UserOut{
			ID: usr.ID,
			UserCreate: UserCreate{
				Name:  usr.Name,
				Email: usr.Email,
			},
		})
	}

	return usrs, nil
}

func (e *EntUserRepository) Create(usr *UserCreate, ctx context.Context) error {
	err := usr.Validate()
	if err != nil {
		return err
	}

	_, err = e.db.User.
		Create().
		SetName(usr.Name).
		SetEmail(usr.Email).
		SetIsSuperuser(usr.IsSuperuser).
		Save(ctx)

	return err
}

func (e *EntUserRepository) Update(user *UserCreate, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (e *EntUserRepository) Delete(id int, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

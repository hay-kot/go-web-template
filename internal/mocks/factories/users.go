package factories

import (
	"github.com/hay-kot/git-web-template/internal/dtos"
	"github.com/hay-kot/git-web-template/pkgs/faker"
)

func UserFactory() dtos.UserCreate {
	f := faker.NewFaker()
	return dtos.UserCreate{
		Name:        f.RandomString(10),
		Email:       f.RandomEmail(),
		Password:    f.RandomString(10),
		IsSuperuser: f.RandomBool(),
	}
}
package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/hay-kot/git-web-template/pkgs/faker"
	"github.com/stretchr/testify/assert"
)

func UserFactory() UserCreate {
	return UserCreate{
		Name:        faker.GetRandomString(10),
		Email:       faker.GetRandomEmail(),
		Password:    faker.GetRandomString(10),
		IsSuperuser: faker.GetRandomBool(),
	}
}

func Test_EntUserRepo_GetOneEmail(t *testing.T) {
	assert := assert.New(t)
	user := UserFactory()
	ctx := context.Background()

	testRepos.Users.Create(&user, ctx)

	foundUser, err := testRepos.Users.GetOneEmail(user.Email, ctx)

	assert.NotNil(foundUser)
	assert.Nil(err)
	assert.Equal(user.Email, foundUser.Email)
	assert.Equal(user.Name, foundUser.Name)

	// Cleanup
	testRepos.Users.Delete(foundUser.Id, ctx)
}

func Test_EntUserRepo_GetOneId(t *testing.T) {
	assert := assert.New(t)
	user := UserFactory()
	ctx := context.Background()

	userOut, _ := testRepos.Users.Create(&user, ctx)
	foundUser, err := testRepos.Users.GetOneId(userOut.Id, ctx)

	assert.NotNil(foundUser)
	assert.Nil(err)
	assert.Equal(user.Email, foundUser.Email)
	assert.Equal(user.Name, foundUser.Name)

	// Cleanup
	testRepos.Users.Delete(userOut.Id, ctx)
}

func Test_EntUserRepo_GetAll(t *testing.T) {
	// Setup
	toCreate := []UserCreate{
		UserFactory(),
		UserFactory(),
		UserFactory(),
		UserFactory(),
	}

	ctx := context.Background()

	created := []UserOut{}

	for _, usr := range toCreate {
		usrOut, _ := testRepos.Users.Create(&usr, ctx)
		created = append(created, usrOut)
	}

	// Validate
	allUsers, err := testRepos.Users.GetAll(ctx)

	assert.Nil(t, err)
	assert.Equal(t, len(created), len(allUsers))

	for _, usr := range created {
		fmt.Printf("%+v\n", usr)
		assert.Contains(t, allUsers, usr)
	}

	for _, usr := range created {
		testRepos.Users.Delete(usr.Id, ctx)
	}
}

func Test_EntUserRepo_Update(t *testing.T) {
	t.Skip()
}

package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/hay-kot/git-web-template/internal/dtos"
	"github.com/hay-kot/git-web-template/pkgs/faker"
	"github.com/stretchr/testify/assert"
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
	toCreate := []dtos.UserCreate{
		UserFactory(),
		UserFactory(),
		UserFactory(),
		UserFactory(),
	}

	ctx := context.Background()

	created := []dtos.UserOut{}

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

func Test_EntUserRepo_Delete(t *testing.T) {
	// Create 10 Users
	for i := 0; i < 10; i++ {
		user := UserFactory()
		ctx := context.Background()
		_, _ = testRepos.Users.Create(&user, ctx)
	}

	// Delete all
	ctx := context.Background()
	allUsers, _ := testRepos.Users.GetAll(ctx)

	assert.Greater(t, len(allUsers), 0)
	testRepos.Users.DeleteAll(ctx)

	allUsers, _ = testRepos.Users.GetAll(ctx)
	assert.Equal(t, len(allUsers), 0)

}

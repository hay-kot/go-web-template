package v1

import (
	"context"
	"testing"

	"github.com/hay-kot/git-web-template/backend/internal/dtos"
	"github.com/hay-kot/git-web-template/backend/internal/mocks"
	"github.com/hay-kot/git-web-template/backend/internal/mocks/factories"
)

var mockHandler = &Handlersv1{}
var users = []dtos.UserOut{}

func userPool() func() {
	create := []dtos.UserCreate{
		factories.UserFactory(),
		factories.UserFactory(),
		factories.UserFactory(),
		factories.UserFactory(),
	}

	userOut := []dtos.UserOut{}

	for _, user := range create {
		usrOut, _ := mockHandler.repos.Users.Create(&user, context.Background())
		userOut = append(userOut, usrOut)
	}

	users = userOut

	purge := func() {
		mockHandler.repos.Users.DeleteAll(context.Background())
	}

	return purge
}

func TestMain(m *testing.M) {
	// Set Handler Vars
	mockHandler.log = mocks.GetStructLogger()
	mockHandler.jwt = mocks.GetJWTAuth()
	repos, closeDb := mocks.GetEntRepos()
	mockHandler.repos = repos

	defer closeDb()

	purge := userPool()
	defer purge()

	m.Run()
}

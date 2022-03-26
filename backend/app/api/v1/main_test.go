package v1

import (
	"context"
	"testing"

	"github.com/hay-kot/git-web-template/backend/internal/mocks"
	"github.com/hay-kot/git-web-template/backend/internal/mocks/factories"
	"github.com/hay-kot/git-web-template/backend/internal/types"
)

var mockHandler = &Handlersv1{}
var users = []types.UserOut{}

func userPool() func() {
	create := []types.UserCreate{
		factories.UserFactory(),
		factories.UserFactory(),
		factories.UserFactory(),
		factories.UserFactory(),
	}

	userOut := []types.UserOut{}

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
	repos, closeDb := mocks.GetEntRepos()
	mockHandler.repos = repos
	mockHandler.services = mocks.GetMockServices(repos)

	defer closeDb()

	purge := userPool()
	defer purge()

	m.Run()
}

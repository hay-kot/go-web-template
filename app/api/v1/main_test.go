package v1

import (
	"testing"

	"github.com/hay-kot/git-web-template/internal/mocks"
)

var mockHandler = &Handlersv1{}

func TestMain(m *testing.M) {
	// Set Handler Vars
	mockHandler.log = mocks.GetConsoleLogger()
	mockHandler.jwt = mocks.GetJWTAuth()
	repos, closeDb := mocks.GetEntRepos()
	mockHandler.repos = repos

	defer closeDb()

	m.Run()
}

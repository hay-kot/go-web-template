package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hay-kot/git-web-template/internal/dtos"
	"github.com/hay-kot/git-web-template/internal/mocks/factories"
	"github.com/stretchr/testify/assert"
)

type usersResponse struct {
	Users []dtos.UserCreate `json:"users"`
}

type userResponse struct {
	User dtos.UserCreate `json:"user"`
}

func userPool(t *testing.T) ([]dtos.UserOut, func()) {
	users := []dtos.UserCreate{
		factories.UserFactory(),
		factories.UserFactory(),
		factories.UserFactory(),
		factories.UserFactory(),
	}

	userOut := []dtos.UserOut{}

	for _, user := range users {
		usrOut, _ := mockHandler.repos.Users.Create(&user, context.Background())
		userOut = append(userOut, usrOut)
	}

	purge := func() {
		mockHandler.repos.Users.DeleteAll(context.Background())
	}

	return userOut, purge
}

func purgeUsers() {
	mockHandler.repos.Users.DeleteAll(context.Background())
}

func Test_HandleAdminUserGetAll_Success(t *testing.T) {
	purgeUsers()
	users, purge := userPool(t)
	defer purge()

	r := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/admin/users", nil)

	mockHandler.HandleAdminUserGetAll()(r, req)

	response := usersResponse{
		Users: []dtos.UserCreate{},
	}

	_ = json.Unmarshal(r.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, len(users), len(response.Users))

	knowEmail := []string{
		users[0].Email,
		users[1].Email,
		users[2].Email,
		users[3].Email,
	}

	for _, user := range users {
		assert.Contains(t, knowEmail, user.Email)
	}

}

func Test_HandleAdminUserGet_Success(t *testing.T) {
	purgeUsers()
	users, purge := userPool(t)
	defer purge()

	targetUser := users[2]
	r := httptest.NewRecorder()
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/admin/users/%v", targetUser.Id), nil)

	mockHandler.HandleAdminUserGet()(r, req)

	response := userResponse{
		User: dtos.UserCreate{},
	}

	_ = json.Unmarshal(r.Body.Bytes(), &response)
	assert.ObjectsAreEqual(targetUser, response.User)

}

func Test_HandleAdminUserCreate_Success(t *testing.T) {
	t.Skip()
}

func Test_HandleAdminUserUpdate_Success(t *testing.T) {
	t.Skip()
}

func Test_HandleAdminUserUpdate_Delete(t *testing.T) {
	t.Skip()
}

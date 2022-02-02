package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hay-kot/git-web-template/internal/dtos"
	"github.com/hay-kot/git-web-template/internal/mocks/chimocker"
	"github.com/hay-kot/git-web-template/internal/mocks/factories"
	"github.com/hay-kot/git-web-template/pkgs/server"
	"github.com/stretchr/testify/assert"
)

const (
	UrlUser      = "/api/v1/admin/users"
	UrlUserId    = "/api/v1/admin/users/%v"
	UrlUserIdChi = "/api/v1/admin/users/{id}"
)

type usersResponse struct {
	Users []dtos.UserOut `json:"users"`
}

type userResponse struct {
	User dtos.UserOut `json:"user"`
}

func Test_HandleAdminUserGetAll_Success(t *testing.T) {
	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, UrlUser, nil)

	mockHandler.HandleAdminUserGetAll()(r, req)

	response := usersResponse{
		Users: []dtos.UserOut{},
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
	targetUser := users[2]
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(UrlUserId, targetUser.Id), nil)

	req = chimocker.WithUrlParam(req, "id", fmt.Sprintf("%v", targetUser.Id))

	mockHandler.HandleAdminUserGet()(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	response := userResponse{
		User: dtos.UserOut{},
	}

	_ = json.Unmarshal(res.Body.Bytes(), &response)
	assert.Equal(t, targetUser.Id, response.User.Id)
}

func Test_HandleAdminUserCreate_Success(t *testing.T) {
	payload := factories.UserFactory()

	r := httptest.NewRecorder()

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, UrlUser, bytes.NewBuffer(body))
	req.Header.Set(server.ContentType, server.ContentJSON)

	mockHandler.HandleAdminUserCreate()(r, req)

	assert.Equal(t, http.StatusCreated, r.Code)

	usr, err := mockHandler.repos.Users.GetOneEmail(payload.Email, context.Background())

	assert.NoError(t, err)
	assert.Equal(t, payload.Email, usr.Email)
	assert.Equal(t, payload.Name, usr.Name)
	assert.NotEqual(t, payload.Password, usr.Password) // smoke test - check password is hashed

	_ = mockHandler.repos.Users.Delete(usr.Id, context.Background())
}

func Test_HandleAdminUserUpdate_Success(t *testing.T) {
	t.Skip()
}

func Test_HandleAdminUserUpdate_Delete(t *testing.T) {
	t.Skip()
}

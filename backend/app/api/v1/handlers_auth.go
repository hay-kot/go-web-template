package v1

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/dtos"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (h *Handlersv1) createToken(userId uuid.UUID, ctx context.Context) (string, error) {
	newToken := hasher.GenerateToken()

	token, err := h.repos.AuthTokens.CreateToken(dtos.UserAuthTokenCreate{
		UserId:    userId,
		TokenHash: newToken.Hash,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}, ctx)

	if err != nil {
		return "", err
	}

	if token.TokenHash == nil {
		return "", errors.New("token is empty")
	}

	return "Bearer " + newToken.Raw, err
}

// handleAuthLogin returns a handler to handle username/password authentication for users of the API.
func (h *Handlersv1) HandleAuthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginForm := LoginForm{}
		err := server.Decode(r, &loginForm)

		if err != nil {
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		usr, err := h.repos.Users.GetOneEmail(loginForm.Username, r.Context())

		if err != nil || !hasher.CheckPasswordHash(loginForm.Password, usr.Password) {
			server.RespondError(w, http.StatusUnauthorized, errors.New("invalid username or password"))
			return
		}

		bearer, _ := h.createToken(usr.Id, r.Context())

		err = server.Respond(w, http.StatusOK, TokenResponse{
			Token: bearer,
		})

		if err != nil {
			h.log.Error(err, logger.Props{
				"user": usr.Email,
			})

			return
		}
	}
}

// handleAuthRefresh returns a handler that will issue a new token from an existing token.
// This does not validate that the user still exists within the database.
func (h *Handlersv1) HandleAuthRefresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		currentToken := strings.TrimPrefix(bearer, "Bearer ")

		hash := hasher.HashToken(currentToken)

		usr, err := h.repos.AuthTokens.GetUserFromToken(hash, r.Context())

		newToken, _ := h.createToken(usr.Id, r.Context())

		if err != nil {
			server.RespondUnauthorized(w)
			return
		}

		err = server.Respond(w, http.StatusOK, TokenResponse{
			Token: newToken,
		})

		if err != nil {
			return
		}
	}
}

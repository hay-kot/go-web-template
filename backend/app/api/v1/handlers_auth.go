package v1

import (
	"errors"
	"net/http"
	"time"

	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	BearerToken string    `json:"token"`
	ExpiresAt   time.Time `json:"expiresAt"`
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

		newToken, err := h.services.Auth.Login(r.Context(), loginForm.Username, loginForm.Password)

		err = server.Respond(w, http.StatusOK, TokenResponse{
			BearerToken: "Bearer " + newToken.Raw,
			ExpiresAt:   newToken.ExpiresAt,
		})

		if err != nil {
			h.log.Error(err, logger.Props{
				"user": loginForm.Username,
			})

			return
		}
	}
}

func (h *Handlersv1) HandleAuthLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := services.GetUserTokenFromContext(r.Context())

		if token == "" {
			server.RespondError(w, http.StatusUnauthorized, errors.New("no token within request context"))
			return
		}

		err := h.services.Auth.Logout(r.Context(), token)

		if err != nil {
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		err = server.Respond(w, http.StatusOK, nil)
	}
}

// handleAuthRefresh returns a handler that will issue a new token from an existing token.
// This does not validate that the user still exists within the database.
func (h *Handlersv1) HandleAuthRefresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestToken := services.GetUserTokenFromContext(r.Context())

		if requestToken == "" {
			server.RespondError(w, http.StatusUnauthorized, errors.New("no user token found"))
			return
		}

		newToken, err := h.services.Auth.RenewToken(r.Context(), requestToken)

		if err != nil {
			server.RespondUnauthorized(w)
			return
		}

		err = server.Respond(w, http.StatusOK, newToken)

		if err != nil {
			return
		}
	}
}

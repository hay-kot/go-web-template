package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func issueToken(auth *jwtauth.JWTAuth, username string) (string, error) {
	_, token, err := auth.Encode(map[string]interface{}{
		"username": username,
	})

	if err != nil {
		return "", err
	}

	return "Bearer " + token, err
}

// handleAuthLogin returns a handler to handle username/password authentication for users of the API.
func (s *Handlersv1) HandleAuthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginForm := LoginForm{}
		err := server.Decode(r, &loginForm)

		if err != nil {
			return
		}

		usr, err := s.repos.Users.GetOneEmail(loginForm.Username, r.Context())

		if err != nil || !hasher.CheckPasswordHash(loginForm.Password, usr.Password) {
			server.RespondError(w, http.StatusUnauthorized, errors.New("invalid username or password"))
			return
		}

		bearer, _ := issueToken(s.jwt, usr.Email)

		err = server.Respond(w, http.StatusOK, TokenResponse{
			Token: bearer,
		})

		if err != nil {
			return
		}
	}
}

// handleAuthRefresh returns a handler that will issue a new JWT from an existing JWT token.
// This does not validate that the user still exists within the database.
func (s *Handlersv1) HandleAuthRefresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		username := claims["username"].(string)

		bearer, _ := issueToken(s.jwt, username)

		err := server.Respond(w, http.StatusOK, TokenResponse{
			Token: bearer,
		})

		if err != nil {
			return
		}
	}
}

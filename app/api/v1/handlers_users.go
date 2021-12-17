package v1

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/pkgs/server"
)

func (s *Handlersv1) HandleUserSelf() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		username := claims["username"].(string)

		usr, err := s.repos.Users.GetOneEmail(username, r.Context())

		if err != nil {
			_ = server.RespondError(w, http.StatusInternalServerError, err)
		}

		// Return Username
		_ = server.Respond(w, http.StatusOK, usr)
	}
}

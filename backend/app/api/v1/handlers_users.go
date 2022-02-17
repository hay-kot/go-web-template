package v1

import (
	"errors"
	"net/http"

	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

func (s *Handlersv1) HandleUserSelf() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr := services.UserFromContext(r.Context())

		if usr == nil {
			s.log.Error(errors.New("no user within request context"), nil)
			server.RespondInternalServerError(w)
			return
		}

		_ = server.Respond(w, http.StatusOK, usr)
	}
}

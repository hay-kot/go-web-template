package v1

import (
	"net/http"

	"github.com/hay-kot/git-web-template/backend/internal/dtos"
	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

func (s *Handlersv1) HandleUserSelf() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr := r.Context().Value(services.ContextUser).(dtos.UserOut)

		// if usr == dtos.UserOut{} {
		// 	s.log.Error(errors.New("no user within request context"), nil)
		// 	server.RespondInternalServerError(w)
		// }

		// Return Username
		_ = server.Respond(w, http.StatusOK, usr)
	}
}

package v1

import (
	"errors"
	"net/http"

	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

func (ctrl *V1Controller) HandleUserSelf() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := services.GetUserTokenFromContext(r.Context())

		usr, err := ctrl.svc.User.GetSelf(r.Context(), token)

		if usr.IsNull() || err != nil {
			ctrl.log.Error(errors.New("no user within request context"), nil)
			server.RespondInternalServerError(w)
			return
		}

		_ = server.Respond(w, http.StatusOK, usr)
	}
}

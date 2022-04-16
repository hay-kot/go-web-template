package base

import (
	"net/http"

	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

type ReadyFunc func() bool

type BaseController struct {
	log *logger.Logger
	svr *server.Server
}

func NewBaseController(log *logger.Logger, svr *server.Server) *BaseController {
	h := &BaseController{
		log: log,
		svr: svr,
	}
	return h
}

func (ctrl *BaseController) HandleBase(versions ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := types.ApiSummary{
			Healthy:  true,
			Versions: versions,
			Title:    "Go API Template",
			Message:  "Welcome to the Go API Template Application!",
		}

		err := server.Respond(w, http.StatusOK, data)

		if err != nil {
			ctrl.log.Error(err, nil)
			server.RespondInternalServerError(w)
		}
	}
}

func (ctrl *BaseController) HandleReady(ready ReadyFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ready() {
			server.Respond(w, http.StatusOK, server.
				Wrap("status", "available").
				Message("The service is ready to use"),
			)
		} else {
			server.Respond(w, http.StatusServiceUnavailable, server.
				Wrap("status", "unavailable").
				Message("The service is not ready to use"),
			)
		}
	}
}

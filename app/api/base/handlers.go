package base

import (
	"net/http"

	"github.com/hay-kot/git-web-template/pkgs/logger"
	"github.com/hay-kot/git-web-template/pkgs/server"
)

type Handlersv1 struct {
	log *logger.Logger
	svr *server.Server
}

func NewHandlerV1(log *logger.Logger, svr *server.Server) *Handlersv1 {
	h := &Handlersv1{
		log: log,
		svr: svr,
	}
	return h
}

type BaseRouteResponse struct {
	Healthy  bool     `json:"health,omitempty"`
	Versions []string `json:"versions,omitempty"`
	Title    string   `json:"title,omitempty"`
	Message  string   `json:"message,omitempty"`
}

func (h *Handlersv1) HandleBase(versions ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := BaseRouteResponse{
			Healthy:  true,
			Versions: versions,
			Title:    "Go API Template",
			Message:  "Welcome to the Go API Template Application!",
		}

		err := server.Respond(w, http.StatusOK, data)

		if err != nil {
			h.log.Error(err, nil)
			server.RespondInternalServerError(w)
		}
	}
}

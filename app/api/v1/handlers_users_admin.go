package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hay-kot/git-web-template/pkgs/server"
)

func (s *Handlersv1) HandleAdminUserGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.repos.Users.GetAll(r.Context())

		if err != nil {
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusOK, server.Wrap("users", users))
	}
}

func (s *Handlersv1) HandleAdminUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		user, err := s.repos.Users.GetOneId(id, r.Context())

		if err != nil {
			s.log.Error(err, nil)
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}
		server.Respond(w, http.StatusOK, server.Wrap("user", user))

	}
}

func (s *Handlersv1) HandleAdminUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *Handlersv1) HandleAdminUserUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Handlersv1) HandleAdminUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

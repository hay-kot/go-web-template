package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

func (ctrl *V1Controller) HandleAdminUserGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := ctrl.svc.Admin.GetAll(r.Context())

		if err != nil {
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusOK, server.Wrap("users", users))
	}
}

func (ctrl *V1Controller) HandleAdminUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := uuid.Parse(chi.URLParam(r, "id"))

		if err != nil {
			ctrl.log.Debug(err.Error(), logger.Props{
				"scope":   "admin",
				"details": "failed to convert id to valid UUID",
			})
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		user, err := ctrl.svc.Admin.GetByID(r.Context(), uid)

		if err != nil {
			ctrl.log.Error(err, nil)
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}
		server.Respond(w, http.StatusOK, server.Wrap("user", user))

	}
}

func (ctrl *V1Controller) HandleAdminUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createData := types.UserCreate{}

		if err := server.Decode(r, &createData); err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to decode user create data",
			})
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		err := createData.Validate()

		if err != nil {
			server.RespondError(w, http.StatusUnprocessableEntity, err)
			return
		}

		hashedPw, err := hasher.HashPassword(createData.Password)

		if err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to hash password",
			})

			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		createData.Password = hashedPw
		userOut, err := ctrl.svc.Admin.Create(r.Context(), &createData)

		if err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to create user",
			})

			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusCreated, server.Wrap("user", userOut))
	}
}

func (ctrl *V1Controller) HandleAdminUserUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (ctrl *V1Controller) HandleAdminUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

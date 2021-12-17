package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/app/api/base"
	v1 "github.com/hay-kot/git-web-template/app/api/v1"
	"github.com/hay-kot/git-web-template/internal/repo"
)

const prefix = "/api"

// registerRoutes registers all the routes for the API
func (s *app) newRouter(repos *repo.AllRepos) *chi.Mux {
	r := chi.NewRouter()
	setGlobalMiddleware(r)

	// =========================================================================
	// Base Routes

	// Server Favicon
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	baseHandler := base.NewHandlerV1(s.logger)
	r.Get(prefix, baseHandler.HandleBase("v1"))

	// =========================================================================
	// API Version 1

	v1Base, v1Handlers := v1.NewHandlerV1(prefix, repos, s.jwt, s.logger)
	r.Post(v1Base("/login"), v1Handlers.HandleAuthLogin())
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(s.jwt))
		r.Use(mwAuth)
		r.Get(v1Base("/users/self"), v1Handlers.HandleUserSelf())
	})

	r.Group(func(r chi.Router) {
		r.Use(mwAdmin)
		r.Get(v1Base("/admin/users"), v1Handlers.HandleAdminUserGetAll())
		r.Post(v1Base("/admin/users"), v1Handlers.HandleAdminUserCreate())
		r.Get(v1Base("/admin/users/:id"), v1Handlers.HandleAdminUserGet())
		r.Put(v1Base("/admin/users/:id"), v1Handlers.HandleAdminUserCreate())
		r.Delete(v1Base("/admin/users/:id"), v1Handlers.HandleAdminUserCreate())
	})

	return r
}

// LogRoutes logs the routes of the server that are registered within Server.registerRoutes(). This is useful for debugging.
// See https://github.com/go-chi/chi/issues/332 for details and inspiration.
func (s *app) LogRoutes(r *chi.Mux) {
	desiredSpaces := 10

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		text := "[" + method + "]"

		for len(text) < desiredSpaces {
			text = text + " "
		}

		s.logger.Debug("Registered: %s%s\n", text, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}

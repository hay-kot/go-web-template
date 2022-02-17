package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hay-kot/git-web-template/backend/app/api/base"
	v1 "github.com/hay-kot/git-web-template/backend/app/api/v1"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
)

const prefix = "/api"

// registerRoutes registers all the routes for the API
func (a *app) newRouter(repos *repo.AllRepos) *chi.Mux {
	r := chi.NewRouter()
	a.setGlobalMiddleware(r)

	// =========================================================================
	// Base Routes

	// Server Favicon
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	baseHandler := base.NewHandlerV1(a.logger, a.server)
	r.Get(prefix, baseHandler.HandleBase("v1"))
	r.Get(prefix+"/status", baseHandler.HandleReady(func() bool { return true }))

	// =========================================================================
	// API Version 1

	v1Base, v1Handlers := v1.NewHandlerV1(prefix, repos, a.logger)
	r.Post(v1Base("/login"), v1Handlers.HandleAuthLogin())
	r.Group(func(r chi.Router) {
		r.Use(a.mwAuthToken)
		r.Get(v1Base("/users/self"), v1Handlers.HandleUserSelf())
		r.Post(v1Base("/users/logout"), v1Handlers.HandleAuthLogout())
	})

	r.Group(func(r chi.Router) {
		r.Use(a.mwAuthToken)
		r.Get(v1Base("/admin/users"), v1Handlers.HandleAdminUserGetAll())
		r.Post(v1Base("/admin/users"), v1Handlers.HandleAdminUserCreate())
		r.Get(v1Base("/admin/users/{id}"), v1Handlers.HandleAdminUserGet())
		r.Put(v1Base("/admin/users/{id}"), v1Handlers.HandleAdminUserCreate())
		r.Delete(v1Base("/admin/users/{id}"), v1Handlers.HandleAdminUserCreate())
	})

	return r
}

// LogRoutes logs the routes of the server that are registered within Server.registerRoutes(). This is useful for debugging.
// See https://github.com/go-chi/chi/issues/332 for details and inspiration.
func (a *app) LogRoutes(r *chi.Mux) {
	desiredSpaces := 10

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		text := "[" + method + "]"

		for len(text) < desiredSpaces {
			text = text + " "
		}

		fmt.Printf("Registered Route: %s%s\n", text, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}

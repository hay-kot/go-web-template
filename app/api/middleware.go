package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/pkgs/server"
	"github.com/lestrrat-go/jwx/jwt"
)

func setGlobalMiddleware(r *chi.Mux) {
	// =========================================================================
	// Middleware
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
}

// mwAuth is a middleware that will check the JWT token
func mwAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			_ = server.RespondError(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			_ = server.RespondError(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

// mwAuth is a middleware that will check the JWT token
func mwAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			_ = server.RespondError(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			_ = server.RespondError(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/hay-kot/git-web-template/internal/config"
	"github.com/hay-kot/git-web-template/pkgs/logger"
	"github.com/hay-kot/git-web-template/pkgs/server"
	"github.com/lestrrat-go/jwx/jwt"
)

func (a *app) setGlobalMiddleware(r *chi.Mux) {
	// =========================================================================
	// Middleware
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(mwStripTrailingSlash)

	// Use struct logger in production for requests, but use
	// pretty console logger in development.
	if a.Conf.Mode == config.ModeDevelopment {
		r.Use(middleware.Logger)
	} else {
		r.Use(a.mwStructLogger)
	}
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

// mqStripTrailingSlash is a middleware that will strip trailing slashes from the request path.
func mwStripTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

func (a *app) mwStructLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		url := fmt.Sprintf("%s://%s%s %s", scheme, r.Host, r.RequestURI, r.Proto)

		a.logger.Info(fmt.Sprintf("[%s] %s", r.Method, url), logger.Props{
			"id":     middleware.GetReqID(r.Context()),
			"method": r.Method,
			"url":    url,
			"remote": r.RemoteAddr,
		})

		// Do stuff here
		next.ServeHTTP(w, r)
	})
}

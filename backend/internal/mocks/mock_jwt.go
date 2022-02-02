package mocks

import "github.com/go-chi/jwtauth/v5"

func GetJWTAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte("secret"), nil)
}

package provider

import (
	"github.com/go-chi/jwtauth"
)

func NewJWTAuth(env *EnvProvider) *jwtauth.JWTAuth {
	auth := jwtauth.New("HS256", []byte(env.jwtSecret), nil)

	return auth
}

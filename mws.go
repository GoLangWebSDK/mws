package mws

import (
	"net/http"

	"github.com/GoLangWebSDK/mws/jwt"
	"github.com/GoLangWebSDK/mws/log"
)

func NewLog() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return log.Middleware(next)
	}
}

func NewJWT(validator jwt.Bearer, routes []string) func(next http.Handler) http.Handler {
	jwt.PublicRoutes = routes
	jwt.Validator = validator
	return func(next http.Handler) http.Handler {
		return jwt.Middleware(next)
	}
}

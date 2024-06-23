package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Bearer interface {
	Validate(token string) error
}

var Validator Bearer
var PublicRoutes = []string{
	"/login",
	"/register",
	"/password_recovery",
	"/password_recovery/{token}",
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if routeIsPublic(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
		}

		bearerToken, err := getBearerToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if Validator == nil {
			fmt.Println("no validator")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = Validator.Validate(bearerToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func routeIsPublic(route string) bool {
	isPublic := false

	for _, publicRoute := range PublicRoutes {
		if strings.EqualFold(route, publicRoute) {
			isPublic = true
		}
	}

	return isPublic
}

func getBearerToken(r *http.Request) (string, error) {
	headerParts := strings.Split(r.Header.Get("Authorization"), " ")

	if len(headerParts) == 2 && strings.EqualFold(headerParts[0], "Bearer") {
		return headerParts[1], nil
	}

	return "", errors.New("invalid Authorization header")
}

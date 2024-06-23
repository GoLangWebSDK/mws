package main

import (
	"net/http"

	"github.com/GoLangWebSDK/mws"
	"github.com/GoLangWebSDK/rest"
)

func main() {

	router := rest.NewRouter()

	publicRoutes := []string{
		"/api/users/login",
		"/api/users/register",
		"/api/users/register/invite",
		"/api/users/password_recovery",
		"/api/users/password_recovery/{token}",
	}

	router.Use(mws.NewJWT(&User{}, publicRoutes))
	router.Use(mws.NewLog())

	ctrl := rest.NewController(router)

	ctrl.Get("/", func(session *rest.Session) {
		session.Response.WriteHeader(http.StatusOK)
		session.Response.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", router)
}

type User struct {
	Email    string
	Password string
}

func (u *User) Validate(token string) error {
	return nil
}

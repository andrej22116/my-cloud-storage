package controllers

import (
	"net/http"

	"./database"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	userArgs := database.UserArguments{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}

	userData, err := database.Registration(&gDatabase, userArgs)
	if err != nil {
		w.Header().
	}
	else {

	}
}

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

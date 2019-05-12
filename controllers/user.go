package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./database"
	"./filesystem"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	user := database.UserArguments{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
		w.Header().Add("Error", err.Error())
		return
	}

	err = database.Registration(gDatabase, user)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
	}

	filesystem.CreateFolder(filesystem.RootPath, user.Login)
}

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	user := database.UserArguments{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
		w.Header().Add("Error", err.Error())
		return
	}

	userData, err := database.Authorization(gDatabase, user)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
	} else {
		makeJsonHeader(w, "POST", userData)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token := database.Token{}
	err := jsonFromBody(r, &token)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		w.Header().Add("Error", err.Error())
		return
	}

	database.Logout(gDatabase, token.Token)
}

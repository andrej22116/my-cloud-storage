package main

import (
	"net/http"

	"./controllers"
	"github.com/gorilla/mux"
)

func main() {

	err := controllers.OpenDatabaseConnection()
	if err != nil {
		panic(err)
	}

	/*
		var user database.UserData
		user.Nickname = "admin"

		var newFile database.FileInfo
		newFile.Name = "test.txt"
		newFile.IsFolder = false
		err = database.CreateFile(db, user, "/admin", newFile)
		if err != nil {
			panic(err)
		}*/

	router := mux.NewRouter()

	///router.HandleFunc("/load", controllers.SendFile)
	//router.HandleFunc("/load", controllers.SendPrivateFile).Methods("POST")
	router.HandleFunc("/upload", controllers.UploadFileHandler).Methods("POST")
	router.HandleFunc("/files", controllers.GetAllFiles).Methods("GET")
	router.HandleFunc("/registration", controllers.RegistrationHandler).Methods("POST")
	router.HandleFunc("/authorization", controllers.AuthorizationHandler).Methods("POST")
	router.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./views/")))

	http.ListenAndServe(":8080", router)
}

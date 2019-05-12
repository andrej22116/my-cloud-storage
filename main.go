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

	router := mux.NewRouter()

	router.HandleFunc("/registration", controllers.RegistrationHandler).Methods("POST")
	router.HandleFunc("/authorization", controllers.AuthorizationHandler).Methods("POST")
	router.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")

	router.HandleFunc("/add/folder", controllers.CreateFolderHandler).Methods("POST")
	router.HandleFunc("/add/file", controllers.UploadFileHandler).Methods("POST")
	router.HandleFunc("/modify", controllers.ModifyFileHandler).Methods("POST")
	router.HandleFunc("/remove", controllers.RemoveFileHandler).Methods("POST")
	router.HandleFunc("/files", controllers.FileListHandler).Methods("POST")
	router.HandleFunc("/upload", controllers.BeforeUploadFileHandler).Methods("POST")
	router.HandleFunc("/upload/{uploadToken}", controllers.UploadFileHandler).Methods("POST")
	router.HandleFunc("/load", controllers.BeforeLoadFileHandler).Methods("POST")
	router.HandleFunc("/load/{loadToken}", controllers.LoadFileHandler).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./views/")))

	http.ListenAndServe(":8080", router)
}

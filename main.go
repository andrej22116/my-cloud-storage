package main

import (
	"fmt"
	"net/http"
	"os"

	"./controllers"
	"github.com/gorilla/mux"
)

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Command args: host port dbname user password")
		return
	}

	err := controllers.OpenDatabaseConnection(os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/registration", controllers.RegistrationHandler).Methods("POST")
	router.HandleFunc("/authorization", controllers.AuthorizationHandler).Methods("POST")
	router.HandleFunc("/testtoken", controllers.TestTokenHandler).Methods("POST")
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

	fmt.Println("Server started!")
	http.ListenAndServe(":8080", router)
}

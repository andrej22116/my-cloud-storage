package main

import (
	"fmt"

	"./controllers"
	"./controllers/database"
)

func main() {

	err := controllers.OpenDatabaseConnection()
	if err != nil {
		panic(err)
	}

	userData, err := database.Authorization(db, database.UserArguments{
		Login:    "admin",
		Password: "admin",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(userData.AccessToken)

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

	files, err := database.GetUserRoot(db, userData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lol")
	for _, file := range files {
		fmt.Println(file.Name)
	}

	//router := mux.NewRouter()

	///router.HandleFunc("/load", controllers.SendFile)
	//router.HandleFunc("/load", controllers.SendPrivateFile).Methods("POST")
	//router.HandleFunc("/files", controllers.GetAllFiles).Methods("GET")

	//http.ListenAndServe(":8080", router)
}

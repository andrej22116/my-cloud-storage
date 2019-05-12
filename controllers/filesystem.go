package controllers

import (
	"fmt"
	"net/http"

	"./database"
	"./filesystem"
)

/*
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, header, _ := r.FormFile("file")
	defer file.Close()

	err := filesystem.CreateFile(filesystem.RootPath, header.Filename, file)

	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
	}
}*/

func CreateFolderHandler(w http.ResponseWriter, r *http.Request) {
	var userArgs struct {
		Token string `json:"token"`
		Path  string `json:"path"`
		Name  string `json:"name"`
	}

	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	userData, err := checkAccess(userArgs.Token)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	if userArgs.Path == "/" {
		userArgs.Path = ""
	}

	databasePath := "/" + userData.Nickname + userArgs.Path
	fullPath := filesystem.RootPath + databasePath
	fmt.Println(userData.Nickname)
	fmt.Println(databasePath)
	fmt.Println(fullPath)
	fmt.Println(userArgs.Name)

	err = filesystem.CreateFolder(fullPath, userArgs.Name)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	err = database.CreateFile(gDatabase, databasePath, database.FileInfo{
		Name:     userArgs.Name,
		IsFolder: true,
	})

	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	fmt.Println("5")
}

func ModifyFileHandler(w http.ResponseWriter, r *http.Request) {
	var userArgs struct {
		Token   string `json:"token"`
		Path    string `json:"path"`
		OldName string `json:"oldName"`
		NewName string `json:"newName"`
	}

	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	userData, err := checkAccess(userArgs.Token)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	if userArgs.Path == "/" {
		userArgs.Path = ""
	}

	databasePath := "/" + userData.Nickname + userArgs.Path
	fullPath := filesystem.RootPath + databasePath

	err = filesystem.Rename(fullPath, userArgs.OldName, userArgs.NewName)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	err = database.ModifyFile(gDatabase, databasePath,
		database.FileInfo{
			Name: userArgs.OldName,
		},
		database.FileInfo{
			Name: userArgs.NewName,
		})

	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}
}

func RemoveFileHandler(w http.ResponseWriter, r *http.Request) {
	var userArgs struct {
		Token string `json:"token"`
		Path  string `json:"path"`
		Name  string `json:"name"`
	}

	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	userData, err := checkAccess(userArgs.Token)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	if userArgs.Path == "/" {
		userArgs.Path = ""
	}

	databasePath := "/" + userData.Nickname + userArgs.Path
	fullPath := filesystem.RootPath + databasePath

	err = filesystem.Remove(fullPath, userArgs.Name)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	err = database.RemoveFile(gDatabase, databasePath, database.FileInfo{
		Name: userArgs.Name,
	})

	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}
}

func FileListHandler(w http.ResponseWriter, r *http.Request) {
	var userArgs struct {
		Token string `json:"token"`
		Path  string `json:"path"`
	}

	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	userData, err := checkAccess(userArgs.Token)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	if userArgs.Path == "/" {
		userArgs.Path = ""
	}

	databasePath := "/" + userData.Nickname + userArgs.Path

	files, err := database.GetFileListFromPath(gDatabase, databasePath)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	makeJsonHeader(w, "POST", files)
}

package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"./database"
	"./filesystem"
)

/*
SendPrivateFile send file to owner user
*/
func BeforeLoadFileHandler(w http.ResponseWriter, r *http.Request) {
	var userArgs struct {
		Token string `json:"token"`
		Path  string `json:"path"`
		Name  string `json:"name"`
	}

	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)

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

	token, err := database.CreateNewUploadToken(gDatabase, userArgs.Token, databasePath, userArgs.Name)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	makeJsonHeader(w, "POST", map[string]string{
		"loadToken": strings.TrimPrefix(token, "\\"),
	})
}

func LoadFileHandler(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	loadToken := "\\" + variables["loadToken"]

	_, filePath, fileName, err := database.DataByUploadToken(gDatabase, loadToken)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	err = database.DeleteUploadToken(gDatabase, loadToken)
	if err != nil {
		fmt.Println(err)
	}

	fullPath := filesystem.RootPath + filePath + "/" + fileName

	downloadBytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusNotFound)
		return
	}

	fileSize := len(string(downloadBytes))

	w.Header().Set("Content-Type", "application/force-download")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(downloadBytes))

	//io.Copy(w, bytes.NewReader(downloadBytes))
}

func BeforeUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	var userArgs struct {
		Token string `json:"token"`
		Path  string `json:"path"`
		Name  string `json:"name"`
	}

	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)

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

	token, err := database.CreateNewUploadToken(gDatabase, userArgs.Token, databasePath, userArgs.Name)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	makeJsonHeader(w, "POST", map[string]string{
		"uploadToken": strings.TrimPrefix(token, "\\"),
	})
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	uploadToken := "\\" + variables["uploadToken"]

	_, filePath, _, err := database.DataByUploadToken(gDatabase, uploadToken)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	err = database.DeleteUploadToken(gDatabase, uploadToken)
	if err != nil {
		fmt.Println(err)
	}

	file, header, _ := r.FormFile("file")
	defer file.Close()

	fullPath := filesystem.RootPath + filePath
	err = filesystem.CreateFile(fullPath, header.Filename, file)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	err = database.CreateFile(gDatabase, filePath, database.FileInfo{
		Name:     header.Filename,
		IsFolder: false,
	})

	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}
}

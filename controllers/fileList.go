package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FileInfo struct {
	Name     string `json:"name"`
	IsFolder bool   `json:"isFolder"`
	Time     string `json:"date"`
}

func GetAllFiles(w http.ResponseWriter, r *http.Request) {
	founedFiles, _ := ioutil.ReadDir("./files/")

	//resp := make([])

	var files struct {
		Data []FileInfo //`json:"files"`
	}

	for _, file := range founedFiles {
		fileObj := FileInfo{file.Name(), file.IsDir(), file.ModTime().String()}
		files.Data = append(files.Data, fileObj)
		//
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "application/json")
	json.NewEncoder(w).Encode(files)
}

package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/*
SendPrivateFile send file to owner user
*/
func SendPrivateFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("file")
	filePath := "./files/" + fileName

	downloadBytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	mime := http.DetectContentType(downloadBytes)

	fileSize := len(string(downloadBytes))

	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(downloadBytes))
}

/*
SendPublicFile send file to owner user
*/
func SendPublicFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("file")
	filePath := "./files/" + fileName

	downloadBytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	mime := http.DetectContentType(downloadBytes)

	fileSize := len(string(downloadBytes))

	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(downloadBytes))
}

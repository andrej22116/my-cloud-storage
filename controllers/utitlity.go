package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"./database"
)

func makeJsonHeader(w http.ResponseWriter, method string, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", method)
	w.Header().Set("Access-Control-Allow-Headers", "application/json")
	err := json.NewEncoder(w).Encode(obj)

	return err
}

func makeFileHeader(w http.ResponseWriter, mime, fileName string, fileSize int) {
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
}

func makeErrorHeader(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func userDataByRequestToken(body io.ReadCloser) (database.UserData, error) {
	token := database.Token{}
	err := json.NewDecoder(body).Decode(&token)
	if err != nil {
		return database.UserData{}, err
	}

	return database.CheckAccess(gDatabase, token.Token)
}

func jsonFromBody(r *http.Request, obj interface{}) error {
	return json.NewDecoder(r.Body).Decode(obj)
}

///////////////////////////////////////////////////////////
type InvalidTokenError struct {
	What string
}

func (e InvalidTokenError) Error() string {
	return fmt.Sprint(e.What)
}

func checkAccess(token string) (database.UserData, error) {
	userData, err := database.CheckAccess(gDatabase, token)
	if err != nil {
		return database.UserData{}, err
	}

	if len(userData.Nickname) == 0 {
		return database.UserData{}, InvalidTokenError{
			What: "Invalid token!",
		}
	}

	return userData, nil
}

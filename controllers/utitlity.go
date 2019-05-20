package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"./database"
)

/*
makeJsonHeader - создаёт заголовок для ответа с включённым JSON объектом
*/
func makeJsonHeader(w http.ResponseWriter, method string, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", method)
	w.Header().Set("Access-Control-Allow-Headers", "application/json")
	err := json.NewEncoder(w).Encode(obj)

	return err
}

/*
makeFileHeader - готовит заголовок для загрузки файла на сервер
*/
func makeFileHeader(w http.ResponseWriter, mime, fileName string, fileSize int) {
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
}

/*
makeErrorHeader - Собирает заголовок с ошибкой
*/
func makeErrorHeader(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

/*
jsonFromBody - десереализует JSON в структуру
*/
func jsonFromBody(r *http.Request, obj interface{}) error {
	return json.NewDecoder(r.Body).Decode(obj)
}

/*
InvalidTokenError - ошибка
*/
type InvalidTokenError struct {
	What string
}

/*
InvalidTokenError Error - реализация интерфейса
*/
func (e InvalidTokenError) Error() string {
	return fmt.Sprint(e.What)
}

/*
checkAccess - проверка токена на валидность
*/
func checkAccess(token string) (database.UserData, error) {
	// Запрос к БД
	userData, err := database.CheckAccess(gDatabase, token)
	if err != nil {
		return database.UserData{}, err
	}

	// Дополнительные проверки
	if len(userData.Nickname) == 0 {
		return database.UserData{}, InvalidTokenError{
			What: "Invalid token!",
		}
	}

	return userData, nil
}

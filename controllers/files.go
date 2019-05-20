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
BeforeLoadFileHandler проверяет параметры пользователя и отправляет ему ключ для загрузки с сервера
*/
func BeforeLoadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Переменная в виде структуры JSON пакета
	var userArgs struct {
		// Название, тип, название в JSON
		Token string `json:"token"`
		Path  string `json:"path"`
		Name  string `json:"name"`
	}

	// Десереализуем JSON в созданную переменную
	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Проверяем токен пользователя
	userData, err := checkAccess(userArgs.Token)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Если хотим получить доступ к корню - удаляем слеш, если он есть
	if userArgs.Path == "/" {
		userArgs.Path = ""
	}

	// Создаём путь, как он будет храниться в БД
	databasePath := "/" + userData.Nickname + userArgs.Path

	// Пытаемся получить загрузочный токен
	token, err := database.CreateNewUploadToken(gDatabase, userArgs.Token, databasePath, userArgs.Name)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Создаём ответ, в который кладём JSON с токеном для загрузки файла
	makeJsonHeader(w, "POST", map[string]string{
		"loadToken": strings.TrimPrefix(token, "\\"),
	})
}

/*
LoadFileHandler проверяет ключ в адресе и отправляет пользователю файл
*/
func LoadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Получем параметры, которые обработал mux - фрэймворк
	variables := mux.Vars(r)
	// Получаем токен из url запроса
	loadToken := "\\" + variables["loadToken"]

	// Получаем информацию о загружаемом файле
	_, filePath, fileName, err := database.DataByUploadToken(gDatabase, loadToken)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Удаляем токен файловой загрузки
	err = database.DeleteUploadToken(gDatabase, loadToken)
	if err != nil {
		fmt.Println(err)
	}

	// Генерируем полный путь к файлу
	fullPath := filesystem.RootPath + filePath + "/" + fileName

	// Читаем его
	downloadBytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusNotFound)
		return
	}

	// Записываем размер
	fileSize := len(string(downloadBytes))

	// Создаём заголовок
	w.Header().Set("Content-Type", "application/force-download")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	// Записываем содержимое файла в ответ
	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(downloadBytes))
}

/*
BeforeUploadFileHandler проверяет параметры пользователя и отправляет ему ключ для загрузки на сервер
*/
func BeforeUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Переменная в виде структуры JSON пакета
	var userArgs struct {
		// Название, тип, название в JSON
		Token string `json:"token"`
		Path  string `json:"path"`
		Name  string `json:"name"`
	}

	// Десереализуем JSON в созданную переменную
	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Проверяем токен пользователя
	userData, err := checkAccess(userArgs.Token)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Если хотим получить доступ к корню - удаляем слеш, если он есть
	if userArgs.Path == "/" {
		userArgs.Path = ""
	}

	// Создаём путь, как он будет храниться в БД
	databasePath := "/" + userData.Nickname + userArgs.Path

	// Создаём новый токен загрузки файла на сервер
	token, err := database.CreateNewUploadToken(gDatabase, userArgs.Token, databasePath, userArgs.Name)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Отпарвляем пользователю токен для загрузки файла
	makeJsonHeader(w, "POST", map[string]string{
		"uploadToken": strings.TrimPrefix(token, "\\"),
	})
}

/*
UploadFileHandler проверяет ключ в адресе и загружает пользовательский файл
*/
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Получем параметры, которые обработал mux - фрэймворк
	variables := mux.Vars(r)
	// Получаем токен из url запроса
	uploadToken := "\\" + variables["uploadToken"]

	// Получаем информацию о загружаемом файле
	_, filePath, _, err := database.DataByUploadToken(gDatabase, uploadToken)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Удаляет токен загрузки
	err = database.DeleteUploadToken(gDatabase, uploadToken)
	if err != nil {
		fmt.Println(err)
	}

	// Получаем данные формы
	file, header, _ := r.FormFile("file")
	// Готовим функцию закрытия (будет вызвана в после выхода из основной функции)
	defer file.Close()

	// Создаём файл
	fullPath := filesystem.RootPath + filePath
	err = filesystem.CreateFile(fullPath, header.Filename, file)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Создаём запись в БД
	err = database.CreateFile(gDatabase, filePath, database.FileInfo{
		Name:     header.Filename,
		IsFolder: false,
	})

	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		filesystem.Remove(fullPath, header.Filename)
		return
	}
}

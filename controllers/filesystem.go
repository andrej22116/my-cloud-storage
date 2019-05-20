package controllers

import (
	"fmt"
	"net/http"

	"./database"
	"./filesystem"
)

/*
CreateFolderHandler - обработчик запроса на создание новой папки.
*/
func CreateFolderHandler(w http.ResponseWriter, r *http.Request) {
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
		// Печатаем в консоль, что что-то пошло не так.
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
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
	// И каким он будет для сервера
	fullPath := filesystem.RootPath + databasePath

	// Создаём папку
	err = filesystem.CreateFolder(fullPath, userArgs.Name)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Просим БД это дело зафиксировать
	err = database.CreateFile(gDatabase, databasePath, database.FileInfo{
		Name:     userArgs.Name,
		IsFolder: true,
	})

	// В случае ошибки отменяем изменения
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		filesystem.Remove(fullPath, userArgs.Name)
		return
	}
}

/*
ModifyFileHandler - обработчик запроса на переименовывание файла
*/
func ModifyFileHandler(w http.ResponseWriter, r *http.Request) {
	// Переменная в виде структуры JSON пакета
	var userArgs struct {
		// Название, тип, название в JSON
		Token   string `json:"token"`
		Path    string `json:"path"`
		OldName string `json:"oldName"`
		NewName string `json:"newName"`
	}

	// Десереализуем JSON в созданную переменную
	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
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
	// И каким он будет для сервера
	fullPath := filesystem.RootPath + databasePath

	// Переименовываем физический файл
	err = filesystem.Rename(fullPath, userArgs.OldName, userArgs.NewName)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Делаем запрос на изменение имени
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

/*
RemoveFileHandler - обработчик запроса на удаление файла
*/
func RemoveFileHandler(w http.ResponseWriter, r *http.Request) {
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
		makeErrorHeader(w, http.StatusBadRequest)
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
	// И каким он будет для сервера
	fullPath := filesystem.RootPath + databasePath

	// Удаляем файл
	err = filesystem.Remove(fullPath, userArgs.Name)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Удаляем данные о файле из БД
	err = database.RemoveFile(gDatabase, databasePath, database.FileInfo{
		Name: userArgs.Name,
	})

	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}
}

/*
FileListHandler - обработчик запроса на получение списка файлов в каталоге
*/
func FileListHandler(w http.ResponseWriter, r *http.Request) {
	// Переменная в виде структуры JSON пакета
	var userArgs struct {
		// Название, тип, название в JSON
		Token string `json:"token"`
		Path  string `json:"path"`
	}

	// Десереализуем JSON в созданную переменную
	err := jsonFromBody(r, &userArgs)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
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

	// Запрашиваем список файлов из БД
	files, err := database.GetFileListFromPath(gDatabase, databasePath)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Создаём ответ с JSON внутри
	makeJsonHeader(w, "POST", files)
}

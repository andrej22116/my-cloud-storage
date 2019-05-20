package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
FileInfo - структура с информацией о файле
*/
type FileInfo struct {
	Name     string `json:"name"`
	IsFolder bool   `json:"isFolder"`
	Time     string `json:"date"`
}

/*
FolderNotExistsError - структура ошибки отсутствия папки
*/
type FolderNotExistsError struct {
	What string
}

/*
FolderNotExistsError Error - реализация интерфейса ошибки
*/
func (e FolderNotExistsError) Error() string {
	return fmt.Sprint(e.What)
}

/*
FileNotExistsError - структура отсутствия файла
*/
type FileNotExistsError struct {
	What string
}

/*
FileNotExistsError Error - реализация интерфейса ошибки
*/
func (e FileNotExistsError) Error() string {
	return fmt.Sprint(e.What)
}

/*
Константы с запросами
*/
const (
	getFolderContent            = "select * from get_folder_content($1);"
	testFolderExists            = "select * from test_folder_exists($1);"
	testFileExists              = "select * from test_file_exists($1, $2);"
	createFileInFolder          = "select * from create_file_in_folder($1, $2, $3);"
	renameFileInFolder          = "select * from rename_file_in_folder($1, $2, $3);"
	deleteFileInFolder          = "select * from delete_file_in_folder($1, $2);"
	createUploadTokenForSession = "select * from create_upload_token_for_session($1, $2, $3);"
	dataForUploadToken          = "select * from data_for_upload_token($1);"
	deleteUploadTokenForSession = "select * from delete_upload_token_for_session($1);"
)

/*
GetFileListFromPath - возвращает список файлов в каталоге
*/
func GetFileListFromPath(database *sql.DB, path string) ([]FileInfo, error) {
	rows, err := database.Query(getFolderContent, path)

	if err != nil {
		return nil, err
	}

	result := []FileInfo{}

	defer rows.Close()
	for rows.Next() {
		var fileInfo FileInfo
		rows.Scan(&fileInfo.Name, &fileInfo.Time, &fileInfo.IsFolder)
		result = append(result, fileInfo)
	}

	return result, nil
}

/*
CreateFile - отправляет запрос на создание файла в БД
*/
func CreateFile(database *sql.DB, path string, fileInfo FileInfo) error {
	// Проверка существования папки
	err := checkFolderExists(database, path)
	if err != nil {
		return err
	}

	// Выполняем запрос
	_, err = database.Exec(createFileInFolder, path, fileInfo.Name, fileInfo.IsFolder)

	// Возвращаем результат
	return err
}

/*
ModifyFile - изменяет имя файла
*/
func ModifyFile(database *sql.DB, path string, oldFileInfo FileInfo, newFileInfo FileInfo) error {
	// Проверяем наличие файла
	err := checkFileExists(database, path, oldFileInfo.Name)
	if err != nil {
		return err
	}

	// Выполняем запрос
	_, err = database.Exec(renameFileInFolder, path, oldFileInfo.Name, newFileInfo.Name)

	// Возвращаем результат
	return err
}

/*
RemoveFile - удаляет файл
*/
func RemoveFile(database *sql.DB, path string, fileInfo FileInfo) error {
	err := checkFileExists(database, path, fileInfo.Name)
	if err != nil {
		return err
	}

	_, err = database.Exec(deleteFileInFolder, path, fileInfo.Name)

	return nil
}

/*
checkFolderExists - проверяет существование папки
*/
func checkFolderExists(database *sql.DB, path string) error {
	// Выполняется запрос
	rows, err := database.Query(testFolderExists, path)
	if err != nil {
		return err
	}

	// подготавливаем функцию закрытия
	defer rows.Close()
	var exists bool
	// Читаем результат
	rows.Next()
	rows.Scan(&exists)

	// Если результата нет - ошибка
	if !exists {
		return FolderNotExistsError{
			"Folder " + path + " not exists!",
		}
	}

	return nil
}

/*
checkFileExists - проверяет существование файла
*/
func checkFileExists(database *sql.DB, path string, fileName string) error {
	// Выполняется запрос
	rows, err := database.Query(testFileExists, path, fileName)
	if err != nil {
		return err
	}

	// подготавливаем функцию закрытия
	defer rows.Close()
	var exists bool
	// Читаем результат
	rows.Next()
	rows.Scan(&exists)

	// Если результата нет - ошибка
	if !exists {
		return FolderNotExistsError{
			"File " + path + "/" + fileName + " not exists!",
		}
	}

	return nil
}

/*
CreateNewUploadToken - запрос на создание загрузочного токена
*/
func CreateNewUploadToken(database *sql.DB, token, path, fileName string) (string, error) {
	// Выполняем запрос
	rows, err := database.Query(createUploadTokenForSession, token, path, fileName)
	if err != nil {
		return "", err
	}

	// Подготавливаем функцию закрытия
	defer rows.Close()

	//Читаем и возвращаем токен
	uploadToken := ""

	rows.Next()
	rows.Scan(&uploadToken)

	return uploadToken, nil
}

/*
DataByUploadToken - получаем информацию о загружаемых данных по токену
*/
func DataByUploadToken(database *sql.DB, token string) (string, string, string, error) {
	// Выполняем запрос
	rows, err := database.Query(dataForUploadToken, token)
	if err != nil {
		return "", "", "", err
	}

	// Подготавливаем функцию закрытия
	defer rows.Close()

	// Читаем информацию и возвращаем её
	userToken := ""
	filePath := ""
	fileName := ""

	rows.Next()
	rows.Scan(&userToken, &filePath, &fileName)

	return userToken, filePath, fileName, nil
}

/*
DeleteUploadToken - запрос на удаление загрузочного токена
*/
func DeleteUploadToken(database *sql.DB, token string) error {
	// Выполняем запрос и возвращаем результат
	_, err := database.Exec(deleteUploadTokenForSession, token)
	return err
}

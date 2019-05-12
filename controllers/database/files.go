package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type FileInfo struct {
	Name     string `json:"name"`
	IsFolder bool   `json:"isFolder"`
	Time     string `json:"date"`
}

type FolderNotExistsError struct {
	What string
}

func (e FolderNotExistsError) Error() string {
	return fmt.Sprint(e.What)
}

type FileNotExistsError struct {
	What string
}

func (e FileNotExistsError) Error() string {
	return fmt.Sprint(e.What)
}

const (
	getFolderContent            = "select * from get_folder_content($1);"
	testFolderExists            = "select * from test_folder_exists($1);"
	testFileExists              = "select * from test_file_exists($1, $2);"
	createFileInFolder          = "select * from create_file_in_folder($1, $2, $3);"
	deleteFileInFolder          = "select * from delete_file_in_folder($1, $2);"
	createUploadTokenForSession = "select * from create_upload_token_for_session($1, $2, $3);"
	dataForUploadToken          = "select * from data_for_upload_token($1);"
	deleteUploadTokenForSession = "select * from delete_upload_token_for_session($1);"
)

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

func CreateFile(database *sql.DB, path string, fileInfo FileInfo) error {
	err := checkFolderExists(database, path)
	if err != nil {
		return err
	}

	_, err = database.Exec(createFileInFolder, path, fileInfo.Name, fileInfo.IsFolder)

	return err
}

func ModifyFile(database *sql.DB, path string, oldFileInfo FileInfo, newFileInfo FileInfo) error {

	/*
		_, err = database.Exec(create_file_in_folder, path, fileInfo.Name, fileInfo.IsFolder)

		return err*/

	return nil
}

func RemoveFile(database *sql.DB, path string, fileInfo FileInfo) error {
	err := checkFileExists(database, path, fileInfo.Name)
	if err != nil {
		return err
	}

	_, err = database.Exec(deleteFileInFolder, path, fileInfo.Name)

	return nil
}

func checkFolderExists(database *sql.DB, path string) error {
	rows, err := database.Query(testFolderExists, path)
	if err != nil {
		return err
	}

	defer rows.Close()
	var exists bool
	rows.Next()
	rows.Scan(&exists)

	if !exists {
		return FolderNotExistsError{
			"Folder " + path + " not exists!",
		}
	}

	return nil
}

func checkFileExists(database *sql.DB, path string, fileName string) error {
	rows, err := database.Query(testFileExists, path, fileName)
	if err != nil {
		return err
	}

	defer rows.Close()
	var exists bool
	rows.Next()
	rows.Scan(&exists)

	if !exists {
		return FolderNotExistsError{
			"File " + path + "/" + fileName + " not exists!",
		}
	}

	return nil
}

func CreateNewUploadToken(database *sql.DB, token, path, fileName string) (string, error) {
	rows, err := database.Query(createUploadTokenForSession, token, path, fileName)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	uploadToken := ""

	rows.Next()
	rows.Scan(&uploadToken)

	return uploadToken, nil
}

func DataByUploadToken(database *sql.DB, token string) (string, string, string, error) {
	rows, err := database.Query(dataForUploadToken, token)
	if err != nil {
		return "", "", "", err
	}

	defer rows.Close()

	userToken := ""
	filePath := ""
	fileName := ""

	rows.Next()
	rows.Scan(&userToken, &filePath, &fileName)

	return userToken, filePath, fileName, nil
}

func DeleteUploadToken(database *sql.DB, token string) error {
	_, err := database.Exec(deleteUploadTokenForSession, token)
	return err
}

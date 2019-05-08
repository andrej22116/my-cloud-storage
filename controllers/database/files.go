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
	getFolderContent   = "select * from get_folder_content($1);"
	testFolderExists   = "select * from test_folder_exists($1);"
	testFileExists     = "select * from test_file_exists($1, $2);"
	createFileInFolder = "select * from create_file_in_folder($1, $2, $3);"
	deleteFileInFolder = "select * from delete_file_in_folder($1, $2);"
)

func GetUserRoot(database *sql.DB, userData UserData) ([]FileInfo, error) {
	return GetUserPath(database, userData, "/"+userData.Nickname)
}

func GetUserPath(database *sql.DB, userData UserData, path string) ([]FileInfo, error) {
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

func CreateFile(database *sql.DB, userData UserData, path string, fileInfo FileInfo) error {
	err := checkFolderExists(database, path)
	if err != nil {
		return err
	}

	_, err = database.Exec(createFileInFolder, path, fileInfo.Name, fileInfo.IsFolder)

	return err
}

func ModifyFile(database *sql.DB, userData UserData, path string, oldFileInfo FileInfo, newFileInfo FileInfo) error {

	/*
		_, err = database.Exec(create_file_in_folder, path, fileInfo.Name, fileInfo.IsFolder)

		return err*/

	return nil
}

func DeleteFile(database *sql.DB, userData UserData, path string, fileInfo FileInfo) error {
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

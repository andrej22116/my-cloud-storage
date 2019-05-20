package filesystem

import (
	"os"
)

/*
Rename - Переименовывает файл или папку
*/
func Rename(path, oldName, newName string) error {
	// Переименовываем и возвращаем результат
	return os.Rename(path+"/"+oldName, path+"/"+newName)
}

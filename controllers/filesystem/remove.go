package filesystem

import "os"

/*
Remove - удаляет файл или папку
*/
func Remove(path, name string) error {
	// Удаляем и возвращаем результат
	return os.RemoveAll(path + "/" + name)
}

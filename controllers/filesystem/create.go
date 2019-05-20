package filesystem

import (
	"io"
	"mime/multipart"
	"os"
)

/*
CreateFile - создаёт файл в файловой системе
*/
func CreateFile(path, name string, file multipart.File) error {
	// Создаём файл и открываем на запись
	out, err := os.OpenFile(path+"/"+name,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	// И копируем в него всё что могём из полученных данных
	defer out.Close()
	_, err = io.Copy(out, file)

	return err
}

/*
CreateFolder - создаёт папку в файловой системе
*/
func CreateFolder(path string, name string) error {
	// Создаёт папку и возвращает результат
	return os.Mkdir(path+"/"+name, os.ModeDir)
}

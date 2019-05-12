package filesystem

import (
	"io"
	"mime/multipart"
	"os"
)

func CreateFile(path, name string, file multipart.File) error {
	out, err := os.OpenFile(path+"/"+name,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	defer out.Close()
	io.Copy(out, file)

	return err
}

func CreateFolder(path string, name string) error {
	return os.Mkdir(path+"/"+name, os.ModeDir)
}

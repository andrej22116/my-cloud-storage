package filesystem

import "os"

func Rename(path, oldName, newName string) error {
	return os.Rename(path+"/"+oldName, path+"/"+newName)
}

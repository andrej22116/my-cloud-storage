package filesystem

import "os"

func Remove(path, name string) error {
	return os.RemoveAll(path + "/" + name)
}

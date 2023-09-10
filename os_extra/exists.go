package osextra

import "os"

func Exists(path ...string) bool {
	for _, currentPath := range path {
		if _, err := os.Stat(currentPath); os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateIfNotExists(path string, perm os.FileMode) error {
	if !Exists(path) {
		return os.MkdirAll(path, perm)
	}
	return nil
}

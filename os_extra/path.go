package osextra

import (
	"path/filepath"
	"strings"
)

func ValidatePath(path string) (string, bool) {
	path = strings.TrimPrefix(strings.ReplaceAll(path, "\\", "/"), "/")
	depth := 0
	directories := strings.Split(path, "/")
	for _, directory := range directories {
		if directory == ".." {
			depth--
		} else {
			depth++
		}
	}
	return filepath.Join(directories...), depth >= 0
}

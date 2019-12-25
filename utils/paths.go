package utils

import (
	"os"
	"path"
	"strings"
)

// GetAbs util
func GetAbs(names []string) []string {

	cwd, err := os.Getwd()
	CheckError(err)

	for i, name := range names {
		if !path.IsAbs(name) {
			names[i] = path.Join(cwd, name)
		}
	}

	return names
}

// IsRel util
func IsRel(path string) bool {
	return strings.HasPrefix(path, ".")
}

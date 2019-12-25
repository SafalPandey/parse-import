package utils

import (
	"errors"
	"os"
)

// CheckError util
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// GetExt util
func GetExt(fpath string, count int) (bool, string, error) {
	extArr := []string{"", ".tsx", ".ts", ".js", ".json"}

	if count >= len(extArr) {
		return false, "", errors.New("Oops no more extensions available: " + fpath)
	}

	_, err := os.Stat(fpath + extArr[count])

	if err != nil {
		return false, extArr[count], nil
	}

	return true, extArr[count], nil
}

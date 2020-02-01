package core

import (
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	"../utils"
)

// LocalDirs global var
var LocalDirs []string

// BaseDirAbsPathMap global var
var BaseDirAbsPathMap = make(map[string]string)

// SplitChar global constant
var SplitChar byte

// Language global constant
var Language string

// Extensions global constant
var Extensions []string

// PathDelimiter global constant
var PathDelimiter string

// ImportPattern global constant
var ImportPattern *regexp.Regexp

// FindLocalDirs global constant
var FindLocalDirs func(string)

// ComputeConstants computes global constants
func ComputeConstants() {
	SplitChar = SplitCharMap[Language]
	Extensions = ExtensionMap[Language]
	PathDelimiter = PathDelimiterMap[Language]
	FindLocalDirs = FindLocalDirsMap[Language]
	ImportPattern = regexp.MustCompile(ImportPatternMap[Language])
}

// SetLocalDirs sets global constant LocalDirs and BaseDirAbsPathMap
func SetLocalDirs(entryPoint string) {
	baseDirAbsPath := path.Dir(entryPoint)

	files, err := ioutil.ReadDir(baseDirAbsPath)
	utils.CheckError(err)

	for _, file := range files {
		localDir := strings.Split(file.Name(), ".")[0]
		LocalDirs = append(LocalDirs, localDir)
		BaseDirAbsPathMap[localDir] = baseDirAbsPath
	}
}

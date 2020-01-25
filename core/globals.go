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

// BaseDirAbsPath global var
var BaseDirAbsPathMap = make(map[string]string)

var SplitChar byte
var Language string
var Extensions []string
var PathDelimiter string
var ImportPattern *regexp.Regexp
var FindLocalDirs func(string)

func ComputeConstants() {
	SplitChar = SplitCharMap[Language]
	Extensions = ExtensionMap[Language]
	PathDelimiter = PathDelimiterMap[Language]
	FindLocalDirs = FindLocalDirsMap[Language]
	ImportPattern = regexp.MustCompile(ImportPatternMap[Language])
}

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

package core

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	"../utils"
)

// LocalDirs global var
var LocalDirs []string

// BaseDirAbsPath global var
var BaseDirAbsPath string

var SplitChar byte
var Language string
var Extensions []string
var PathDelimiter string
var ImportPattern *regexp.Regexp

func ComputeConstants() {
	SplitChar = SplitCharMap[Language]
	Extensions = ExtensionMap[Language]
	PathDelimiter = PathDelimiterMap[Language]
	ImportPattern = regexp.MustCompile(ImportPatternMap[Language])
}

// FindLocalDirs util
func FindLocalDirs(tsconfigPath string) {
	data, err := ioutil.ReadFile(tsconfigPath)
	utils.CheckError(err)

	var obj map[string]map[string]string
	json.Unmarshal(data, &obj)

	BaseDirAbsPath = path.Join(path.Dir(tsconfigPath), obj["compilerOptions"]["baseUrl"])

	files, _ := ioutil.ReadDir(BaseDirAbsPath)
	for _, file := range files {
		LocalDirs = append(LocalDirs, strings.Split(file.Name(), ".")[0])
	}
}

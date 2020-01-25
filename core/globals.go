package core

import (
	"regexp"
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

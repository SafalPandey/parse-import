package core

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"strings"

	"../utils"
)

var ImportPatternMap = map[string]string{
	"ts": `(?sm)(?:^import (?P<name>.+)\s*from\s+(?P<module>\S+))|(?:^(?:const|let|var)\s*(?P<name>\S+)\s*=\s*require\((?P<module>\S+)\))`,
	"py": `(?sm)(?:^import\s+(?P<module>\S+))|(?:^from (?P<module>.+)\s*import\s+(?P<name>\S+))`,
}

var ExtensionMap = map[string][]string{
	"ts": []string{"/index.ts", "/index.tsx", "/index.js", "", ".tsx", ".ts", ".js", ".json"},
	"py": []string{".py"},
}

var SplitCharMap = map[string]byte{
	"ts": ';',
	"py": '\n',
}

var PathDelimiterMap = map[string]string{
	"ts": "/",
	"py": ".",
}

var FindLocalDirsMap = map[string]func(string){
	"ts": func(tsconfigPath string) {
		data, err := ioutil.ReadFile(tsconfigPath)
		utils.CheckError(err)

		var obj map[string]map[string]string
		json.Unmarshal(data, &obj)

		baseDirAbsPath := path.Join(path.Dir(tsconfigPath), obj["compilerOptions"]["baseUrl"])

		files, _ := ioutil.ReadDir(baseDirAbsPath)
		for _, file := range files {
			localDir := strings.Split(file.Name(), ".")[0]
			LocalDirs = append(LocalDirs, localDir)
			BaseDirAbsPathMap[localDir] = baseDirAbsPath
		}
	},
}

package core

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"../utils"
)

var ImportPatternMap = map[string]string{
	"ts": `(?sm)(?:^import (?P<name>.+)\s*from\s+(?P<module>\S+))|(?:^\s*(?:const|let|var)\s*(?P<name>\S+)\s*=\s*require\((?P<module>\S+)\))`,
	"py": `(?sm)(?:^import\s+(?P<module>\S+))|(?:^from (?P<module>\S+)\s*import\s+(?P<name>\S+))`,
}

var ExtensionMap = map[string][]string{
	"ts": []string{"/index.ts", "/index.tsx", "/index.js", "", ".tsx", ".ts", ".js", ".json"},
	"py": []string{".py", "main.py", ""},
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

		entryPoint := path.Join(path.Join(path.Dir(tsconfigPath), obj["compilerOptions"]["baseUrl"]), "index.ts")
		SetLocalDirs(entryPoint)
	},

	"py": SetLocalDirs,
}

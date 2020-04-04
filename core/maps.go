package core

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"../utils"
)

// ImportPatternMap map constant
var ImportPatternMap = map[string][]string{
	"ts": []string{
		`^import (?P<name>.+)\s*from\s+(?P<module>\S+)`,
		`^\s*(?:const|let|var)\s+(?P<name>\S+)\s*=\s*require\((?P<module>\S+)\)`,
	},
	"py": []string{
		`(?sm)(?:^import\s+(?P<module>\S+)`,
		`^from (?P<module>\S+)\s*import\s+(?P<name>\S+)`,
	},
}

// ExtensionMap map constant
var ExtensionMap = map[string][]string{
	"ts": []string{"/index.ts", "/index.tsx", "/index.js", "", ".tsx", ".ts", ".js", ".json"},
	"py": []string{".py", "main.py", ""},
}

// SplitCharMap map constant
var SplitCharMap = map[string]byte{
	"ts": ';',
	"py": '\n',
}

// PathDelimiterMap map constant
var PathDelimiterMap = map[string]string{
	"ts": "/",
	"py": ".",
}

// FindLocalDirsMap map constant
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

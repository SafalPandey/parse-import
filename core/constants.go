package core

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

package core

import (
	"bufio"
	"math"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"

	"../types"
	"../utils"
)

var pat = regexp.MustCompile(`^import (?P<name>.+) from (?P<module>.+)`)

// ParseImport will mutate the passed map with all the dependent imports and their info
func ParseImport(files []string, importMap map[string]interface{}) {
	importMap1 := make(map[string]interface{})
	importMap2 := make(map[string]interface{})

	halfIndex := int(math.Ceil(float64(len(files)) / 2))

	var grp sync.WaitGroup

	grp.Add(1)
	go subParse(files[:halfIndex], importMap1, &grp)

	grp.Add(1)
	go subParse(files[halfIndex:], importMap2, &grp)

	grp.Wait()

	utils.MergeMaps(importMap1, importMap2)
	utils.MergeMaps(importMap, importMap1)
}

func subParse(files []string, importMap map[string]interface{}, grp *sync.WaitGroup) {
	for _, fileName := range files {
		infos := getImports(fileName)

		localPaths, importMap := updateMap(infos, importMap)

		localPaths = utils.Filter(localPaths, func(x string) bool {
			_, ok := importMap[x]

			return ok
		})
		ParseImport(localPaths, importMap)

	}

	grp.Done()
}

func getImports(fileName string) []types.ImportInfo {
	var imports []types.ImportInfo

	file, err := os.Open(fileName)
	utils.CheckError(err)

	defer file.Close()

	lineNum := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		submatches := pat.FindStringSubmatch(line)

		if len(submatches) != 0 {
			name := submatches[1]
			module := submatches[2]
			filePath := strings.Trim(module, "';")
			isDir := false

			isRel := utils.IsRel(filePath)
			pathIsFromBaseDir := utils.StartsWithAnyOf(LocalDirs, filePath)

			if isRel || pathIsFromBaseDir {
				if pathIsFromBaseDir {
					filePath = path.Join(BaseDirAbsPath, filePath)
				} else {
					filePath = path.Join(path.Dir(fileName), filePath)
				}

				i := 0
				ext := ""
				done := false

				for !done {
					done, ext, err = utils.GetExt(filePath, i)
					utils.CheckError(err)
					i++
				}

				fi, err := os.Stat(filePath + ext)

				if err == nil && fi.Mode().IsDir() {
					isDir = true
					filePath += "/"
				} else {
					filePath += ext
				}
			}

			imports = append(imports, types.ImportInfo{
				Line:       lineNum,
				Name:       name,
				Module:     module,
				Path:       filePath,
				IsDir:      isDir,
				ImportedIn: fileName,
			})
		}

		lineNum++
	}

	return imports
}

func updateMap(paths []types.ImportInfo, importMap map[string]interface{}) ([]string, map[string]interface{}) {
	var localPaths []string
	for _, p := range paths {
		if !path.IsAbs(p.Path) {
			importMap[p.Path] = types.MapNode{IsLocal: false, Path: p.Path, Info: p}
		} else {
			importMap[p.Path] = utils.BuildMap(importMap, strings.Split(p.Path, "/")[1:], p)

			if !p.IsDir {
				localPaths = append(localPaths, p.Path)
			}
		}
	}

	return localPaths, importMap
}

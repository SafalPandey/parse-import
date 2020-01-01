package core

import (
	"bufio"
	"fmt"
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
	var parseGrp sync.WaitGroup

	parsedMap := make(map[string]bool)
	infoChan := make(chan []types.ImportInfo)

	for _, file := range files {
		parsedMap[file] = true
	}

	parseGrp.Add(1)
	go parse(files, infoChan, &parseGrp)

	// Update import map when new info is avalable in channel
	go func() {
		for infos := range infoChan {

			localPaths, _ := updateMap(infos, importMap)

			localPaths = utils.Filter(localPaths, func(x string) bool {
				exists := parsedMap[x]

				return !exists
			})

			for _, x := range localPaths {
				parsedMap[x] = true
			}

			go parse(localPaths, infoChan, &parseGrp)
		}
	}()

	parseGrp.Wait()
	close(infoChan)
	fmt.Println(len(parsedMap))
}

func parse(files []string, infoChan chan<- []types.ImportInfo, parseGrp *sync.WaitGroup) {
	halfIndex := int(math.Ceil(float64(len(files)) / 2))

	parseGrp.Add(1)
	go subParse(files[:halfIndex], infoChan, parseGrp)

	subParse(files[halfIndex:], infoChan, parseGrp)
}

func subParse(files []string, infoChan chan<- []types.ImportInfo, parseGrp *sync.WaitGroup) {
	for _, fileName := range files {
		infos := getImports(fileName)

		parseGrp.Add(1)
		infoChan <- infos
	}
	parseGrp.Done()
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
			filePath := strings.Trim(module, "'\";")
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

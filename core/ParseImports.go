package core

import (
	"bufio"
	"errors"
	"math"
	"os"
	"path"
	"strings"
	"sync"

	"../types"
	"../utils"
)

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

			localPaths = utils.FilterString(localPaths, func(x string) bool {
				exists := parsedMap[x]

				parsedMap[x] = true

				return !exists
			})

			go parse(localPaths, infoChan, &parseGrp)
		}
	}()

	parseGrp.Wait()
	close(infoChan)
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
	scanner.Split(utils.GetSplitterFunc(SplitChar))

	for scanner.Scan() {
		line := scanner.Text()
		submatches := utils.FindNamedMatches(ImportPattern, line)

		if len(submatches) != 0 {
			name := submatches["name"]
			module := submatches["module"]

			modulePath := strings.Join(
				strings.Split(
					strings.Trim(module, "'\";"),
					PathDelimiter,
				),
				"/",
			)

			isDir := false

			isRel := utils.IsRel(modulePath)
			pathIsFromBaseDir, baseDir := utils.StartsWithAnyOf(LocalDirs, modulePath, "/")

			if isRel || pathIsFromBaseDir {
				if pathIsFromBaseDir {
					modulePath = path.Join(BaseDirAbsPathMap[baseDir], modulePath)
				} else {
					modulePath = path.Join(path.Dir(fileName), modulePath)
				}

				modulePath, isDir = getFilePath(modulePath)
			}

			imports = append(imports, types.ImportInfo{
				Path:  modulePath,
				IsDir: isDir,
				Importers: []types.ImportedIn{
					{
						Name:   name,
						Module: module,
						Line:   lineNum,
						Path:   fileName,
					},
				},
			})
		}

		lineNum++
	}

	return imports
}

func updateMap(paths []types.ImportInfo, importMap map[string]interface{}) ([]string, map[string]interface{}) {
	var localPaths []string

	for _, p := range paths {
		isLocal := false
		var importedIn []types.ImportedIn

		if importMap[p.Path] != nil {
			importedIn = importMap[p.Path].(types.MapNode).Info.Importers
		}

		if path.IsAbs(p.Path) {
			isLocal = true
			localPaths = append(localPaths, p.Path)
			// importMap[p.Path] = utils.BuildNestedMap(importMap[p.Path], strings.Split(p.Path, "/")[1:], p)
		}

		importMap[p.Path] = types.MapNode{
			IsLocal: isLocal,
			Path:    p.Path,
			Info: types.ImportInfo{
				Path:      p.Path,
				IsDir:     p.IsDir,
				Importers: append(importedIn, p.Importers...),
			},
		}
	}

	return localPaths, importMap
}

// Iterates over Extensions (languageSpecific) and
// returns filePath and isDir boolean if a file exists
// throws error if no file is found.
func getFilePath(fpath string) (string, bool) {
	for i := range Extensions {
		path := fpath + Extensions[i]

		fi, err := os.Stat(path)

		if err == nil {
			return path, fi.IsDir()
		}
	}

	panic(errors.New("Oops no more extensions available: " + fpath))
}

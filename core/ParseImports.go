package core

import (
	"errors"
	"io/ioutil"
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
	invalidEntrypointMap := make(map[string]string)

	for _, file := range files {
		fi, err := os.Stat(file)
		utils.CheckError(err)

		if fi.IsDir() {
			invalidEntrypointMap[file] = "Entrypoint cannot be a directory. Please pass a file instead."
		}

		parsedMap[file] = true
	}

	if len(invalidEntrypointMap) > 0 {
		message := "Entrypoint is invalid:\n"

		for file, errMsg := range invalidEntrypointMap {
			message += "  \"" + file + "\" <- " + errMsg + "\n"
		}

		panic(message)
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

	contents, err := ioutil.ReadFile(fileName)
	utils.CheckError(err)

	allMatches := utils.FindAllNamedMatches(ImportPattern, string(contents))

	for _, submatches := range allMatches {
		name := submatches["name"]
		module := submatches["module"]

		modulePath := strings.Join(strings.Split(strings.Trim(module, "'\";"), PathDelimiter), "/")

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
					Path:   fileName,
				},
			},
		})
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

			if !p.IsDir {
				localPaths = append(localPaths, p.Path)
			}
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

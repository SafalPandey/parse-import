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

// ParseImport will create a map of imports found in provided files and their info
func ParseImport(files []string) map[string]interface{} {
	var parseGrp sync.WaitGroup

	parsedMap := make(map[string]bool)
	importMap := make(map[string]interface{})
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

	return importMap
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
	importPaths := []string{}

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

		importPaths = append(importPaths, modulePath)

		imports = append(imports, types.ImportInfo{
			Path:    modulePath,
			IsDir:   isDir,
			Imports: []string{},
			Importers: []types.ImportedIn{
				{
					Name:   name,
					Module: module,
					Path:   fileName,
				},
			},
		})
	}

	return append(imports, types.ImportInfo{
		Path:    fileName,
		IsDir:   false,
		Imports: importPaths,
	})
}

func updateMap(infos []types.ImportInfo, importMap map[string]interface{}) ([]string, map[string]interface{}) {
	var localPaths []string

	for _, i := range infos {
		isLocal := false
		var imports []string
		var importedIn []types.ImportedIn

		if importMap[i.Path] != nil {
			imports = importMap[i.Path].(types.MapNode).Info.Imports
			importedIn = importMap[i.Path].(types.MapNode).Info.Importers
		}

		if path.IsAbs(i.Path) {
			isLocal = true

			if !i.IsDir {
				localPaths = append(localPaths, i.Path)
			}
		}

		importMap[i.Path] = types.MapNode{
			IsLocal:      isLocal,
			Path:         i.Path,
			IsEntrypoint: false,
			Info: types.ImportInfo{
				Path:      i.Path,
				IsDir:     i.IsDir,
				Imports:   append(imports, i.Imports...),
				Importers: append(importedIn, i.Importers...),
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

// ValidateEntrypoints checks entrypoints to ensure they are valid
func ValidateEntrypoints(files []string) {
	invalidEntrypointMap := make(map[string]string)

	for _, file := range files {
		fi, err := os.Stat(file)
		utils.CheckError(err)

		if fi.IsDir() {
			invalidEntrypointMap[file] = "Entrypoint cannot be a directory. Please pass a file instead."
		}
	}

	if len(invalidEntrypointMap) > 0 {
		message := "Entrypoint is invalid:\n"

		for file, errMsg := range invalidEntrypointMap {
			message += "  \"" + file + "\" <- " + errMsg + "\n"
		}

		panic(message)
	}
}

// SetEntrypoints creates a map of supplied entrypoints assuming they are local
func SetEntrypoints(entrypoints []string, importMap map[string]interface{}) map[string]interface{} {
	for _, file := range entrypoints {
		val := importMap[file].(types.MapNode)
		val.IsEntrypoint = true

		importMap[file] = val
	}

	return importMap
}

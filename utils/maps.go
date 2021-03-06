package utils

import (
	"regexp"

	"../types"
)

// BuildNestedMap recursively
func BuildNestedMap(obj interface{}, keys []string, p types.ImportInfo) map[string]interface{} {
	var exists bool
	newObj := make(map[string]interface{})

	if obj != nil {
		obj, exists = obj.(map[string]interface{})[keys[0]]
	}

	if len(keys) == 1 {
		var importedIn []types.ImportedIn

		if exists {
			importedIn = obj.(types.MapNode).Info.Importers
		}

		newObj[keys[0]] = types.MapNode{
			IsLocal: true,
			Path:    p.Path,
			Info: types.ImportInfo{
				Path:      p.Path,
				IsDir:     p.IsDir,
				Importers: append(p.Importers, importedIn...),
			},
		}
	} else {
		newObj[keys[0]] = BuildNestedMap(obj, keys[1:], p)
	}

	return newObj
}

// MergeMaps helps you merge two maps.
// Warning: This utils mutates the first map.
// TODO: Make this a pure function.
func MergeMaps(map1 map[string]interface{}, map2 map[string]interface{}) map[string]interface{} {
	for k, v := range map2 {
		map1[k] = v
	}

	return map1
}

// NormalizeMap wil try to normalize a map down using it's key
func NormalizeMap(obj map[string]interface{}) map[string]interface{} {
	// TODO: Implement logic to normalize given object
	return obj
}

// FindNamedMatches creates a map of submatches and returns the map.
func FindNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)
	subexpNames := regex.SubexpNames()

	results := map[string]string{}

	for i, name := range match {
		val, exists := results[subexpNames[i]]

		if !exists || val == "" {
			results[subexpNames[i]] = name
		}
	}

	return results
}

// FindAllNamedMatches creates a slice of maps of submatches and returns the slice.
func FindAllNamedMatches(regex *regexp.Regexp, str string) []map[string]string {
	matches := regex.FindAllStringSubmatch(str, -1)
	subexpNames := regex.SubexpNames()

	var result []map[string]string

	for _, match := range matches {
		matchMap := map[string]string{}

		for i, name := range match {
			val, exists := matchMap[subexpNames[i]]

			if !exists || val == "" {
				matchMap[subexpNames[i]] = name
			}
		}

		result = append(result, matchMap)
	}

	return result
}

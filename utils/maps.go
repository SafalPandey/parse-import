package utils

import (
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
	return obj
}

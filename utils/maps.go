package utils

import (
	"../types"
)

// BuildMap recursively
func BuildMap(obj map[string]interface{}, keys []string, p types.ImportInfo) map[string]interface{} {
	newObj := make(map[string]interface{})
	// newObj = obj

	if len(keys) == 1 {
		newObj[keys[0]] = types.MapNode{
			IsLocal: true,
			Path:    p.Path,
			Info:    p,
		}
	} else {
		newObj[keys[0]] = BuildMap(newObj, keys[1:], p)
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

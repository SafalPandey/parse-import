package utils

import "strings"

// Filter will filter a slice of string using given condition
func Filter(a []string, condition func(string) bool) []string {
	n := 0
	for _, x := range a {
		if condition(x) {
			a[n] = x
			n++
		}
	}
	a = a[:n]

	return a
}

// StartsWithAnyOf will return true and the matched element if a string starts with any element in the given array
func StartsWithAnyOf(a []string, s string, suffix string) (bool, string) {
	for _, element := range a {
		if strings.HasPrefix(s, element+suffix) || s == element {
			return true, element
		}
	}

	return false, ""
}

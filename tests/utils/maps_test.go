package tests

import (
	"reflect"
	"regexp"
	"testing"

	"../../utils"
)

/**
  Should match given string and return a map of submatches.
**/
func TestFindAllNamedMatches(t *testing.T) {
	str := `
    This is a string.
    This is another line in the same string.
    This is last line in string.
  `
	re := regexp.MustCompile(`(?P<this>This) is (?P<a>a\S*) [^\n]*`)

	expected := []map[string]string{
		{
			"":     "This is a string.",
			"a":    "a",
			"this": "This",
		},
		{
			"":     "This is another line in the same string.",
			"a":    "another",
			"this": "This",
		},
	}

	actual := utils.FindAllNamedMatches(re, str)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Filtered result did not match expected output\ngot:\n  \"%s\"\nexpected:\n  \"%s\"", actual, expected)
	}
}

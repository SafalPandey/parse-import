package tests

import (
	"reflect"
	"testing"

	"../../utils"
)

func hasMoreThanOneChar(element string) bool {
	return len(element) > 1
}

/**
  Should filter out strings that do not satisfy the given condition.
**/
func TestFilterStringShouldFilter(t *testing.T) {
	inputSlice := []string{"This", "is", "a", "string", "array."}
	expected := []string{"This", "is", "string", "array."}

	actual := utils.FilterString(inputSlice, hasMoreThanOneChar)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Filtered result did not match expected output, got: \"%s\", expected: \"%s\"", actual, expected)
	}
}

/**
  Should not filter anything if all strings satisfy the given condition.
**/
func TestFilterStringShouldNotFilter(t *testing.T) {
	inputSlice := []string{"This", "is", "an", "array", "of", "strings."}
	expected := []string{"This", "is", "an", "array", "of", "strings."}

	actual := utils.FilterString(inputSlice, hasMoreThanOneChar)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Filtered result did not match expected output, got: \"%s\", expected: \"%s\"", actual, expected)
	}
}

/**
  Should return true and matched element when successfully matched without suffix.
**/
func TestStartsWithAnyOfExists(t *testing.T) {
	inputSlice := []string{"this", "string", "starts", "with"}
	inputString := "this string starts with this"

	expectedBool := true
	expectedMatch := "this"

	actualBool, actualMatch := utils.StartsWithAnyOf(inputSlice, inputString, "")

	if actualBool != expectedBool {
		t.Errorf("Actual boolean output did not match expected boolean output, got: %t, expected: %t", actualBool, expectedBool)
	}

	if actualMatch != expectedMatch {
		t.Errorf("Actual matched output did not match expected matched output, got: \"%s\", expected: \"%s\"", actualMatch, expectedMatch)
	}
}

/**
  Should return false and "" when not matched without suffix.
**/
func TestStartsWithAnyOfExistsNoMatch(t *testing.T) {
	inputSlice := []string{"this", "string", "starts", "with"}
	inputString := "No this string doesn't start with this"

	expectedBool := false
	expectedMatch := ""

	actualBool, actualMatch := utils.StartsWithAnyOf(inputSlice, inputString, "")

	if actualBool != expectedBool {
		t.Errorf("Actual boolean output did not match expected boolean output, got: %t, expected: %t", actualBool, expectedBool)
	}

	if actualMatch != expectedMatch {
		t.Errorf("Actual matched output did not match expected matched output, got: \"%s\", expected: \"%s\"", actualMatch, expectedMatch)
	}
}

/**
  Should return true and matched element when successfully matched with suffix.
**/
func TestStartsWithAnyOfExistsWithSuffix(t *testing.T) {
	inputSlice := []string{"this", "string", "starts", "with"}
	inputString := "this/string/starts/with/this"

	expectedBool := true
	expectedMatch := "this"

	actualBool, actualMatch := utils.StartsWithAnyOf(inputSlice, inputString, "/")

	if actualBool != expectedBool {
		t.Errorf("Actual boolean output did not match expected boolean output, got: %t, expected: %t", actualBool, expectedBool)
	}

	if actualMatch != expectedMatch {
		t.Errorf("Actual matched output did not match expected matched output, got: \"%s\", expected: \"%s\"", actualMatch, expectedMatch)
	}
}

/**
  Should return false and "" when not matched with suffix.
**/
func TestStartsWithAnyOfExistsNoMatchWithSuffix(t *testing.T) {
	inputSlice := []string{"this", "string", "starts", "with"}
	inputString := "No this string doesn't start with this"

	expectedBool := false
	expectedMatch := ""

	actualBool, actualMatch := utils.StartsWithAnyOf(inputSlice, inputString, "/")

	if actualBool != expectedBool {
		t.Errorf("Actual boolean output did not match expected boolean output, got: %t, expected: %t", actualBool, expectedBool)
	}

	if actualMatch != expectedMatch {
		t.Errorf("Actual matched output did not match expected matched output, got: \"%s\", expected: \"%s\"", actualMatch, expectedMatch)
	}
}

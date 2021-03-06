package utils

import "bufio"

// GetSplitterFunc returns a splitter function based for the provided character.
// The returned splitter function can then be passed to scanner.Split function.
//
// The splitter function returned is based on following example from official Go documentation:
// https://golang.org/pkg/bufio/#example_Scanner_emptyFinalToken
func GetSplitterFunc(char byte) func([]byte, bool) (int, []byte, error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		for i := 0; i < len(data); i++ {
			if data[i] == char {
				return i + 1, data[:i], nil
			}
		}

		if !atEOF {
			return 0, nil, nil
		}

		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
}

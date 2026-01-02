package fsx

import (
	"os"
	"strconv"
	"strings"
)

// LoadString reads a file and returns its content as a string, with leading and trailing whitespace removed.
// If an error occurs while reading the file, it is returned instead of the file content.
func LoadString(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

// LoadUint64 reads a file and parses its content as an unsigned 64-bit integer.
// If an error occurs while reading the file, it is returned instead of the parsed value.
// The file content is expected to be a string representation of an unsigned 64-bit integer
// in base 10.
func LoadUint64(path string) (uint64, error) {
	value, err := LoadString(path)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(value, 10, 64)
}

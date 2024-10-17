package simpleio

import (
	"os"
)

// FileHandler opens a file string path and handles any errors that may happen.
func FileHandler(filename string) *os.File {
	if file, err := os.Open(filename); CatchError(err) {
		return file
	}
	return nil
}

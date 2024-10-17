package stdio

import (
	"bufio"
	"github.com/klauspost/pgzip"
	"os"
)

// FileHandle opens a file string path and handles any errors that may happen.
func FileHandle(filename string) *os.File {
	if file, err := os.Open(filename); CatchError(err) {
		return file
	}
	return nil
}

// NewReader creates a new reader
func NewReader(filename string) *Scanalyzer {
	file := FileHandle(filename)
	reader := bufio.NewReader(file)
	if IsGzip(reader) {
		gzipReader, err := pgzip.NewReader(reader)
		CatchError(err)
		return NewScannerio(gzipReader, file)
	}
	return NewScannerio(reader, file)
}

package parseio

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/klauspost/pgzip"
)

// SimpleOpen opens a file and returns the file handle.
func SimpleOpen(filename string) *os.File {
	if file, err := os.Open(filename); ExitOnError(err) {
		return file
	}
	return nil
}

// FileHandler opens a file string path and handles any errors including gzipped files.
func FileHandler(filename string) (*bufio.Reader, *os.File) {
	file := SimpleOpen(filename)
	reader := bufio.NewReader(file)

	if IsGzip(reader) {
		reader = bufio.NewReader(NewGunzip(reader))
	}
	return reader, file
}

// NewGunzip is a helper function to define a pgzip.Reader{} and handles and errors.
func NewGunzip(reader io.Reader) *pgzip.Reader {
	if gunzip, err := pgzip.NewReader(reader); ExitOnError(err) {
		return gunzip
	}
	return nil
}

// IsGzip checks if the file is gzip-compressed by peeking at its magic number
func IsGzip(reader *bufio.Reader) bool {
	if buffer, err := reader.Peek(2); ExitOnError(err) {
		return bytes.Equal(buffer, []byte{0x1f, 0x8b})
	}
	return false
}

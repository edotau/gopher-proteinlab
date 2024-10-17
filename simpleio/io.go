package simpleio

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/klauspost/pgzip"
)

// FileHandler opens a file string path and handles any errors that may happen.
func FileHandler(filename string) *os.File {
	if file, err := os.Open(filename); CatchError(err) {
		return file
	}
	return nil
}

func NewGunzip(reader io.Reader) *pgzip.Reader {
	if gunzip, err := pgzip.NewReader(reader); CatchError(err) {
		return gunzip
	}
	return nil
}

// IsGzip checks if the file is gzip-compressed by peeking at its magic number
func IsGzip(reader *bufio.Reader) bool {
	if buffer, err := reader.Peek(2); CatchError(err) {
		return bytes.Equal(buffer, []byte{0x1f, 0x8b})
	}
	return false
}

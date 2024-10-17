package simpleio

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/klauspost/pgzip"
)

// FileHandler opens a file string path and handles any errors that may happen.
func FileHandler(filename string) (*bufio.Reader, *os.File) {
	if file, err := os.Open(filename); CatchError(err) {
		reader := bufio.NewReader(file)
		if IsGzip(reader) {
			reader = bufio.NewReader(NewGunzip(reader))
		}
		return reader, file
	}
	return nil, nil
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

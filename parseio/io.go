package parseio

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/klauspost/pgzip"
)

// CodeReader struct wraps around bufio.Reader while handling gzzip files as well.
type CodeReader struct {
	*bufio.Reader
	close func() error
}

// SimpleOpen opens a file and returns the file handle.
func SimpleOpen(filename string) *os.File {
	if file, err := os.Open(filename); ExitOnError(err) {
		return file
	}
	return nil
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

// FileHandler opens a file string path and handles any errors including gzipped files.
func FileHandler(filename string) (*bufio.Reader, *os.File) {
	file := SimpleOpen(filename)
	reader := bufio.NewReader(file)

	if IsGzip(reader) {
		reader = bufio.NewReader(NewGunzip(reader))
	}
	return reader, file
}

// NewCodeReader creates a new CodeReader wrapping bufio.Reader and adds a close method.
func NewCodeReader(filename string) *CodeReader {
	reader, file := FileHandler(filename)
	return &CodeReader{
		Reader: bufio.NewReader(reader),
		close:  file.Close,
	}
}


// Close is the method to close the underlying resource, such as a file.
func (r *CodeReader) Close() error {
	if r.close != nil {
		return r.close()
	}
	return nil
}

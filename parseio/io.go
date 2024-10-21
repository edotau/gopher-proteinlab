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

// Scanalyzer structwraps around bufio.Scanner and adds a close method.
type Scanalyzer struct {
	*bufio.Scanner
	close func() error
}

// VimOpen opens a file and it handles errors gracefully.
func VimOpen(filename string) *os.File {
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
	file := VimOpen(filename)
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

// NewScanner creates a new Scanalyzer scanner.
func NewScanner(filename string) *Scanalyzer {
	reader, file := FileHandler(filename)

	return &Scanalyzer{
		Scanner: bufio.NewScanner(reader),
		close:   file.Close,
	}
}

// Read implements io.Reader, reading data into b from the file or gzip stream.
func (reader *CodeReader) Read(b []byte) (n int, err error) {
	return reader.Reader.Read(b)
}

// Close is the method to close the underlying resource, such as a file.
func (r *CodeReader) Close() error {
	if r.close != nil {
		return r.close()
	}
	return nil
}

// Close is the method to close the underlying resource, such as a file.
func (s *Scanalyzer) Close() error {
	if s.close != nil {
		return s.close()
	}
	return nil
}

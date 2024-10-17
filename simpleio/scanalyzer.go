package simpleio

import (
	"bufio"
	"io"
	"os"
)

// Scanalyzer structwraps around bufio.Scanner and adds a close method.
type Scanalyzer struct {
	*bufio.Scanner
	close func() error
}

// Close is the method to close the underlying resource, such as a file.
func (s *Scanalyzer) Close() error {
	if s.close != nil {
		return s.close()
	}
	return nil
}

// NewScanner creates a new *Scanalyzer struct.
func NewScanner(filename string) *Scanalyzer {
	file := FileHandler(filename)
	reader := bufio.NewReader(file)

	if IsGzip(reader) {
		return NewScannerio(NewGunzip(reader), file)
	}
	return NewScannerio(reader, file)
}

// NewScannerio creates a new Scanalyzer from any io.Reader with an optional close function.
func NewScannerio(r io.Reader, file *os.File) *Scanalyzer {
	return &Scanalyzer{
		Scanner: bufio.NewScanner(r),
		close:   file.Close,
	}
}

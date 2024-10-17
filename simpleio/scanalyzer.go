package simpleio

import (
	"bufio"
)

// Scanalyzer structwraps around bufio.Scanner and adds a close method.
type Scanalyzer struct {
	*bufio.Scanner
	close func() error
}

// NewScanner creates a new Scanalyzer scanner.
func NewScanner(filename string) *Scanalyzer {
	reader, file := FileHandler(filename)

	return &Scanalyzer{
		Scanner: bufio.NewScanner(reader),
		close:   file.Close,
	}
}

// Close is the method to close the underlying resource, such as a file.
func (s *Scanalyzer) Close() error {
	if s.close != nil {
		return s.close()
	}
	return nil
}

package simpleio

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
)

func openFile(filename string) (*bufio.Scanner, *os.File, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}

	// Create a buffered reader for peeking and reading
	reader := bufio.NewReader(file)

	// Check if the file is gzip-compressed using Peek
	isGzipped, err := isGzip(filename)
	if err != nil {
		file.Close()
		return nil, nil, fmt.Errorf("error checking if file is gzip: %v", err)
	}

	// If gzip, decompress it
	if isGzipped {
		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			file.Close()
			return nil, nil, fmt.Errorf("error reading gzip file: %v", err)
		}
		return bufio.NewScanner(gzipReader), file, nil
	}

	// Non-gzip file, return a regular scanner
	return bufio.NewScanner(reader), file, nil
}

// isGzip checks if the file is gzip-compressed by peeking at its magic number
func isGzip(filename string) (bool, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return false, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Peek at the first two bytes (magic number) without consuming them
	buffer, err := reader.Peek(2)
	if err != nil {
		return false, fmt.Errorf("error peeking file: %v", err)
	}

	// Check if the file starts with the gzip magic number
	return bytes.Equal(buffer, []byte{0x1f, 0x8b}), nil
}

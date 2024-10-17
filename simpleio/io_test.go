package simpleio

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/klauspost/pgzip"
)

func TestSimpleOpen(t *testing.T) {
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	defer log.SetOutput(os.Stderr) // Restore default log output

	defer func() {
		if r := recover(); r != nil {
			expectedMessage := "open nonexistentfile: no such file or directory\n"
			if !bytes.Contains(logOutput.Bytes(), []byte(expectedMessage)) {
				t.Errorf("Expected log message %q, got %q", expectedMessage, logOutput.String())
			}
		} else {
			t.Errorf("Expected panic but did not get one")
		}
	}()

	// Test SimpleOpen with an invalid file
	if file := SimpleOpen("nonexistentfile"); file != nil {
		t.Errorf("Expected nil, got file")
	}

	// Create a temporary file for testing
	if tmpfile, err := os.Create("testdata/lines.txt"); err == nil {
		defer os.Remove(tmpfile.Name())

		if file := SimpleOpen(tmpfile.Name()); file == nil {
			t.Errorf("Expected file, got nil")
		} else {
			CatchError(file.Close())
		}
	}
}

func TestIsGzip(t *testing.T) {
	invalid := bufio.NewReader(bytes.NewReader([]byte("invalid data")))
	if IsGzip(invalid) {
		t.Errorf("Expected: IsGzip(invalid) == false, but true\n")
	}

	if file, err := os.Open("testdata/uniprot-test.dat.gz"); !CatchError(err) {
		defer CatchError(file.Close())
		if _, err := pgzip.NewReader(bufio.NewReader(file)); !CatchError(err) {
			t.Fatalf("Failed to write to gzip writer: %v", err)
		}
	}
}

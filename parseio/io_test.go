package parseio

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

	// Test VimOpen with an invalid file
	if file := VimOpen("nonexistentfile"); file != nil {
		t.Errorf("Expected nil, got file")
	}

	// Create a temporary file for testing
	if tmpfile, err := os.Create("testdata/lines.txt"); err == nil {
		defer os.Remove(tmpfile.Name())

		if file := VimOpen(tmpfile.Name()); file == nil {
			t.Errorf("Expected file, got nil")
		} else {
			ExitOnError(file.Close())
		}
	}
}

func TestIsGzip(t *testing.T) {
	invalid := bufio.NewReader(bytes.NewReader([]byte("invalid data")))
	if IsGzip(invalid) {
		t.Errorf("Expected: IsGzip(invalid) == false, but true\n")
	}

	if file, err := os.Open("testdata/uniprot-test.dat.gz"); !ExitOnError(err) {
		defer ExitOnError(file.Close())
		if _, err := pgzip.NewReader(bufio.NewReader(file)); !ExitOnError(err) {
			t.Fatalf("Failed to write to gzip writer: %v", err)
		}
	}
}

func TestScanalyzer(t *testing.T) {
	scanner := NewScanner("testdata/uniprot-test.dat.gz")
	if err := scanner.Close(); err != nil {
		t.Errorf("Error: Expected scanner.Close() with no error. %v", err)
	}
}
func TestNewScannerio(t *testing.T) {
	var lines []string
	if tmpfile, err := os.Create("testdata/lines.txt"); ExitOnError(err) {
		defer os.Remove(tmpfile.Name())

		if _, err = tmpfile.WriteString("line1\nline2\nline3\n"); ExitOnError(err) {
			t.Logf("Error: tmpfile.WriteString() = %v\n", err)
		}

		scanner := NewScanner(tmpfile.Name())
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		expected := []string{"line1", "line2", "line3"}

		if len(lines) != len(expected) {
			t.Errorf("Error: number of lines do not match %v, %v", len(lines), len(expected))
		}
		for i := 0; i < len(lines); i++ {
			if lines[i] != expected[i] {
				t.Errorf("Error: NewScannerio() is not parsing the lines correctly.'n")
			}
		}
	}

}

package simpleio

import (
	"bufio"
	"os"
	"testing"
)
func TestScanalyzer(t *testing.T) {
	file := FileHandler("testdata/uniprot-test.dat.gz")
	
	scanner := NewScannerio(bufio.NewReader(file), file)
	if err := scanner.Close(); err != nil {
		t.Errorf("Error: Expected scanner.Close() with no error. %v", err)
	}
}

func TestNewScannerio(t *testing.T) {
	tmpfile, err := os.Create("testdata/lines.txt")
	CatchError(err)
	defer os.Remove(tmpfile.Name())

	if _, err = tmpfile.WriteString("line1\nline2\nline3\n"); CatchError(err) {
		CatchError(tmpfile.Close())
	}

	file := FileHandler(tmpfile.Name())
	scanner := NewScannerio(bufio.NewReader(file), file)

	var lines []string
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
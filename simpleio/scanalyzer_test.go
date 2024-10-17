package simpleio

import (
	"os"
	"testing"
)

func TestScanalyzer(t *testing.T) {
	scanner := NewScanner("testdata/uniprot-test.dat.gz")
	if err := scanner.Close(); err != nil {
		t.Errorf("Error: Expected scanner.Close() with no error. %v", err)
	}
}

func TestNewScannerio(t *testing.T) {
	var lines []string
	if tmpfile, err := os.Create("testdata/lines.txt"); CatchError(err) {
		defer os.Remove(tmpfile.Name())

		if _, err = tmpfile.WriteString("line1\nline2\nline3\n"); CatchError(err) {
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

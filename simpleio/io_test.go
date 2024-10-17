package simpleio

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/klauspost/pgzip"
)

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

package simpleio

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/klauspost/pgzip"
)

func TestIsGzip(t *testing.T) {
	invalid := bufio.NewReader(bytes.NewReader([]byte("invalid data")))
	if IsGzip(invalid) {
		t.Errorf("Expected: IsGzip(invalid) == false, but true\n")
	}
	if gunzip, err := pgzip.NewReader(FileHandler("testdata/uniprot-test.dat.gz")); !CatchError(err) {
		t.Fatalf("Failed to write to gzip writer: %v", err)
	} else {
		CatchError(gunzip.Close())
	}
}

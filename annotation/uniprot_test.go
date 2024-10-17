package annotation

import (
	"testing"
)

func TestUniProtXMLReader(t *testing.T) {
	UniProtXMLReader("testdata/uniprot.xml.gz")
}

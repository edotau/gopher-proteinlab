package annotation

import (
	"strings"
	"testing"
)

func TestParseGenBank(t *testing.T) {
	// Simulate a small portion of a GenBank entry
	genbankData := `
LOCUS       LISOD                    756 bp    DNA     linear   BCT 30-JUN-1993
DEFINITION  Listeria ivanovii sod gene for superoxide dismutase.
ACCESSION   X64011 S78972
VERSION     X64011.1  GI:44010
KEYWORDS    sod gene; superoxide dismutase.
SOURCE      Listeria ivanovii
  ORGANISM  Listeria ivanovii
            Bacteria; Firmicutes; Bacillales; Listeriaceae; Listeria. 
FEATURES             Location/Qualifiers
     CDS             109..717
                     /gene="sod"
                     /product="superoxide dismutase"
ORIGIN      
        1 cgttatttaa ggtgttacat agttctatgg aaatagggtc tatacctttc gccttacaat
       61 gtaatttctt ..........
//
`
	// Expected output
	expectedEntry := &GenBankEntry{
		Locus:      "LISOD                    756 bp    DNA     linear   BCT 30-JUN-1993",
		Definition: "Listeria ivanovii sod gene for superoxide dismutase.",
		Accession:  []string{"X64011", "S78972"},
		Version:    "X64011.1  GI:44010",
		Keywords:   []string{"sod gene", "superoxide dismutase"},
		Source:     "Listeria ivanovii",
		Organism:   "Listeria ivanovii",
		Features: []GenBankFeature{
			{
				Key:      "CDS",
				Location: "109..717",
				Qualifiers: map[string]string{
					"/gene":    "sod",
					"/product": "superoxide dismutase",
				},
			},
		},
		Sequence: "cgttatttaa ggtgttacat agttctatgg aaatagggtc tatacctttc gccttacaat gtaatttctt ..........",
	}

	// Fix the scanner initialization for GenBank
	reader := strings.NewReader(genbankData)

	// Assuming parseGenBank expects *parseio.Scanalyzer
	entry, err := parseGenBank(reader)
	if err != nil {
		t.Fatalf("Unexpected error while parsing GenBank data: %v", err)
	}

	// Compare the parsed entry with the expected entry
	if !EqualGenBankEntry(entry, expectedEntry) {
		t.Errorf("Parsed GenBank entry does not match expected result\nGot:\n%s\nExpected:\n%s", entry.ToString(), expectedEntry.ToString())
	}
}

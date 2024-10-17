package annotation

import (
	"bufio"
	"strings"
	"testing"

	"gopher-proteinlab/parseio"

)


func TestParseEMBL(t *testing.T) {
	// Simulate an EMBL entry using a string
	emblData := `
ID   EMBL000001;
AC   X56734; M74088;
KW   gene; protein; virus;
FT   source          1..1234
FT                   /organism="Homo sapiens"
FT                   /mol_type="mRNA"
FT                   /db_xref="taxon:9606"
FT   CDS             1..1234
FT                   /gene="example_gene"
FT                   /product="example protein"
SQ   Sequence 1234 BP; 614 A; 324 C; 170 G; 126 T; 0 other;
     aaagtttatttagagactaatttgctaaacatattcgcaggcgggcatg
     ttagctatgcggtggcaggttggcactgagctcaggagccggtcgtgcg
//
`
	// Expected result for the EMBLEntry
	expectedEntry := &EMBLEntry{
		ID:        "EMBL000001",
		Accession: []string{"X56734", "M74088"},
		Keywords:  []string{"gene", "protein", "virus"},
		Features: []Feature{
			{
				Key:      "source",
				Location: "1..1234",
				Qualifiers: map[string]string{
					"/organism": "Homo sapiens",
					"/mol_type": "mRNA",
					"/db_xref":  "taxon:9606",
				},
			},
			{
				Key:      "CDS",
				Location: "1..1234",
				Qualifiers: map[string]string{
					"/gene":    "example_gene",
					"/product": "example protein",
				},
			},
		},
		Sequence: "aaagtttatttagagactaatttgctaaacatattcgcaggcgggcatgttagctatgcggtggcaggttggcactgagctcaggagccggtcgtgcg",
	}

	// Use a bufio.Scanner for testing
	emblReader := strings.NewReader(emblData) // Use strings.NewReader directly for string input

// Wrap the bufio.Scanner in Scanalyzer
	emblScanner := &parseio.Scanalyzer{
		Scanner: bufio.NewScanner(emblReader), // Wrap the bufio.Scanner
	}
	// Call the parseEMBL function
	entry, err := parseEMBL(emblScanner)
	if err != nil {
		t.Fatalf("parseEMBL failed: %v", err)
	}

	// Use helper function to compare
	if !EqualEmblEntry(entry, expectedEntry) {
		t.Errorf("Parsed EMBLEntry does not match expected entry.\nParsed: %s\nExpected: %s", entry.ToString(), expectedEntry.ToString())
	}
}
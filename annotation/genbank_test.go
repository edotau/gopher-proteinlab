package annotation

// import (
// 	"bufio"
// 	"gopher-proteinlab/parseio"
// 	"reflect"
// 	"strings"
// 	"testing"
// )

// func TestParseGenBank(t *testing.T) {
// 	// Simulate a small portion of a GenBank entry
// 	genbankData := `
// LOCUS       LISOD                    756 bp    DNA     linear   BCT 30-JUN-1993
// DEFINITION  Listeria ivanovii sod gene for superoxide dismutase.
// ACCESSION   X64011 S78972
// VERSION     X64011.1  GI:44010
// KEYWORDS    sod gene; superoxide dismutase.
// SOURCE      Listeria ivanovii
// 	ORGANISM  Listeria ivanovii
//             Bacteria; Firmicutes; Bacillales; Listeriaceae; Listeria.
// FEATURES             Location/Qualifiers
// 	CDS             109..717
//                     /gene="sod"
//                     /product="superoxide dismutase"
// ORIGIN
//         1 cgttatttaa ggtgttacat agttctatgg aaatagggtc tatacctttc gccttacaat
//        61 gtaatttctt ..........
// //
// `

// 	// Create a scanner for the test data
// 	genbankReader := strings.NewReader(genbankData)
// 	genbankScanner := &parseio.Scanalyzer{
// 		Scanner: bufio.NewScanner(genbankReader), // Wrap the bufio.Scanner
// 	}

// 	// Parse the GenBank entry
// 	entry, err := parseGenBank(genbankScanner)
// 	if err != nil {
// 		t.Fatalf("parseGenBank failed: %v", err)
// 	}

// 	expectedEntry := &GenBankEntry{
// 		Locus:      "LISOD",
// 		Definition: "Listeria ivanovii sod gene for superoxide dismutase.",
// 		Accession:  []string{"X64011", "S78972"},
// 		Version:    "X64011.1  GI:44010",
// 		Keywords:   []string{"sod gene", "superoxide dismutase"},
// 		Source:     "Listeria ivanovii",
// 		Organism:   "Listeria ivanovii Bacteria; Firmicutes; Bacillales; Listeriaceae; Listeria.",
// 		Features: []GenBankFeature{
// 			{
// 				Key:      "CDS",
// 				Location: "109..717",
// 				Qualifiers: map[string]string{
// 					"/gene":    "sod",
// 					"/product": "superoxide dismutase",
// 				},
// 			},
// 		},
// 		Sequence: "cgttatttaaggtgttacatagttctatggaaatagggtctatacctttcgccttacaatgtaatttctt..........",
// 	}

// 	if !reflect.DeepEqual(entry, expectedEntry) {
// 		t.Errorf("Parsed entry does not match expected entry.\nParsed: %s\nExpected: %s", entry.ToJson(), expectedEntry.ToJson())
// 	}
// }

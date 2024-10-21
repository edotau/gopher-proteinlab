package uniprot

import (
	"testing"
)

// TestUniProtEntryToJSON tests the ToJson method for the UniProtEntry struct.
func TestUniProtEntryToJSON(t *testing.T) {
	entry := Entry{
		Accession: []string{"P12345", "Q67890"},
		Name:      []string{"Example Protein"},
		Protein: Protein{
			RecommendedName: RecommendedName{
				FullName: EvidencedString{Value: "Example Protein Full Name"},
			},
		},
		Organism: Organism{
			Name: []OrganismName{
				{Name: "scientific", {{Type: "Homo sapiens", Id: "1"}},
			},
			
		},
		Sequence: Sequence{
			Value:    "MSEQENCE",
			Length:   8,
			Mass:     1234,
			Checksum: "ABCD1234",
			Modified: "2024-10-17",
			Version:  1,
		},
		Created:  "2024-10-17",
		Modified: "2024-10-17",
		Version:  1,
		References: []Reference{
			{Citation: Citation{Title: "Reference 1"}},
			{Citation: Citation{Title: "Reference 2"}},
		},
	}

	expectedOutput := "{\n  \"Accession\": [\n    \"P12345\",\n    \"Q67890\"\n  ],\n  \"Name\": [\n    \"Example Protein\"\n  ],\n  \"Protein\": {\n    \"RecommendedName\": {\n      \"FullName\": \"Example Protein Full Name\"\n    }\n  },\n  \"Organism\": {\n    \"Name\": [\n      {\n        \"Type\": \"scientific\",\n        \"Name\": \"Homo sapiens\"\n      },\n      {\n        \"Type\": \"common\",\n        \"Name\": \"Human\"\n      }\n    ]\n  },\n  \"Sequence\": {\n    \"Value\": \"MSEQENCE\",\n    \"Length\": 8,\n    \"Mass\": 1234,\n    \"Checksum\": \"ABCD1234\",\n    \"Modified\": \"2024-10-17\",\n    \"Version\": 1\n  },\n  \"Created\": \"2024-10-17\",\n  \"Modified\": \"2024-10-17\",\n  \"Version\": 1,\n  \"References\": [\n    {\n      \"Key\": \"\",\n      \"Citation\": {\n        \"Title\": \"Reference 1\",\n        \"Type\": \"\"\n      }\n    },\n    {\n      \"Key\": \"\",\n      \"Citation\": {\n        \"Title\": \"Reference 2\",\n        \"Type\": \"\"\n      }\n    }\n  ]\n}"

	output := entry.ToJson()

	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

// // TestUniProtEntryToString tests the ToString method for the UniProtEntry struct.
// func TestUniProtEntryToString(t *testing.T) {
// 	entry := Entry{
// 		Accession: []string{"P12345", "Q67890"},
// 		Name:      []string{"Example Protein"},
// 		Protein: Protein{
// 			RecommendedName: RecommendedName{
// 				FullName: EvidencedString{Value: "Example Protein Full Name"},
// 			},
// 		},
// 		Organism: Organism{
// 			Names: []OrganismName{
// 				{Type: "scientific", Value: "Homo sapiens"},
// 				{Type: "common", Value: "Human"},
// 			},
// 		},
// 		Sequence: Sequence{
// 			Value:    "MSEQENCE",
// 			Length:   8,
// 			Mass:     1234,
// 			Checksum: "ABCD1234",
// 			Modified: "2024-10-17",
// 			Version:  1,
// 		},
// 		Created:  "2024-10-17",
// 		Modified: "2024-10-17",
// 		Version:  1,
// 		References: []Reference{
// 			{Citation: Citation{Title: "Reference 1"}},
// 			{Citation: Citation{Title: "Reference 2"}},
// 		},
// 	}

// 	expectedOutput := `Accession: P12345, Q67890
// Name: Example Protein
// Protein: Example Protein Full Name
// Organism: scientific: Homo sapiens; common: Human;
// Sequence: MSEQENCE
// Created: 2024-10-17
// Modified: 2024-10-17
// Version: 1
// References: Reference 1; Reference 2;
// `
// 	output := (entry.ToJson())
// 	if output != expectedOutput {
// 		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOutput, output)
// 	}
// }

// // TestUniProtXMLReader tests the UniProtXMLReader function for reading XML files.
// func TestUniProtXMLReader(t *testing.T) {
// 	expectedEntry := &Entry{
// 		Accession: []string{"P0C9F0"},
// 		Name:      []string{"1001R_ASFK5"},
// 		Protein: Protein{
// 			RecommendedName: RecommendedName{
// 				FullName: EvidencedString{Value: "Protein MGF 100-1R"},
// 			},
// 		},
// 		Organism: Organism{
// 			Names: []OrganismName{
// 				{Type: "scientific", Value: "African swine fever virus (isolate Pig/Kenya/KEN-50/1950)"},
// 				{Type: "common", Value: "ASFV"},
// 			},
// 		},
// 		Sequence: Sequence{
// 			Value:    "MVRLFYNPIKYLFYRRSCKKRLRKALKKLNFYHPPKECCQIYRLLENAPGGTYFITENMTNELIMIAKDPVDKKIKSVKLYLTGNYIKINQHYYINIYMYLMRYNQIYKYPLICFSKYSKIL",
// 			Length:   122,
// 			Mass:     14969,
// 			Checksum: "C5E63C34B941711C",
// 			Modified: "2009-05-05",
// 			Version:  1,
// 		},
// 		Created:  "2009-05-05",
// 		Modified: "2023-11-08",
// 		Version:  11,
// 		References: []Reference{
// 			{
// 				Key: "1",
// 				Citation: Citation{
// 					Title: "African swine fever virus genomes.",
// 					Type:  "submission",
// 				},
// 			},
// 		},
// 	}
// 	xmlReader, xmlFile := parseio.FileHandler("testdata/uniprot.xml.gz")
// 	defer xmlFile.Close()
// 	decoder := xml.NewDecoder(xmlReader)

// 	entry, err := parseUniProt(decoder)
// 	parseio.ExitOnError(err)

// 	if !EqualEntries(entry, expectedEntry) {
// 		t.Fatalf("Entries are not equal\n%s\n!=\n %s\n", entry.ToString(), expectedEntry.ToString())
// 	}
// }

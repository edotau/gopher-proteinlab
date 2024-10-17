package annotation

import (
	"testing"
)

func TestEntryToJSON(t *testing.T) {
	entry := Entry{
		Accession: []string{"P12345", "Q67890"},
		Name:      []string{"Example Protein"},
		Protein: ProteinInfo{
			RecommendedName: RecommendedName{
				FullName: "Example Protein Full Name",
			},
		},
		Organism: Organism{
			Name: []OrganismName{
				{Type: "scientific", Name: "Homo sapiens"},
				{Type: "common", Name: "Human"},
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

func TestEntryToString(t *testing.T) {
	entry := Entry{
		Accession: []string{"P12345", "Q67890"},
		Name:      []string{"Example Protein"},
		Protein: ProteinInfo{
			RecommendedName: RecommendedName{
				FullName: "Example Protein Full Name",
			},
		},
		Organism: Organism{
			Name: []OrganismName{
				{Type: "scientific", Name: "Homo sapiens"},
				{Type: "common", Name: "Human"},
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

	expectedOutput := `Accession: P12345, Q67890
Name: Example Protein
Protein: Example Protein Full Name
Organism: scientific: Homo sapiens; common: Human; 
Sequence: MSEQENCE
Created: 2024-10-17
Modified: 2024-10-17
Version: 1
References: Reference 1; Reference 2; 
`
	output := (entry.ToString())
	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

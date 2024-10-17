package annotation

import (
	"testing"
)

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

	output := entry.ToString()

	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

package uniprot

import (
	"encoding/xml"
	"gopher-proteinlab/parseio"
	"testing"
)

func TestUniProtProteinEntry(t *testing.T) {
	data := `
	<protein>
		<recommendedName>
			<fullName>Example Protein Full Name</fullName>
			<shortName>Example Short Name</shortName>
			<ecNumber>1.1.1.1</ecNumber>
		</recommendedName>
		<alternativeName>
			<fullName>Alternative Full Name</fullName>
			<shortName>Alt Short Name</shortName>
		</alternativeName>
		<submittedName>
			<fullName>Submitted Full Name</fullName>
			<ecNumber>2.2.2.2</ecNumber>
		</submittedName>
	</protein>
	`

	expected := ProteinEntry{
		RecommendedName: &ProteinName{
			FullName:  NameEntry{Value: "Example Protein Full Name"},
			ShortName: []NameEntry{{Value: "Example Short Name"}},
			ECNumber:  []NameEntry{{Value: "1.1.1.1"}},
		},
		AlternativeName: []ProteinName{
			{
				FullName:  NameEntry{Value: "Alternative Full Name"},
				ShortName: []NameEntry{{Value: "Alt Short Name"}},
			},
		},
		SubmittedName: []ProteinName{
			{
				FullName: NameEntry{Value: "Submitted Full Name"},
				ECNumber: []NameEntry{{Value: "2.2.2.2"}},
			},
		},
	}

	var protein ProteinEntry
	err := xml.Unmarshal([]byte(data), &protein)
	parseio.ExitOnError(err)

	if !protein.Equal(expected) {
		t.Errorf("Unmarshaled data does not match expected data.\nGot: %s\nExpected: %s", protein.ToJson(), expected.ToJson())
	}
}

package annotation

import (
	"encoding/xml"
	"fmt"
	"gopher-proteinlab/parseio"
	"testing"
)

func TestUniProtXMLReader(t *testing.T) {
	expectedEntry := &Entry{
		Accession: []string{"P0C9F0"},
		Name:      []string{"1001R_ASFK5"},
		Protein: ProteinInfo{
			RecommendedName: RecommendedName{
				FullName: "Protein MGF 100-1R",
			},
		},
		Organism: Organism{
			Name: []OrganismName{
				{Type: "scientific", Name: "African swine fever virus (isolate Pig/Kenya/KEN-50/1950)"},
				{Type: "common", Name: "ASFV"},
			},
		},
		Sequence: Sequence{
			Value:    "MVRLFYNPIKYLFYRRSCKKRLRKALKKLNFYHPPKECCQIYRLLENAPGGTYFITENMTNELIMIAKDPVDKKIKSVKLYLTGNYIKINQHYYINIYMYLMRYNQIYKYPLICFSKYSKIL",
			Length:   122,
			Mass:     14969,
			Checksum: "C5E63C34B941711C",
			Modified: "2009-05-05",
			Version:  1,
		},
		Created:  "2009-05-05",
		Modified: "2023-11-08",
		Version:  11,
		References: []Reference{
			{
				Key: "1",
				Citation: Citation{
					Title: "African swine fever virus genomes.",
					Type:  "submission",
				},
			},
		},
	}
	xmlReader, xmlFile := parseio.FileHandler("testdata/uniprot.xml.gz")
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlReader)

	entry, err := parseUniProt(decoder)
	parseio.ExitOnError(err)

	if !EqualEntries(entry, expectedEntry) {
		fmt.Printf("Entries are not equal\n%s\n!=\n %s\n", entry.ToString(), expectedEntry.ToString())
	}
}

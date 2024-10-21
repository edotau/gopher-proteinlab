package uniprot

import (
	"encoding/xml"
	"testing"
)

// TestOrganismXMLParsing tests the XML parsing of the Organism struct
func TestOrganismXMLParsing(t *testing.T) {
	xmlData := `
	<organism evidence="1">
		<name type="scientific">Homo sapiens</name>
		<name type="common">Human</name>
		<dbReference type="NCBI Taxonomy" id="9606"/>
		<lineage>
			<taxon>Eukaryota</taxon>
			<taxon>Metazoa</taxon>
			<taxon>Chordata</taxon>
			<taxon>Mammalia</taxon>
			<taxon>Primates</taxon>
			<taxon>Hominidae</taxon>
			<taxon>Homo</taxon>
			<taxon>Homo sapiens</taxon>
		</lineage>
	</organism>`

	expected := Organism{
		Evidence: "1",
		Name: []NameEntry{
			{Type: "scientific", Value: "Homo sapiens"},
			{Type: "common", Value: "Human"},
		},
		DBReference: []DBReference{
			{Type: "NCBI Taxonomy", ID: "9606"},
		},
		Lineage: &Lineage{
			Taxon: []string{
				"Eukaryota", "Metazoa", "Chordata", "Mammalia", "Primates",
				"Hominidae", "Homo", "Homo sapiens",
			},
		},
	}

	var actual Organism
	err := xml.Unmarshal([]byte(xmlData), &actual)
	if err != nil {
		t.Fatalf("Error unmarshaling XML: %v", err)
	}

	if !actual.Equal(expected) {
		t.Errorf("Unmarshaled data does not match expected.\nGot: %+v\nExpected: %+v", actual, expected)
	}
}

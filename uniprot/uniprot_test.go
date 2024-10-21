package uniprot

import (
	"encoding/xml"
	"gopher-proteinlab/parseio"
	"testing"
)

// TestUniProtXMLReader tests the UniProtXMLReader function by comparing XML strings.
func TestUniProtXMLReader(t *testing.T) {
	// Parse the date strings to time.Time format
	expectedEntry := Entry{
		Accession: "P0C9F0",
		Name:      "1001R_ASFK5",
		Protein: ProteinEntry{
			RecommendedName: ProteinName{
				FullName: NameEntry{Value: "Protein MGF 100-1R"},
			},
		},
		Gene: []Gene{
			{
				Name: []NameEntry{
					{
						Type:  "ordered locus",
						Value: "Ken-018",
					},
				},
			},
		},
		Organism: Organism{
			Name: []NameEntry{
				{Type: "scientific", Value: "African swine fever virus (isolate Pig/Kenya/KEN-50/1950)"},
				{Type: "common", Value: "ASFV"},
			},
			DBReference: []DBReference{
				{
					Type: "NCBI Taxonomy",
					ID:   "561445",
				},
			},
			Lineage: &Lineage{
				Taxon: []string{
					"Viruses", "Varidnaviria", "Bamfordvirae", "Nucleocytoviricota",
					"Pokkesviricetes", "Asfuvirales", "Asfarviridae", "Asfivirus",
					"African swine fever virus",
				},
			},
		},
		OrganismHost: []Organism{
			{
				Name: []NameEntry{
					{Type: "scientific", Value: "Ornithodoros"},
					{Type: "common", Value: "relapsing fever ticks"},
				},
				DBReference: []DBReference{
					{Type: "NCBI Taxonomy", ID: "6937"},
				},
			},
			{
				Name: []NameEntry{
					{Type: "scientific", Value: "Phacochoerus aethiopicus"},
					{Type: "common", Value: "Warthog"},
				},
				DBReference: []DBReference{
					{Type: "NCBI Taxonomy", ID: "85517"},
				},
			},
			{
				Name: []NameEntry{
					{Type: "scientific", Value: "Phacochoerus africanus"},
					{Type: "common", Value: "Warthog"},
				},
				DBReference: []DBReference{
					{Type: "NCBI Taxonomy", ID: "41426"},
				},
			},
			{
				Name: []NameEntry{
					{Type: "scientific", Value: "Potamochoerus larvatus"},
					{Type: "common", Value: "Bushpig"},
				},
				DBReference: []DBReference{
					{Type: "NCBI Taxonomy", ID: "273792"},
				},
			},
			{
				Name: []NameEntry{
					{Type: "scientific", Value: "Sus scrofa"},
					{Type: "common", Value: "Pig"},
				},
				DBReference: []DBReference{
					{Type: "NCBI Taxonomy", ID: "9823"},
				},
			},
		},
		References: []Reference{
			{
				Key: "1",
				Citation: Citation{
					Title: "African swine fever virus genomes.",
					Type:  "submission",
					Date:  "2003-03",
					AuthorList: []Person{
						{Name: "Kutish G.F."},
						{Name: "Rock D.L."},
					},
				},
				Scope: []string{"NUCLEOTIDE SEQUENCE [LARGE SCALE GENOMIC DNA]"},
			},
		},
		Comment: []Comment{
			{
				Type: "function",
				Text: []NameEntry{
					{Value: "Plays a role in virus cell tropism, and may be required for efficient virus replication in macrophages.", Evidence: []int{1}},
				},
			},
			{
				Type: "similarity",
				Text: []NameEntry{
					{Value: "Belongs to the asfivirus MGF 100 family.", Evidence: []int{2}},
				},
			},
		},
		DBReference: []DBReference{
			{
				Type: "EMBL",
				ID:   "AY261360",
				Property: []Property{
					{Type: "status", Value: "NOT_ANNOTATED_CDS"},
					{Type: "molecule type", Value: "Genomic_DNA"},
				},
			},
			{Type: "SMR", ID: "P0C9F0"},
			{
				Type: "Proteomes",
				ID:   "UP000000861",
				Property: []Property{
					{Type: "component", Value: "Genome"},
				},
			},
		},
		ProteinExistence: ProteinExistence{Type: "inferred from homology"},
		Feature: []Feature{
			{
				Type:        "chain",
				ID:          "PRO_0000373170",
				Description: "Protein MGF 100-1R",
				Location: Location{
					Begin: &Position{Position: 1},
					End:   &Position{Position: 122},
				},
			},
		},
		Evidence: []Evidence{
			{Type: "ECO:0000250", Key: 1},
			{Type: "ECO:0000305", Key: 2},
		},
		Sequence: Sequence{
			Value:    "MVRLFYNPIKYLFYRRSCKKRLRKALKKLNFYHPPKECCQIYRLLENAPGGTYFITENMTNELIMIAKDPVDKKIKSVKLYLTGNYIKINQHYYINIYMYLMRYNQIYKYPLICFSKYSKIL",
			Length:   122,
			Mass:     14969,
			Checksum: "C5E63C34B941711C",
			Modified: "2009-05-05",
			Version:  1,
		},
		Dataset: "Swiss-Prot",

		Created:  "2009-05-05",
		Modified: "2023-11-08",
		Version:  11,
	}
	// Open and read the XML file
	xmlReader := parseio.NewCodeReader("testdata/uniprot.xml.gz")
	defer xmlReader.Close()

	decoder := xml.NewDecoder(xmlReader)
	var uniprot Uniprot
	err := decoder.Decode(&uniprot)
	parseio.ExitOnError(err)

	entry := uniprot.Entries[0]

	if entry.ToJson() != expectedEntry.ToJson() {
		t.Fatalf("Entries are not equal\nExpected:\n%s\nActual:\n%s\n", entry.ToJson(), expectedEntry.ToJson())
	}
}

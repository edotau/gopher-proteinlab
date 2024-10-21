package uniprot

// import (
// 	"encoding/xml"
// 	"testing"
// )

// // Unit test to unmarshal the XML and validate the struct data.
// func TestUniProtProteinUnmarshal(t *testing.T) {
// 	data := `
// 	<protein>
// 		<recommendedName>
// 			<fullName>Example Protein Full Name</fullName>
// 			<shortName>Example Short Name</shortName>
// 			<ecNumber>1.1.1.1</ecNumber>
// 		</recommendedName>
// 		<alternativeName>
// 			<fullName>Alternative Full Name</fullName>
// 			<shortName>Alt Short Name</shortName>
// 		</alternativeName>
// 		<submittedName>
// 			<fullName>Submitted Full Name</fullName>
// 			<ecNumber>2.2.2.2</ecNumber>
// 		</submittedName>
// 		<allergenName>Allergen Example</allergenName>
// 		<cdAntigenName>CD Antigen 1</cdAntigenName>
// 		<cdAntigenName>CD Antigen 2</cdAntigenName>
// 	</protein>
// 	`

// 	var protein Protein
// 	err := xml.Unmarshal([]byte(data), &protein)
// 	if err != nil {
// 		t.Fatalf("Error unmarshaling XML: %v", err)
// 	}

// 	// Expected output for comparison
// 	expectedProtein := Protein{
// 		Entry: ProteinName{
// 			FullName: Citation{Title: "Example Protein Full Name"},
// 			ShortNames: []Citation{
// 				{Title: "Example Short Name"},
// 			},
// 			ECNumbers: []Citation{
// 				{Title: "1.1.1.1"},
// 			},
// 		},
// 	}

// 	// Compare the unmarshaled data with the expected result
// 	if !protein.Equal(expectedProtein) {
// 		t.Errorf("Unmarshaled data does not match expected data.\nGot: %+v\nExpected: %+v", protein, expectedProtein)
// 	}
// }

// func TestUniProtGeneUnmarshal(t *testing.T) {
// 	data := `
// 	<protein>
// 		<recommendedName>
// 			<fullName>Example Protein Full Name</fullName>
// 			<shortName>Example Short Name</shortName>
// 			<ecNumber>1.1.1.1</ecNumber>
// 		</recommendedName>
// 		<alternativeName>
// 			<fullName>Alternative Full Name</fullName>
// 			<shortName>Alt Short Name</shortName>
// 		</alternativeName>
// 		<submittedName>
// 			<fullName>Submitted Full Name</fullName>
// 			<ecNumber>2.2.2.2</ecNumber>
// 		</submittedName>
// 		<allergenName>Allergen Example</allergenName>
// 		<cdAntigenName>CD Antigen 1</cdAntigenName>
// 		<cdAntigenName>CD Antigen 2</cdAntigenName>
// 		<gene>
// 			<name type="primary">Gene Primary Name</name>
// 			<name type="synonym">Gene Synonym Name</name>
// 		</gene>
// 	</protein>
// 	`

// 	var protein Protein
// 	err := xml.Unmarshal([]byte(data), &protein)
// 	if err != nil {
// 		t.Fatalf("Error unmarshaling XML: %v", err)
// 	}

// 	expectedProtein := Protein{
// 		RecommendedName: RecommendedName{
// 			FullName:  EvidencedString{Value: "Example Protein Full Name"},
// 			ShortName: []EvidencedString{{Value: "Example Short Name"}},
// 			EcNumber:  []EvidencedString{{Value: "1.1.1.1"}},
// 		},
// 		AlternativeNames: []AlternativeName{
// 			{
// 				FullName:  EvidencedString{Value: "Alternative Full Name"},
// 				ShortName: []EvidencedString{{Value: "Alt Short Name"}},
// 			},
// 		},
// 		SubmittedNames: []SubmittedName{
// 			{
// 				FullName: EvidencedString{Value: "Submitted Full Name"},
// 				EcNumber: []EvidencedString{{Value: "2.2.2.2"}},
// 			},
// 		},
// 		AllergenName:   "Allergen Example",
// 		CdAntigenNames: []string{"CD Antigen 1", "CD Antigen 2"},
// 		Genes: []Gene{
// 			{
// 				Names: []GeneName{
// 					{Value: "Gene Primary Name", Type: "primary"},
// 					{Value: "Gene Synonym Name", Type: "synonym"},
// 				},
// 			},
// 		},
// 	}

// 	// Ignore the Name field in comparison
// 	protein.Name = xml.Name{}
// 	expectedProtein.Name = xml.Name{}

// 	if !protein.Equal(expectedProtein) {
// 		t.Errorf("Unmarshaled data does not match expected data.\nGot: %+v\nExpected: %+v", protein, expectedProtein)
// 	}
// }

// func TestOrganismUnmarshal(t *testing.T) {
// 	data := `
// 	<organism evidence="1">
// 		<name type="scientific">Homo sapiens</name>
// 		<name type="common">Human</name>
// 		<dbReference type="NCBI Taxonomy" id="9606"/>
// 		<lineage>
// 			<taxon>Eukaryota</taxon>
// 			<taxon>Metazoa</taxon>
// 			<taxon>Chordata</taxon>
// 		</lineage>
// 	</organism>
// 	`

// 	var organism Organism
// 	err := xml.Unmarshal([]byte(data), &organism)
// 	if err != nil {
// 		t.Fatalf("Error unmarshaling XML: %v", err)
// 	}

// 	expectedOrganism := Organism{
// 		Names: []OrganismName{
// 			{Value: "Homo sapiens", Type: "scientific"},
// 			{Value: "Human", Type: "common"},
// 		},
// 		DbReference: []DbReference{
// 			{Type: "NCBI Taxonomy", ID: "9606"},
// 		},
// 		Lineage: &Lineage{
// 			Taxa: []string{"Eukaryota", "Metazoa", "Chordata"},
// 		},
// 		Evidence: "1",
// 	}

// 	// Ignore the Name field in comparison
// 	organism.Name = xml.Name{}
// 	expectedOrganism.Name = xml.Name{}

// 	// Compare Lineage values, not pointers
// 	if !reflect.DeepEqual(organism.Lineage.Taxa, expectedOrganism.Lineage.Taxa) {
// 		t.Errorf("Unmarshaled Lineage does not match expected data.\nGot: %+v\nExpected: %+v", organism.Lineage.Taxa, expectedOrganism.Lineage.Taxa)
// 	}

// 	// Compare the rest of the structure
// 	if !reflect.DeepEqual(organism, expectedOrganism) {
// 		t.Errorf("Unmarshaled data does not match expected data.\nGot: %+v\nExpected: %+v", organism, expectedOrganism)
// 	}
// }

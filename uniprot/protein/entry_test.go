package uniprot

// import "testing"

// func TestProteinEntryToStr(t *testing.T) {
// 	p := &Protein{
// 		RecommendedName: RecommendedName{
// 			FullName: EvidencedString{Value: "Hemoglobin subunit alpha"},
// 		},
// 		Genes: []Gene{
// 			{Names: []GeneName{{Value: "HBA1"}}},
// 			{Names: []GeneName{{Value: "HBA2"}}},
// 		},
// 		AlternativeNames: []AlternativeName{
// 			{FullName: EvidencedString{Value: "Alpha-1-globin"}},
// 		},
// 		CdAntigenNames: []string{"CD235a"},
// 	}
// 	expected := `Protein: Hemoglobin subunit alpha
// Genes: HBA1, HBA2
// Alternative Names: Alpha-1-globin
// CD Antigen Names: CD235a
// `
// 	s := p.ToString()

// 	if s != expected {
// 		t.Fatalf("Error: protein.ToString()\n %s\n%s\n", s, expected)
// 	}
// }

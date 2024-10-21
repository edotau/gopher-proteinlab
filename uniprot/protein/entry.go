package uniprot

import (
	"encoding/json"
	"gopher-proteinlab/parseio"
	"strings"
)

// ToJson converts a UniProtEntry to a JSON-formatted string.
func (e *Entry) ToJson() string {
	txt := parseio.NewTxtBuilder()
	data, err := json.MarshalIndent(e, "", "  ")
	if parseio.ExitOnError(err) {
		txt.Write(data)
	}
	return txt.String()
}

// func (p *Protein) ToString() string {
// 	txtBuilder := parseio.NewTxtBuilder()
// 	txtBuilder.WriteTag("Protein", p.RecommendedName.FullName.Value)

// 	if len(p.Genes) > 0 {
// 		txtBuilder.WriteString("Genes: ")
// 		for i, gene := range p.Genes {
// 			txtBuilder.WriteString(gene.Names[0].Value)
// 			if i < len(p.Genes)-1 {
// 				txtBuilder.WriteString(", ")
// 			}
// 		}
// 		txtBuilder.WriteByte('\n')
// 	}

// 	if len(p.AlternativeNames) > 0 {
// 		txtBuilder.WriteTag("Alternative Names", strings.Join(collectNames(p.AlternativeNames), ", "))
// 	}

// 	if len(p.SubmittedNames) > 0 {
// 		txtBuilder.WriteTag("Submitted Names", strings.Join(collectNames(p.SubmittedNames), ", "))
// 	}

// 	if p.AllergenName != "" {
// 		txtBuilder.WriteTag("Allergen Name", p.AllergenName)
// 	}

// 	if p.BiotechName != "" {
// 		txtBuilder.WriteTag("Biotech Name", p.BiotechName)
// 	}

// 	if len(p.CdAntigenNames) > 0 {
// 		txtBuilder.WriteTag("CD Antigen Names", strings.Join(p.CdAntigenNames, ", "))
// 	}

// 	if len(p.InnNames) > 0 {
// 		txtBuilder.WriteTag("INN Names", strings.Join(p.InnNames, ", "))
// 	}

// 	return txtBuilder.String()
// }



// // ToString returns a string representation of the UniProtEntry struct.
func (e Entry) ToString() string {
	var words strings.Builder

	// 	writeField(&words, "Accession: ", strings.Join(e.Accession, ", "))
	// 	writeField(&words, "Name: ", strings.Join(e.Name, ", "))
	// 	writeField(&words, "Protein: ", e.Protein.RecommendedName.FullName.Value)

	// 	parseio.HandleStrBuilder(&words, "Organism: ")
	// 	for _, name := range e.Organism.Name {
	// 		parseio.HandleStrBuilder(&words, name.Type)
	// 		parseio.HandleStrBuilder(&words, ": ")
	// 		parseio.HandleStrBuilder(&words, name.Name)
	// 		parseio.HandleStrBuilder(&words, "; ")
	// 	}
	// 	parseio.ExitOnError(words.WriteByte('\n'))

	// 	writeField(&words, "Sequence: ", e.Sequence.Value)
	// 	writeField(&words, "Created: ", e.Created)
	// 	writeField(&words, "Modified: ", e.Modified)
	// 	writeField(&words, "Version: ", fmt.Sprintf("%d", e.Version))

	// 	parseio.HandleStrBuilder(&words, "References: ")
	// 	for _, ref := range e.References {
	// 		parseio.HandleStrBuilder(&words, ref.Citation.Title)
	// 		parseio.HandleStrBuilder(&words, "; ")
	// 	}
	// 	parseio.ExitOnError(words.WriteByte('\n'))
	return words.String()
}

package annotation

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopher-proteinlab/parseio"
)

// Entry defines the interface for parsing biological entries like UniProt and EMBL.
type Entry interface {
	ToJson() string
	ToString() string
}

// ToJson converts a UniProtEntry to a JSON-formatted string.
func (e *UniProtEntry) ToJson() string {
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v", err)
	}
	return string(data)
}

// ToJson converts an EMBLEntry to a JSON-formatted string.
func (e *EMBLEntry) ToJson() string {
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v", err)
	}
	return string(data)
}

// ToJson converts a GenBankEntry to a JSON-formatted string.
func (e *GenBankEntry) ToJson() string {
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v", err)
	}
	return string(data)
}

// ToString returns a string representation of the UniProtEntry struct.
func (e UniProtEntry) ToString() string {
	var words strings.Builder

	writeField(&words, "Accession: ", strings.Join(e.Accession, ", "))
	writeField(&words, "Name: ", strings.Join(e.Name, ", "))
	writeField(&words, "Protein: ", e.Protein.RecommendedName.FullName.Value)

	parseio.HandleStrBuilder(&words, "Organism: ")
	for _, name := range e.Organism.Name {
		parseio.HandleStrBuilder(&words, name.Type)
		parseio.HandleStrBuilder(&words, ": ")
		parseio.HandleStrBuilder(&words, name.Name)
		parseio.HandleStrBuilder(&words, "; ")
	}
	parseio.ExitOnError(words.WriteByte('\n'))

	writeField(&words, "Sequence: ", e.Sequence.Value)
	writeField(&words, "Created: ", e.Created)
	writeField(&words, "Modified: ", e.Modified)
	writeField(&words, "Version: ", fmt.Sprintf("%d", e.Version))

	parseio.HandleStrBuilder(&words, "References: ")
	for _, ref := range e.References {
		parseio.HandleStrBuilder(&words, ref.Citation.Title)
		parseio.HandleStrBuilder(&words, "; ")
	}
	parseio.ExitOnError(words.WriteByte('\n'))
	return words.String()
}

// ToString converts an EMBLEntry to a string.
func (e EMBLEntry) ToString() string {
	var sb strings.Builder
	writeField(&sb, "ID: ", e.ID)
	writeField(&sb, "Accession: ", strings.Join(e.Accession, ", "))
	writeField(&sb, "Keywords: ", strings.Join(e.Keywords, ", "))
	writeField(&sb, "Source: ", e.Source)
	sb.WriteString("Features:\n")
	for _, feature := range e.Features {
		writeField(&sb, "  Key: ", feature.Key)
		writeField(&sb, "  Location: ", feature.Location)
		
		for key, value := range feature.Qualifiers {
			writeField(&sb, fmt.Sprintf("    %s: ", key), value)
		}
	}
	writeField(&sb, "Sequence: ", e.Sequence)
	return sb.String()
}

// ToString returns a string representation of the GenBankEntry struct.
func (e GenBankEntry) ToString() string {
	var sb strings.Builder

	// Helper function to write fields into the string builder
	writeField := func(label, value string) {
		sb.WriteString(label)
		sb.WriteString(value)
		sb.WriteByte('\n')
	}

	writeField("Locus: ", e.Locus)
	writeField("Definition: ", e.Definition)
	writeField("Accession: ", strings.Join(e.Accession, ", "))
	writeField("Version: ", e.Version)
	writeField("Keywords: ", strings.Join(e.Keywords, ", "))
	writeField("Source: ", e.Source)
	writeField("Organism: ", e.Organism)

	// Write features
	sb.WriteString("Features:\n")
	for _, feature := range e.Features {
		writeField("  Key: ", feature.Key)
		writeField("  Location: ", feature.Location)
		for key, value := range feature.Qualifiers {
			writeField(fmt.Sprintf("    %s: ", key), value)
		}
	}

	writeField("Sequence: ", e.Sequence)

	return sb.String()
}

// writeField is a helper function to write a label and its corresponding value to a string builder.
func writeField(buffer *strings.Builder, label, value string) {
	parseio.HandleStrBuilder(buffer, label)
	parseio.HandleStrBuilder(buffer, value)
	parseio.ExitOnError(buffer.WriteByte('\n'))
}

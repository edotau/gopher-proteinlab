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
	parseio.NewTxtBuilder().WriteTag(label, value)
}

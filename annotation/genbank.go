package annotation

import (
	"bufio"
	"fmt"
	"gopher-proteinlab/parseio"
	"io"
	"os"
	"strings"
)

// GenBankEntry represents the structure of a GenBank file entry.
type GenBankEntry struct {
	Locus      string
	Definition string
	Accession  []string
	Version    string
	Keywords   []string
	Source     string
	Organism   string
	References []GenBankReference
	Features   []GenBankFeature
	Sequence   string
}

// GenBankFeature represents a feature in a GenBank file.
type GenBankFeature struct {
	Key        string
	Location   string
	Qualifiers map[string]string
}

// GenBankReference represents a reference in a GenBank file.
type GenBankReference struct {
	Number  string
	Authors string
	Title   string
	Journal string
	Medline string
	Comment string
}

// GenBankReader reads a GenBank .gbff file, decodes its content, and returns the parsed data.
func GenBankReader(filename string) {
	file, err := os.Open(filename)
	parseio.ExitOnError(err)

	defer file.Close()

	for {
		// Parse each entry and process it
		entry, err := parseGenBank(file)
		if err == io.EOF {
			break // End of file reached
		}
		parseio.ExitOnError(err)
		// TODO: Data transformations go here
		fmt.Println(entry.ToString())
	}
}

// parseGenBank reads and parses one GenBank entry at a time.
func parseGenBank(r io.Reader) (*GenBankEntry, error) {
	scanner := bufio.NewScanner(r)
	entry := &GenBankEntry{}
	currentFeature := GenBankFeature{}
	isSequence := false

	for scanner.Scan() {
		line := scanner.Text()

		// LOCUS line
		if strings.HasPrefix(line, "LOCUS") {
			entry.Locus = strings.TrimSpace(line[12:])
		}

		// DEFINITION line
		if strings.HasPrefix(line, "DEFINITION") {
			entry.Definition = strings.TrimSpace(line[12:])
		}

		// ACCESSION line
		if strings.HasPrefix(line, "ACCESSION") {
			accessions := strings.Split(strings.TrimSpace(line[12:]), " ")
			for _, acc := range accessions {
				entry.Accession = append(entry.Accession, strings.TrimSpace(acc))
			}
		}

		// VERSION line
		if strings.HasPrefix(line, "VERSION") {
			entry.Version = strings.TrimSpace(line[12:])
		}

		// KEYWORDS line
		if strings.HasPrefix(line, "KEYWORDS") {
			keywords := strings.Split(strings.TrimSpace(line[12:]), ";")
			for _, keyword := range keywords {
				entry.Keywords = append(entry.Keywords, strings.TrimSpace(keyword))
			}
		}

		// SOURCE line
		if strings.HasPrefix(line, "SOURCE") {
			entry.Source = strings.TrimSpace(line[12:])
		}

		// ORGANISM line
		if strings.HasPrefix(line, "  ORGANISM") {
			entry.Organism = strings.TrimSpace(line[12:])
		}

		// REFERENCE lines
		if strings.HasPrefix(line, "REFERENCE") {
			ref := GenBankReference{
				Number: strings.TrimSpace(line[12:]),
			}
			entry.References = append(entry.References, ref)
		}

		// AUTHORS line within REFERENCE
		if strings.HasPrefix(line, "  AUTHORS") {
			entry.References[len(entry.References)-1].Authors = strings.TrimSpace(line[12:])
		}

		// TITLE line within REFERENCE
		if strings.HasPrefix(line, "  TITLE") {
			entry.References[len(entry.References)-1].Title = strings.TrimSpace(line[12:])
		}

		// JOURNAL line within REFERENCE
		if strings.HasPrefix(line, "  JOURNAL") {
			entry.References[len(entry.References)-1].Journal = strings.TrimSpace(line[12:])
		}

		// MEDLINE line within REFERENCE
		if strings.HasPrefix(line, "  MEDLINE") {
			entry.References[len(entry.References)-1].Medline = strings.TrimSpace(line[12:])
		}

		// FEATURES section
		if strings.HasPrefix(line, "FEATURES") {
			isSequence = false
		}

		// Parse features
		if strings.HasPrefix(line, "     ") && strings.HasPrefix(strings.TrimSpace(line), "/") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				qualifierKey := strings.TrimSpace(parts[0])
				qualifierValue := strings.Trim(strings.TrimSpace(parts[1]), "\"")
				currentFeature.Qualifiers[qualifierKey] = qualifierValue
			}
		} else if strings.HasPrefix(line, "     ") && !strings.HasPrefix(strings.TrimSpace(line), "/") {
			// Start a new feature
			if currentFeature.Key != "" {
				entry.Features = append(entry.Features, currentFeature)
			}
			currentFeature = GenBankFeature{
				Key:        strings.Fields(line)[0],
				Location:   strings.Join(strings.Fields(line)[1:], " "),
				Qualifiers: make(map[string]string),
			}
		}

		// SEQUENCE section
		if strings.HasPrefix(line, "ORIGIN") {
			isSequence = true
			entry.Sequence = ""
		}

		// Parse sequence data
		if isSequence && strings.HasPrefix(line, " ") && len(line) > 10 {
			sequence := strings.ReplaceAll(line[10:], " ", "")
			entry.Sequence += sequence
		}

		// End of entry
		if strings.HasPrefix(line, "//") {
			break
		}
	}

	// Append the last feature
	if currentFeature.Key != "" {
		entry.Features = append(entry.Features, currentFeature)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return entry, nil
}

// EqualGenBankEntry is a helper function to compare two GenBankEntry structs.
func EqualGenBankEntry(e1, e2 *GenBankEntry) bool {
	// Compare Locus
	if e1.Locus != e2.Locus {
		return false
	}

	// Compare Accession
	if len(e1.Accession) != len(e2.Accession) {
		return false
	}
	for i := range e1.Accession {
		if e1.Accession[i] != e2.Accession[i] {
			return false
		}
	}

	// Compare Definition
	if e1.Definition != e2.Definition {
		return false
	}

	// Compare Keywords
	if len(e1.Keywords) != len(e2.Keywords) {
		return false
	}
	for i := range e1.Keywords {
		if e1.Keywords[i] != e2.Keywords[i] {
			return false
		}
	}

	// Compare Features
	if len(e1.Features) != len(e2.Features) {
		return false
	}
	for i := range e1.Features {
		if e1.Features[i].Key != e2.Features[i].Key || e1.Features[i].Location != e2.Features[i].Location {
			return false
		}
		if len(e1.Features[i].Qualifiers) != len(e2.Features[i].Qualifiers) {
			return false
		}
		for k, v := range e1.Features[i].Qualifiers {
			if e2.Features[i].Qualifiers[k] != v {
				return false
			}
		}
	}

	// Compare Sequence
	if e1.Sequence != e2.Sequence {
		return false
	}

	return true
}

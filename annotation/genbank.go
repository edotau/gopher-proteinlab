package annotation

import (

	"fmt"
	"strings"

	"gopher-proteinlab/parseio"
)

// GenBankEntry represents a parsed GenBank file entry.
type GenBankEntry struct {
	Locus      string            // Locus line containing information about the sequence
	Definition string            // Definition line describing the sequence
	Accession  []string          // Accession numbers of the sequence
	Version    string            // Version information of the sequence
	Keywords   []string          // Keywords associated with the sequence
	Source     string            // Source organism or cell line for the sequence
	Organism   string            // Full organism classification of the sequence
	References []GenBankReference // List of references in the sequence
	Features   []GenBankFeature   // List of features such as genes and coding sequences
	Sequence   string            // The nucleotide or protein sequence
}

// GenBankFeature represents a feature in a GenBank file.
type GenBankFeature struct {
	Key        string            // Feature key (e.g., CDS, gene)
	Location   string            // Location of the feature in the sequence
	Qualifiers map[string]string // Additional qualifiers for the feature
}

// GenBankReference represents a reference section in a GenBank file.
type GenBankReference struct {
	Number  string // Reference number
	Authors string // Authors of the reference
	Title   string // Title of the referenced work
	Journal string // Journal of publication
	Medline string // Medline information
	Comment string // Additional comments about the reference
}

// GenBankReader reads a GenBank file, parses its content, and returns the parsed entry data.
// It processes the file in chunks to avoid loading the entire file into memory.
func parseGenBank(scanner *parseio.Scanalyzer) (*GenBankEntry, error) {
	entry := &GenBankEntry{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Use switch case to handle different line types in the GenBank file
		switch {
		case strings.HasPrefix(line, "LOCUS"):
			entry.Locus = strings.Fields(line[12:])[0]  // Capture only the first word
		case strings.HasPrefix(line, "DEFINITION"):
			entry.Definition = strings.TrimSpace(line[12:])
		case strings.HasPrefix(line, "ACCESSION"):
			entry.Accession = strings.Fields(line[12:])
		case strings.HasPrefix(line, "VERSION"):
			entry.Version = strings.TrimSpace(line[12:])
		case strings.HasPrefix(line, "KEYWORDS"):
			entry.Keywords = strings.Split(strings.TrimSpace(line[12:]), ";")
		case strings.HasPrefix(line, "SOURCE"):
			entry.Source = strings.TrimSpace(line[12:])
		case strings.HasPrefix(line, "ORGANISM"):
			entry.Organism = readMultiLineOrganism(scanner, line)
		case strings.HasPrefix(line, "REFERENCE"):
			entry.References = append(entry.References, readReference(scanner, line))
		case strings.HasPrefix(line, "FEATURES"):
			entry.Features = readFeatures(scanner)
		case strings.HasPrefix(line, "ORIGIN"):
			entry.Sequence = readSequence(scanner)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return entry, nil
}



// readMultiLineOrganism reads multi-line organism information from the GenBank file.
func readMultiLineOrganism(scanner *parseio.Scanalyzer, line string) string {
	// Start with the first line of the organism
	organism := strings.TrimSpace(line[12:])
	for scanner.Scan() {
		nextLine := strings.TrimSpace(scanner.Text())
		if len(nextLine) == 0 || strings.HasPrefix(nextLine, "REFERENCE") || strings.HasPrefix(nextLine, "FEATURES") {
			break
		}
		organism += " " + nextLine
	}
	return organism
}


// readReference reads a reference section from the GenBank file and returns a GenBankReference.
func readReference(scanner *parseio.Scanalyzer, firstLine string) GenBankReference {
	var reference GenBankReference
	reference.Number = strings.TrimSpace(firstLine[12:])

	// Continue reading the reference block
	for scanner.Scan() {
		refLine := strings.TrimSpace(scanner.Text())

		switch {
		case strings.HasPrefix(refLine, "AUTHORS"):
			reference.Authors = strings.TrimSpace(refLine[12:])
		case strings.HasPrefix(refLine, "TITLE"):
			reference.Title = strings.TrimSpace(refLine[12:])
		case strings.HasPrefix(refLine, "JOURNAL"):
			reference.Journal = strings.TrimSpace(refLine[12:])
		case strings.HasPrefix(refLine, "MEDLINE"):
			reference.Medline = strings.TrimSpace(refLine[12:])
		case strings.HasPrefix(refLine, "COMMENT"):
			reference.Comment = strings.TrimSpace(refLine[12:])
		default:
			// Stop reading once we reach a non-reference line
			return reference
		}
	}
	return reference
}

// readFeatures reads the features section from the GenBank file and returns a slice of GenBankFeatures.
func readFeatures(scanner *parseio.Scanalyzer) []GenBankFeature {
	var features []GenBankFeature
	var currentFeature GenBankFeature

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Stop processing features when ORIGIN (sequence) is reached
		if strings.HasPrefix(line, "ORIGIN") {
			break
		}

		// Handle feature lines (starting with feature key, like CDS)
		if strings.HasPrefix(line, "     ") && !strings.HasPrefix(line, "/") {
			// Append the current feature if it exists
			if currentFeature.Key != "" {
				features = append(features, currentFeature)
			}

			// Create a new feature
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				currentFeature = GenBankFeature{
					Key:        parts[0],
					Location:   parts[1],
					Qualifiers: make(map[string]string),
				}
			}
		} else if strings.HasPrefix(line, "/") { // Handle qualifier lines (e.g., /gene="sod")
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				qualifierKey := strings.TrimSpace(parts[0])
				qualifierValue := strings.Trim(strings.TrimSpace(parts[1]), "\"")
				currentFeature.Qualifiers[qualifierKey] = qualifierValue
			}
		}
	}

	// Append the last feature if present
	if currentFeature.Key != "" {
		features = append(features, currentFeature)
	}

	return features
}


// readSequence reads the sequence data from the GenBank file.
func readSequence(scanner *parseio.Scanalyzer) string {
	var sequenceBuilder strings.Builder
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "//" {
			break
		}
		// Remove line numbers and spaces from the sequence
		sequenceParts := strings.Fields(line)
		if len(sequenceParts) > 1 {
			sequenceBuilder.WriteString(strings.Join(sequenceParts[1:], ""))
		}
	}
	return sequenceBuilder.String()
}

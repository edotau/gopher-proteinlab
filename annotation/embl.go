package annotation

import (
	"fmt"
	"io"

	"gopher-proteinlab/parseio"
	"strings"
)

// EMBLEntry represents the structure of an EMBL file entry.
type EMBLEntry struct {
	ID        string    `json:"id"`
	Accession []string  `json:"accession"`
	Keywords  []string  `json:"keywords"`
	Source    string    `json:"source"`
	Features  []Feature `json:"features"`
	Sequence  string    `json:"sequence"`
}

// // Feature represents a feature in the EMBL file.
// type Feature struct {
// 	Key        string            `json:"key"`
// 	Location   string            `json:"location"`
// 	Qualifiers map[string]string `json:"qualifiers"`
// }

// EMBLReader reads an EMBL .dat file, decodes its content, and returns the parsed data one entry at a time.
func EMBLReader(filename string) {
	emblScanner := parseio.NewScanner(filename) // This will be a wrapper around bufio.Scanner
	defer emblScanner.Close()

	for {
		// Parse each entry and process it
		entry, err := parseEMBL(emblScanner)
		if err == io.EOF {
			break // End of file reached
		}
		parseio.ExitOnError(err)

		// Process the entry, e.g., print or store it
		fmt.Println(entry.ToString())
	}
}

// parseEMBL parses one EMBL entry at a time from the provided scanner.
func parseEMBL(scanner *parseio.Scanalyzer) (*EMBLEntry, error) {
	entry := &EMBLEntry{}
	currentFeature := Feature{}

	for scanner.Scan() {
		line := scanner.Text()

		// ID line (entry ID)
		if strings.HasPrefix(line, "ID") {
			entry.ID = strings.TrimSuffix(strings.TrimSpace(line[2:]), ";")
		}

		// AC line (accession numbers)
		if strings.HasPrefix(line, "AC") {
			accessions := strings.Split(strings.TrimSpace(line[2:]), ";")
			for _, i := range accessions {
				if acc := strings.TrimSpace(i); acc != "" {
					entry.Accession = append(entry.Accession, strings.TrimSpace(acc))
				}
			}
		}
		if strings.HasPrefix(line, "KW") {
			words := strings.Split(strings.TrimSpace(line[2:]), ";")
			for _, i := range words {
				if key := strings.TrimSpace(i); key != "" {
					entry.Keywords = append(entry.Keywords, key)
				}
			}
		}
		// FT line (features)
		if strings.HasPrefix(line, "FT") {
			ftLine := strings.TrimSpace(line[5:])
			if strings.HasPrefix(ftLine, "source") || strings.HasPrefix(ftLine, "CDS") {
				// Append the current feature if it exists
				if currentFeature.Key != "" {
					entry.Features = append(entry.Features, currentFeature)
				}
				// Create new feature
				currentFeature = Feature{
					Key:        strings.Fields(ftLine)[0],
					Location:   strings.Fields(ftLine)[1],
					Qualifiers: make(map[string]string),
				}
			} else {
				// Handle qualifier lines
				if strings.HasPrefix(ftLine, "/") {
					parts := strings.SplitN(ftLine, "=", 2)
					if len(parts) == 2 {
						qualifierKey := strings.TrimSpace(parts[0])
						qualifierValue := strings.Trim(strings.TrimSpace(parts[1]), "\"")
						currentFeature.Qualifiers[qualifierKey] = qualifierValue
					}
				}
			}
		}

		// SQ line (sequence)
		if strings.HasPrefix(line, "SQ") {
			var sequence strings.Builder
			for scanner.Scan() {
				seqLine := scanner.Text()
				if seqLine == "//" {
					break
				}
				sequence.WriteString(strings.TrimSpace(seqLine))
			}
			entry.Sequence = sequence.String()
		}
	}

	// Append the last feature if it exists
	if currentFeature.Key != "" {
		entry.Features = append(entry.Features, currentFeature)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return entry, nil
}

// EqualEmblEntry is a helper function to compare two EMBLEntry structs.
func EqualEmblEntry(e1, e2 *EMBLEntry) bool {
	// Compare IDs
	if e1.ID != e2.ID {
		return false
	}

	// Compare Accession numbers
	if len(e1.Accession) != len(e2.Accession) {
		return false
	}
	for i := range e1.Accession {
		if e1.Accession[i] != e2.Accession[i] {
			return false
		}
	}

	// Compare Keywords
	if len(e1.Keywords) != len(e2.Keywords) {
		fmt.Println("ERROR HERE")
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

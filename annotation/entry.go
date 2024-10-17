package annotation

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopher-proteinlab/parseio"
	"strings"
)

type Entry struct {
	Accession  []string    `xml:"accession"`
	Name       []string    `xml:"name"`
	Protein    ProteinInfo `xml:"protein"`
	Organism   Organism    `xml:"organism"`
	Sequence   Sequence    `xml:"sequence"`
	Created    string      `xml:"created,attr"`
	Modified   string      `xml:"modified,attr"`
	Version    int         `xml:"version,attr"`
	References []Reference `xml:"reference"`
}

type ProteinInfo struct {
	RecommendedName RecommendedName `xml:"recommendedName"`
}

type RecommendedName struct {
	FullName string `xml:"fullName"`
}

type Organism struct {
	Name []OrganismName `xml:"name"`
}

type OrganismName struct {
	Type string `xml:"type,attr"`
	Name string `xml:",chardata"`
}

type Sequence struct {
	Value    string `xml:",chardata"`
	Length   int    `xml:"length,attr"`
	Mass     int    `xml:"mass,attr"`
	Checksum string `xml:"checksum,attr"`
	Modified string `xml:"modified,attr"`
	Version  int    `xml:"version,attr"`
}

type Reference struct {
	Key      string   `xml:"key,attr"`
	Citation Citation `xml:"citation"`
}

type Citation struct {
	Title string `xml:"title"`
	Type  string `xml:"type,attr"`
}

// writeField is a helper function to write a field
func writeField(buffer *strings.Builder, label, value string) {
	parseio.HandleStrBuilder(buffer, label)
	parseio.HandleStrBuilder(buffer, value)
	parseio.ExitOnError(buffer.WriteByte('\n'))
}

func xmlToString(v interface{}) string {
	output, err := xml.MarshalIndent(v, "", "  ")
	parseio.ExitOnError(err)
	return string(output)
}

// ToString returns a string representation of the Entry struct.
func (e Entry) ToString() string {
	var words strings.Builder

	writeField(&words, "Accession: ", strings.Join(e.Accession, ", "))
	writeField(&words, "Name: ", strings.Join(e.Name, ", "))
	writeField(&words, "Protein: ", e.Protein.RecommendedName.FullName)

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

// ToJson is a Method to convert Entry to a string
func (e *Entry) ToJson() string {
	dictionary, err := json.MarshalIndent(e, "", "  ")
	parseio.ExitOnError(err)
	return string(dictionary)
}

func EqualEntries(e1, e2 *Entry) bool {
	if len(e1.Accession) != len(e2.Accession) {
		return false
	}
	for i := range e1.Accession {
		if e1.Accession[i] != e2.Accession[i] {
			return false
		}
	}

	if len(e1.Name) != len(e2.Name) {
		return false
	}
	for i := range e1.Name {
		if e1.Name[i] != e2.Name[i] {
			return false
		}
	}

	if e1.Protein.RecommendedName.FullName != e2.Protein.RecommendedName.FullName {
		return false
	}

	if len(e1.Organism.Name) != len(e2.Organism.Name) {
		return false
	}
	for i := range e1.Organism.Name {
		if e1.Organism.Name[i] != e2.Organism.Name[i] {
			return false
		}
	}

	if e1.Sequence != e2.Sequence {
		return false
	}

	if e1.Created != e2.Created {
		return false
	}

	if e1.Modified != e2.Modified {
		return false
	}

	if e1.Version != e2.Version {
		return false
	}

	if len(e1.References) != len(e2.References) {
		return false
	}

	for i := range e1.References {
		if e1.References[i] != e2.References[i] {
			return false
		}
	}

	return true
}

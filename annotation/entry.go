package annotation

import (
	"fmt"
	"gopher-proteinlab/simpleio"
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
	simpleio.HandleStrBuilder(buffer, label)
	simpleio.HandleStrBuilder(buffer, value)
	simpleio.CatchError(buffer.WriteByte('\n'))
}

// ToString returns a string representation of the Entry struct.
func (e Entry) ToString() string {
	var words strings.Builder

	writeField(&words, "Accession: ", strings.Join(e.Accession, ", "))
	writeField(&words, "Name: ", strings.Join(e.Name, ", "))
	writeField(&words, "Protein: ", e.Protein.RecommendedName.FullName)

	simpleio.HandleStrBuilder(&words, "Organism: ")
	for _, name := range e.Organism.Name {
		simpleio.HandleStrBuilder(&words, name.Type)
		simpleio.HandleStrBuilder(&words, ": ")
		simpleio.HandleStrBuilder(&words, name.Name)
		simpleio.HandleStrBuilder(&words, "; ")
	}
	simpleio.CatchError(words.WriteByte('\n'))

	writeField(&words, "Sequence: ", e.Sequence.Value)
	writeField(&words, "Created: ", e.Created)
	writeField(&words, "Modified: ", e.Modified)
	writeField(&words, "Version: ", fmt.Sprintf("%d", e.Version))

	simpleio.HandleStrBuilder(&words, "References: ")
	for _, ref := range e.References {
		simpleio.HandleStrBuilder(&words, ref.Citation.Title)
		simpleio.HandleStrBuilder(&words, "; ")
	}
	simpleio.CatchError(words.WriteByte('\n'))
	return words.String()
}

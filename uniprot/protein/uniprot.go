package annotation

import (
	"encoding/xml"
	"fmt"
	"io"

	"gopher-proteinlab/parseio"
	"gopher-proteinlab/protein"
)

// UniProtXML represents the structure of the root element for a UniProt XML file.
type UniProtXML struct {
	XMLName xml.Name       `xml:"uniprot"`
	Entries []UniProtEntry `xml:"entry"`
}

// UniProtEntry represents the data for each entry in the UniProt XML file.
type UniProtEntry struct {
	Accession  []string               `xml:"accession"`
	Name       []string               `xml:"name"`
	Protein    protein.UniProtProtein `xml:"protein"`
	Organism   protein.Organism       `xml:"organism"`
	Sequence   UniProtSequence        `xml:"sequence"`
	Created    string                 `xml:"created,attr"`
	Modified   string                 `xml:"modified,attr"`
	Version    int                    `xml:"version,attr"`
	References []UniProtReference     `xml:"reference"`
}

// UniProtOrganismName holds the organism name with its type (e.g., scientific or common).
type UniProtOrganismName struct {
	Type string `xml:"type,attr"`
	Name string `xml:",chardata"`
}

// UniProtSequence represents the sequence data for a UniProt entry.
type UniProtSequence struct {
	Value    string `xml:",chardata"`
	Length   int    `xml:"length,attr"`
	Mass     int    `xml:"mass,attr"`
	Checksum string `xml:"checksum,attr"`
	Modified string `xml:"modified,attr"`
	Version  int    `xml:"version,attr"`
}

// UniProtReference holds reference data for citations in a UniProt entry.
type UniProtReference struct {
	Key      string          `xml:"key,attr"`
	Citation UniProtCitation `xml:"citation"`
}

// UniProtCitation contains citation details for a reference.
type UniProtCitation struct {
	Title string `xml:"title"`
	Type  string `xml:"type,attr"`
}

// UniProtXMLReader reads a UniProt XML file, decodes its content, and returns the parsed data.
func UniProtXMLReader(filename string) {
	xmlReader := parseio.NewCodeReader(filename)
	defer xmlReader.Close()

	// Iterate through the tokens in the XML file
	decoder := xml.NewDecoder(xmlReader)
	for entry, err := parseUniProt(decoder); err != io.EOF; entry, err = parseUniProt(decoder) {
		parseio.ExitOnError(err)
		fmt.Println(entry)
	}
}

// parseUniProt parses the UniProt XML file for individual entries.
func parseUniProt(decoder *xml.Decoder) (*UniProtEntry, error) {
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "entry" {
				var entry UniProtEntry
				if err = decoder.DecodeElement(&entry, &se); err != nil {
					return nil, err
				}
				return &entry, nil
			}
		}
	}
}

// ToXmlString converts a UniProtEntry to an XML-formatted string.
func ToXmlString(e UniProtEntry) string {
	output, err := xml.MarshalIndent(e, "", "  ")
	parseio.ExitOnError(err)
	return string(output)
}

// EqualEntries compares two UniProtEntry structs to determine if they are equal.
func EqualEntries(e1, e2 *UniProtEntry) bool {
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

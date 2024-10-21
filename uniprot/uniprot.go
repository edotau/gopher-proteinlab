package uniprot

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopher-proteinlab/parseio"
	"io"
)

// Root element
type Uniprot struct {
	XMLName   xml.Name `xml:"uniprot"`
	Entries   []Entry  `xml:"entry"`
	Copyright string   `xml:"copyright,omitempty"`
}

// Entry definition
type Entry struct {
	Accession        []string         `xml:"accession"`
	Name             []string         `xml:"name"`
	Protein          ProteinEntry     `xml:"protein"`
	Gene             []Gene           `xml:"gene,omitempty"`
	Organism         Organism         `xml:"organism"`
	OrganismHost     []Organism       `xml:"organismHost,omitempty"`
	GeneLocation     []GeneLocation   `xml:"geneLocation,omitempty"`
	References       []Reference      `xml:"reference"`
	Comment          []Comment        `xml:"comment,omitempty"`
	DBReference      []DBReference    `xml:"dbReference,omitempty"`
	ProteinExistence ProteinExistence `xml:"proteinExistence"`
	Keyword          []Keyword        `xml:"keyword,omitempty"`
	Feature          []Feature        `xml:"feature,omitempty"`
	Evidence         []Evidence       `xml:"evidence,omitempty"`
	Sequence         Sequence         `xml:"sequence"`
	Dataset          string           `xml:"dataset,attr"`
	Created          string           `xml:"created,attr"`
	Modified         string           `xml:"modified,attr"`
	Version          int              `xml:"version,attr"`
}

// NameEntry represents a general name entry for both protein and gene names.
type NameEntry struct {
	Type     string `xml:"type,attr,omitempty"`     // Optional field for gene names
	Evidence []int  `xml:"evidence,attr,omitempty"` // Evidence is common in both cases
	Value    string `xml:",chardata"`               // The actual value of the name
}

// Gene represents a gene element.
type Gene struct {
	Name []NameEntry `xml:"name"`
}

// GeneLocationType definition
type GeneLocation struct {
	Type     string      `xml:"type,attr"`
	Name     []NameEntry `xml:"name,omitempty"`
	Evidence []int       `xml:"evidence,attr,omitempty"`
}

// Reference definition
type Reference struct {
	Citation Citation `xml:"citation"`
	Scope    []string `xml:"scope"`
	Source   *Source  `xml:"source,omitempty"`
	Evidence []int    `xml:"evidence,attr,omitempty"`
	Key      string   `xml:"key,attr"`
}

// Citation is used to represent strings with optional evidence attributes.
type Citation struct {
	Title       string        `xml:"title,omitempty"`
	AuthorList  []Person      `xml:"authorList>person,omitempty"`
	EditorList  []Person      `xml:"editorList>person,omitempty"`
	Locator     string        `xml:"locator,omitempty"`
	DBReference []DBReference `xml:"dbReference,omitempty"`
	Type        string        `xml:"type,attr"`
	Date        string        `xml:"date,attr,omitempty"`
	Name        string        `xml:"name,attr,omitempty"`
	Volume      string        `xml:"volume,attr,omitempty"`
	First       string        `xml:"first,attr,omitempty"`
	Last        string        `xml:"last,attr,omitempty"`
}

// CommentType definition
type Comment struct {
	Type string      `xml:"type,attr"`
	Text []NameEntry `xml:"text,omitempty"`
}

// PropertyType definition
type Property struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

// KeywordType definition
type Keyword struct {
	ID       string `xml:"id,attr"`
	Evidence []int  `xml:"evidence,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// FeatureType definition
type Feature struct {
	Type        string   `xml:"type,attr"`
	ID          string   `xml:"id,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Evidence    []int    `xml:"evidence,attr,omitempty"`
	Location    Location `xml:"location"`
}

// LocationType definition
type Location struct {
	Begin    *Position `xml:"begin,omitempty"`
	End      *Position `xml:"end,omitempty"`
	Position *Position `xml:"position,omitempty"`
}

// PositionType definition
type Position struct {
	Position uint64 `xml:"position,attr,omitempty"`
	Status   string `xml:"status,attr,omitempty"`
}

// SequenceType definition
type Sequence struct {
	Length   int    `xml:"length,attr"`
	Mass     int    `xml:"mass,attr"`
	Checksum string `xml:"checksum,attr"`
	Modified string `xml:"modified,attr"`
	Version  int    `xml:"version,attr"`
	Value    string `xml:",chardata"`
}

// EvidenceType definition
type Evidence struct {
	Type         string       `xml:"type,attr"`
	Key          int          `xml:"key,attr"`
	Source       *Source      `xml:"source,omitempty"`
	ImportedFrom *DBReference `xml:"importedFrom,omitempty"`
}

// SourceType definition
type Source struct {
	DBReference *DBReference `xml:"dbReference,omitempty"`
	Ref         int          `xml:"ref,attr,omitempty"`
}

// PersonType definition
type Person struct {
	Name string `xml:"name,attr"`
}

// MoleculeType definition
type Molecule struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
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
func parseUniProt(decoder *xml.Decoder) (*Entry, error) {
	for {
		if tok, err := decoder.Token(); parseio.ExitOnError(err) {
			switch se := tok.(type) {
			case xml.StartElement:
				if se.Name.Local == "entry" {
					var entry Entry
					if err = decoder.DecodeElement(&entry, &se); err != nil {
						return nil, err
					}
					return &entry, nil
				}
			}
		}
	}
}

// ToXmlString converts a UniProtEntry to an XML-formatted string.
func ToString(e Entry) string {
	output, err := xml.MarshalIndent(e, "", "  ")
	parseio.ExitOnError(err)
	return string(output)
}

func (e *Entry) ToJson() string {
	txt := parseio.NewTxtBuilder()
	data, err := json.MarshalIndent(e, "", "  ")
	if parseio.ExitOnError(err) {
		txt.Write(data)
	}
	return txt.String()
}

func XmlString(e Entry) string {
	output, err := xml.MarshalIndent(e, "", "  ")
	parseio.ExitOnError(err)
	return string(output)
}
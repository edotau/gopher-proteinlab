package annotation

import (
	"encoding/xml"
	"fmt"
	"gopher-proteinlab/simpleio"
	"io"
)

type UniProtXML struct {
	XMLName xml.Name `xml:"uniprot"`
	Entries []Entry  `xml:"entry"`
}

// // UniProtXMLReader reads a UniProt XML file, decodes the XML content, and returns the parsed data.
// func UniProtXMLReader(filename string) {
// 	xmlFile := simpleio.SimpleOpen(filename)
// 	defer xmlFile.Close()

// 	decoder := xml.NewDecoder(xmlFile)

// 	// Iterate through the tokens in the XML file
// 	for entry, err := parseUniProt(decoder); err != io.EOF; entry, err = parseUniProt(decoder) {
// 		entry.ToString()
// 	}
// }

// UniProtXMLReader reads a UniProt XML file, decodes the XML content, and returns the parsed data.
func UniProtXMLReader(filename string) error {
	xmlReader, xmlFile := simpleio.FileHandler(filename)
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlReader)

	// Iterate through the tokens in the XML file
	for {
		entry, err := parseUniProt(decoder)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error parsing UniProt XML: %v", err)
		}
		fmt.Println(entry.ToString())
	}

	return nil
}

func parseUniProt(decoder *xml.Decoder) (*Entry, error) {
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
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

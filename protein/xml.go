package protein

import (
	"encoding/xml"
)

// UniProtProtein contains the protein information for a UniProt entry, including gene information.
type UniProtProtein struct {
	XMLName          xml.Name                 `xml:"protein"`
	RecommendedName  UniProtRecommendedName   `xml:"recommendedName"`
	AlternativeNames []UniProtAlternativeName `xml:"alternativeName"`
	SubmittedNames   []UniProtSubmittedName   `xml:"submittedName"`
	AllergenName     string                   `xml:"allergenName,omitempty"`
	BiotechName      string                   `xml:"biotechName,omitempty"`
	CdAntigenNames   []string                 `xml:"cdAntigenName,omitempty"`
	InnNames         []string                 `xml:"innName,omitempty"`
	Genes            []Gene                   `xml:"gene"`
}

// UniProtRecommendedName holds the recommended full name for a protein.
type UniProtRecommendedName struct {
	FullName  EvidencedString   `xml:"fullName"`
	ShortName []EvidencedString `xml:"shortName,omitempty"`
	EcNumber  []EvidencedString `xml:"ecNumber,omitempty"`
}

// UniProtAlternativeName holds the alternative names for a protein.
type UniProtAlternativeName struct {
	FullName  EvidencedString   `xml:"fullName,omitempty"`
	ShortName []EvidencedString `xml:"shortName,omitempty"`
	EcNumber  []EvidencedString `xml:"ecNumber,omitempty"`
}

// UniProtSubmittedName holds the submitted names for a protein.
type UniProtSubmittedName struct {
	FullName EvidencedString   `xml:"fullName"`
	EcNumber []EvidencedString `xml:"ecNumber,omitempty"`
}

// Gene represents a gene element.
type Gene struct {
	Names []GeneName `xml:"name"`
}

// GeneName represents a name for a gene with attributes type and evidence.
type GeneName struct {
	Value    string `xml:",chardata"`
	Type     string `xml:"type,attr"`
	Evidence string `xml:"evidence,attr,omitempty"`
}

// EvidencedString handles strings with optional evidence attributes.
type EvidencedString struct {
	Value string `xml:",chardata"`
}

// Organism represents the organismType complexType in the XML schema.
type Organism struct {
	XMLName     xml.Name       `xml:"organism"`
	Names       []OrganismName `xml:"name"`          // Multiple names for the organism
	DbReference []DbReference  `xml:"dbReference"`   // Cross-references to databases (e.g., NCBI)
	Lineage     *Lineage       `xml:"lineage"`       // Optional lineage of the organism
	Evidence    string         `xml:"evidence,attr"` // Optional evidence attribute
}

// OrganismName represents the organismNameType complexType.
type OrganismName struct {
	Value string `xml:",chardata"` // The name of the organism
	Type  string `xml:"type,attr"` // The type attribute (e.g., scientific, common, synonym, etc.)
}

// DbReference represents the dbReferenceType for cross-referencing to databases.
type DbReference struct {
	Type string `xml:"type,attr"` // Type of database (e.g., NCBI taxonomy)
	ID   string `xml:"id,attr"`   // ID in the cross-referenced database
}

// Lineage represents the lineage of the organism.
type Lineage struct {
	Taxa []string `xml:"taxon"` // A list of taxonomic ranks in the lineage
}

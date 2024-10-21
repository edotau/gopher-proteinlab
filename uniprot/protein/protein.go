package protein

// Citation is used to represent strings with optional evidence attributes.
type Citation struct {
	Value    string `xml:",chardata"`
	Evidence string `xml:"evidence,attr,omitempty"`
}

// References is an alias type for a slice []Citation.
type References []Citation

// NameEntry holds the full name, short names, and EC numbers for a protein.
type NameEntry struct {
	FullName  Citation   `xml:"fullName"`            // Full name of the protein
	ShortName References `xml:"shortName,omitempty"` // Optional short names
	EcNumber  References `xml:"ecNumber,omitempty"`  // Optional EC numbers
}

// NameEntryList is an alias type for a slice []NameEntry.
type NameEntryList []NameEntry

// Protein contains the protein information for a UniProt entry, including gene information.
type Protein struct {
	Name            NameEntry   `xml:"recommendedName,omitempty"`
	AlternativeName NameEntryList `xml:"alternativeName,omitempty"`
	SubmittedName   NameEntryList `xml:"submittedName,omitempty"`
	AllergenName     string        `xml:"allergenName,omitempty"`      // Allergen name, if applicable
	BiotechName      string        `xml:"biotechName,omitempty"`       // Biotechnological name, if applicable
	CdAntigenNames   []string      `xml:"cdAntigenName,omitempty"`     // CD antigen names, if applicable
	InnNames         []string      `xml:"innName,omitempty"`           // International Nonproprietary Names (INN), if applicable
	Domain          NameEntryList `xml:"domain,omitempty"`
	Component       NameEntryList `xml:"component,omitempty"`
}


// GeneName represents a name for a gene with attributes type and evidence.
type GeneName struct {
	Value    string `xml:",chardata"`
	Type     string `xml:"type,attr"`
	Evidence string `xml:"evidence,attr,omitempty"`
}

// Gene represents a gene element.
type Gene struct {
	Names []GeneName `xml:"name"`
}

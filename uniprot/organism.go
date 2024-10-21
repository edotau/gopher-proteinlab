package uniprot

// OrganismType definition
type Organism struct {
	Name        []NameEntry   `xml:"name"`
	DBReference []DBReference `xml:"dbReference"`
	Lineage     *Lineage      `xml:"lineage,omitempty"`
	Evidence    []int         `xml:"evidence,attr,omitempty"`
}

// DBReferenceType definition
type DBReference struct {
	Type     string     `xml:"type,attr"`
	ID       string     `xml:"id,attr"`
	Molecule *Molecule  `xml:"molecule,omitempty"`
	Property []Property `xml:"property,omitempty"`
	Evidence []int      `xml:"evidence,attr,omitempty"`
}

// Lineage definition
type Lineage struct {
	Taxon []string `xml:"taxon"`
}

func (alpha Organism) Equal(beta Organism) bool {
	// Compare Evidence slices
	if len(alpha.Evidence) != len(beta.Evidence) {
		return false
	}
	for i := range alpha.Evidence {
		if alpha.Evidence[i] != beta.Evidence[i] {
			return false
		}
	}

	// Compare Name slices
	if len(alpha.Name) != len(beta.Name) {
		return false
	}

	for i, v := range alpha.Name {
		if !v.Equal(beta.Name[i]) {
			return false
		}
	}

	// Compare DBReference slices
	if len(alpha.DBReference) != len(beta.DBReference) {
		return false
	}
	for i := range alpha.DBReference {
		if !alpha.DBReference[i].Equal(beta.DBReference[i]) {
			return false
		}
	}

	// Compare Lineage (optional, so need to handle nil cases)
	if (alpha.Lineage == nil) != (beta.Lineage == nil) {
		return false
	}
	if alpha.Lineage != nil && !alpha.Lineage.Equal(*beta.Lineage) {
		return false
	}

	return true
}

// Equal method for DBReference
func (alpha DBReference) Equal(beta DBReference) bool {
	if alpha.Type != beta.Type || alpha.ID != beta.ID {
		return false
	}

	// Compare Molecule (optional)
	if (alpha.Molecule == nil) != (beta.Molecule == nil) {
		return false
	}
	if alpha.Molecule != nil && *alpha.Molecule != *beta.Molecule {
		return false
	}

	// Compare Property slices
	if len(alpha.Property) != len(beta.Property) {
		return false
	}
	for i := range alpha.Property {
		if alpha.Property[i] != beta.Property[i] {
			return false
		}
	}

	// Compare Evidence slices
	if len(alpha.Evidence) != len(beta.Evidence) {
		return false
	}
	for i := range alpha.Evidence {
		if alpha.Evidence[i] != beta.Evidence[i] {
			return false
		}
	}

	return true
}

// Equal method for Lineage
func (alpha Lineage) Equal(beta Lineage) bool {
	if len(alpha.Taxon) != len(beta.Taxon) {
		return false
	}
	for i := range alpha.Taxon {
		if alpha.Taxon[i] != beta.Taxon[i] {
			return false
		}
	}
	return true
}

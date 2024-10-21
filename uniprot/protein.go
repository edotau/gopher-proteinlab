package uniprot

import (
	"encoding/json"
	"gopher-proteinlab/parseio"
)

func (e *ProteinEntry) ToJson() string {
	txt := parseio.NewTxtBuilder()
	data, err := json.MarshalIndent(e, "", "  ")
	if parseio.ExitOnError(err) {
		txt.Write(data)
	}
	return txt.String()
}

func (alpha NameEntry) Equal(beta NameEntry) bool {
	if alpha.Value != beta.Value {
		return false
	}
	if alpha.Type != beta.Type {
		return false
	}
	if alpha.Evidence != beta.Evidence {
		return false
	}
	return true
}

func (alpha ProteinName) Equal(beta ProteinName) bool {
	if !alpha.FullName.Equal(beta.FullName) {
		return false
	}

	if len(alpha.ShortName) != len(beta.ShortName) {
		return false
	}
	for i := range alpha.ShortName {
		if !alpha.ShortName[i].Equal(beta.ShortName[i]) {
			return false
		}
	}

	if len(alpha.ECNumber) != len(beta.ECNumber) {
		return false
	}
	for i := range alpha.ECNumber {
		if !alpha.ECNumber[i].Equal(beta.ECNumber[i]) {
			return false
		}
	}

	return true
}

func (alpha ProteinEntry) Equal(beta ProteinEntry) bool {
	// Compare RecommendedName, which is a pointer
	if !alpha.RecommendedName.Equal(beta.RecommendedName) {
		return false
	}

	for i := range alpha.AlternativeName {
		if !alpha.AlternativeName[i].Equal(beta.AlternativeName[i]) {
			return false
		}
	}

	// Compare SubmittedName slices
	if !alpha.SubmittedName.Equal(beta.SubmittedName) {
		return false
	}

	// Compare Domain slices
	if len(alpha.Domain) != len(beta.Domain) {
		return false
	}
	for i := range alpha.Domain {
		if !alpha.Domain[i].Equal(beta.Domain[i]) {
			return false
		}
	}

	// Compare Component slices
	if len(alpha.Component) != len(beta.Component) {
		return false
	}
	for i := range alpha.Component {
		if !alpha.Component[i].Equal(beta.Component[i]) {
			return false
		}
	}

	// Compare AllergenName (which is a pointer)
	if (alpha.AllergenName == nil) != (beta.AllergenName == nil) {
		return false
	}
	if alpha.AllergenName != nil && !alpha.AllergenName.Equal(*beta.AllergenName) {
		return false
	}

	// Compare BiotechName (which is a pointer)
	if (alpha.BiotechName == nil) != (beta.BiotechName == nil) {
		return false
	}
	if alpha.BiotechName != nil && !alpha.BiotechName.Equal(*beta.BiotechName) {
		return false
	}

	// Compare CDAntigenNames slices
	if len(alpha.CDAntigenNames) != len(beta.CDAntigenNames) {
		return false
	}
	for i := range alpha.CDAntigenNames {
		if !alpha.CDAntigenNames[i].Equal(beta.CDAntigenNames[i]) {
			return false
		}
	}

	// Compare INNNames slices
	if len(alpha.InnNames) != len(beta.InnNames) {
		return false
	}
	for i := range alpha.InnNames {
		if !alpha.InnNames[i].Equal(beta.InnNames[i]) {
			return false
		}
	}

	// All fields are equal
	return true
}

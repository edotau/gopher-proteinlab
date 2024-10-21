package uniprot

// // Equal compares two Citation structs for equality.
// func (c Citation) Equal(other Citation) bool {
// 	// Compare the Title field
// 	if c.Title != other.Title {
// 		return false
// 	}

// 	// Compare the Type field
// 	if c.Type != other.Type {
// 		return false
// 	}

// 	// Compare the AuthorList field (length and contents)
// 	if len(c.AuthorList) != len(other.AuthorList) {
// 		return false
// 	}
// 	for i := range c.AuthorList {
// 		if c.AuthorList[i] != other.AuthorList[i] {
// 			return false
// 		}
// 	}

// 	// Compare the DbReference field (length and contents)
// 	if len(c.DbReference) != len(other.DbReference) {
// 		return false
// 	}
// 	for i := range c.DbReference {
// 		if c.DbReference[i] != other.DbReference[i] {
// 			return false
// 		}
// 	}

// 	// If all fields match, return true
// 	return true
// }

// // Equal compares two NameEntry instances for equality.
// func (alpha Protein) Equal(beta Protein) bool {
// 	if !alpha.Entry.FullName.Equal(beta.Entry.FullName) {
// 		return false
// 	}
// 	if len(alpha.Entry.ShortNames) != len(beta.Entry.ShortNames){
// 		return false
// 	} else {
// 		for i := range alpha.Entry.ShortNames {
// 			if !alpha.Entry.ShortNames[i].Equal(beta.Entry.ShortNames[i]) {
// 				return false
// 			}
// 		}
// 	}

// 	if  len(alpha.Entry.ECNumbers) != len(beta.Entry.ECNumbers) {
// 		return false
// 	} else {
// 		for i := range alpha.Entry.ECNumbers {
// 			if !alpha.Entry.ECNumbers[i].Equal(beta.Entry.ECNumbers[i]) {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }
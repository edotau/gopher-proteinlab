package protein

import (
	"fmt"
	"strings"

	"gopher-proteinlab/stdio"
)

// Protein amino acid byte
type Protein byte

// Standard protein constant definitions
const (
	Ala     Protein = 'A' // Alanine (A)
	Arg     Protein = 'R' // Arginine (R)
	Asn     Protein = 'N' // Asparagine (N)
	Asp     Protein = 'D' // Aspartic acid (D)
	Cys     Protein = 'C' // Cysteine (C)
	Gln     Protein = 'Q' // Glutamine (Q)
	Glu     Protein = 'E' // Glutamic acid (E)
	Gly     Protein = 'G' // Glycine (G)
	His     Protein = 'H' // Histidine (H)
	Ile     Protein = 'I' // Isoleucine (I)
	Leu     Protein = 'L' // Leucine (L)
	Lys     Protein = 'K' // Lysine (K)
	Met     Protein = 'M' // Methionine (M)
	Phe     Protein = 'F' // Phenylalanine (F)
	Pro     Protein = 'P' // Proline (P)
	Pyl     Protein = 'O' // Pyrrolysine (O)
	Ser     Protein = 'S' // Serine (S)
	Sec     Protein = 'U' // Selenocysteine (U)
	Thr     Protein = 'T' // Threonine (T)
	Trp     Protein = 'W' // Tryptophan (W)
	Tyr     Protein = 'Y' // Tyrosine (Y)
	Val     Protein = 'V' // Valine (V)
	Asx     Protein = 'B' // Aspartic acid or Asparagine (B)
	Glx     Protein = 'Z' // Glutamine or Glutamic acid (Z)
	Xaa     Protein = 'X' // Any amino acid (X)
	Xle     Protein = 'J' // Leucine or Isoleucine (J)
	Stop    Protein = '*' // Stop/Termination codon (*)
	Unknown Protein = 255 // Unknown amino acid
)

// Unusual amino acids from the next iota value after standard amino acids
const (
	Aad   Protein = iota + 100 // 2-Aminoadipic acid
	bAad                       // 3-Aminoadipic acid
	bAla                       // beta-Alanine (bAla)
	Abu                        // 2-Aminobutyric acid
	Abu4                       // 4-Aminobutyric acid
	Acp                        // 6-Aminocaproic acid
	Ahe                        // 2-Aminoheptanoic acid
	Aib                        // 2-Aminoisobutyric acid
	bAib                       // 3-Aminoisobutyric acid
	Apm                        // 2-Aminopimelic acid
	Dbu                        // 2,4-Diaminobutyric acid
	Des                        // Desmosine
	Dpm                        // 2,2'-Diaminopimelic acid
	Dpr                        // 2,3-Diaminoproprionic acid
	EtGly                      // N-Ethylglycine
	EtAsn                      // N-Ethylasparagine
	Hyl                        // Hydroxylysine
	aHyl                       // allo-Hydroxylysine
	Hyp3                       // 3-Hydroxyproline
	Hyp4                       // 4-Hydroxyproline
	Ide                        // Isodesmosine
	aIle                       // allo-Isoleucine
	MeGly                      // N-Methylglycine
	MeIle                      // N-Methylisoleucine
	MeLys                      // 6-N-Methyllysine
	MeVal                      // N-Methylvaline
	Nva                        // Norvaline
	Nle                        // Norleucine
	Orn                        // Ornithine
)

// AminoAcidMap maps bytes to Protein types
var AminoAcidMap = map[byte]Protein{
	'A': Ala, 'R': Arg, 'N': Asn, 'D': Asp, 'C': Cys, 'Q': Gln, 'E': Glu,
	'G': Gly, 'H': His, 'I': Ile, 'L': Leu, 'K': Lys, 'M': Met, 'F': Phe,
	'P': Pro, 'O': Pyl, 'S': Ser, 'U': Sec, 'T': Thr, 'W': Trp, 'Y': Tyr,
	'V': Val, 'B': Asx, 'Z': Glx, 'X': Xaa, 'J': Xle, '*': Stop,
}

// ProteinToByteMap maps Protein types back to bytes
var ProteinToByteMap = make(map[Protein]byte)

func init() {
	for k, v := range AminoAcidMap {
		ProteinToByteMap[v] = k
	}
}

// ByteToAminoAcid converts a byte representing a standard amino acid to its corresponding Base value.
func ByteToAminoAcid(b byte) (Protein, error) {
	if base, exists := AminoAcidMap[b]; exists {
		return base, nil
	}
	return Unknown, fmt.Errorf("error: '%s' is an invalid amino acid symbol. Ensure the input contains valid characters.", string(b))
}

// ToProteins converts string to a slice of Protein amino acids.
func ToProteins(text string) []Protein {
	var proteins []Protein = make([]Protein, len(text))

	for i := 0; i < len(text); i++ {
		if aa, err := ByteToAminoAcid(text[i]); simpleio.CatchError(err) {
			proteins[i] = aa
		}
	}
	return proteins
}

// ToString converts a slice of Protein to a string.
func ToString(proteins []Protein) string {
	var words strings.Builder
	for _, aa := range proteins {
		simpleio.CatchError(words.WriteByte(ProteinToByteMap[aa]))
	}
	return words.String()
}

// Equal asserts if two slices of Protein amino acids are equal.
func Equal(a, b []Protein) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

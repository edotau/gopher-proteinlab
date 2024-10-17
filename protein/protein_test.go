package protein

import (
	"testing"
)

var testcases = []struct {
	txt      string
	expected []Protein
}{
	{"ARNDCQEGHILKMFPSTWYVBZX", []Protein{
		Ala, Arg, Asn, Asp, Cys, Gln, Glu, Gly, His, Ile, Leu, Lys, Met, Phe, Pro, Ser, Thr, Trp, Tyr, Val, Asx, Glx, Xaa,
	}},
	{"EIVMTQSPSTLSASVGDRVIITCQASEIIHSWLAWYQQKPGKAPKLL", []Protein{
		Glu, Ile, Val, Met, Thr, Gln, Ser, Pro, Ser, Thr, Leu, Ser, Ala, Ser, Val, Gly, Asp, Arg, Val, Ile, Ile, Thr, Cys, Gln, Ala, Ser, Glu, Ile, Ile, His, Ser, Trp, Leu, Ala, Trp, Tyr, Gln, Gln, Lys, Pro, Gly, Lys, Ala, Pro, Lys, Leu, Leu,
	}},
}

func TestToAminoAcids(t *testing.T) {
	for _, test := range testcases {
		for i := 0; i < len(test.txt); i++ {
			aminoAcid, err := ByteToAminoAcid(test.txt[i])
			if err != nil {
				t.Errorf("Error converting %c: %v", test.txt[i], err)
			}
			if aminoAcid != test.expected[i] {
				t.Errorf("Error: ByteToAminoAcid(%c), expected: %c.\n", test.txt[i], ProteinToByteMap[test.expected[i]])
			}
		}
	}
}

func TestToProteinSequence(t *testing.T) {
	for _, test := range testcases {
		if !Equal(ToProteins(test.txt), test.expected) {
			t.Errorf("Error: ToProteinSeq(%s), expected: %s.\n", test.txt, ToString(test.expected))
		}
	}
}

func TestProteinsToString(t *testing.T) {
	for _, test := range testcases {
		if ToString(test.expected) != test.txt {
			t.Errorf("Error: ToString(%s), expected: %s.\n", test.expected, test.txt)
		}
	}
}

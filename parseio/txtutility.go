package parseio

import "strings"

type TxtUtility struct {
	strings.Builder
}

// NewTxtBuilder wraps around strings.Builder for the use of custom methods.
func NewTxtBuilder() *TxtUtility {
	txt := strings.Builder{}

	return &TxtUtility{
		Builder: txt,
	}
}

// Write writes the provided slice of [][byte to the TxtUtility instance, panicing on error.
func (txt *TxtUtility) Write(b []byte) {
	if _, err := txt.Builder.Write(b); err != nil {
		panic(err)
	}
}

// WriteString writes the provided string to the TxtUtility instance, panicing on error.
func (txt *TxtUtility) WriteString(s string) {
	if _, err := txt.Builder.WriteString(s); err != nil {
		panic(err)
	}
}

// WriteByte writes a single byte to the TxtUtility instance.
func (txt *TxtUtility) WriteByte(b byte) {
	ExitOnError(txt.Builder.WriteByte(b))
}

func (txt *TxtUtility) WriteTag(label, value string) {
	txt.WriteString(label)
	txt.WriteByte(':')
	txt.WriteByte(' ')
	txt.WriteString(value)
}

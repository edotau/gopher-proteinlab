package parseio

import (
	"testing"
)

func TestWrite(t *testing.T) {
	txt := NewTxtBuilder()
	expected := "hello"
	txt.Write([]byte(expected))

	if txt.String() != expected {
		t.Errorf("Expected: '%s', test case: '%s'", expected, txt.String())
	}
}

func TestWriteString(t *testing.T) {
	txt := NewTxtBuilder()
	expected := "Test string"
	txt.WriteString(expected)

	if txt.String() != expected {
		t.Errorf("Expected: '%s', test case: '%s'", expected, txt.String())
	}
}

func TestWriteByte(t *testing.T) {
	txt := NewTxtBuilder()
	expected := "a"
	txt.WriteByte('a')

	if txt.String() != expected {
		t.Errorf("Expected: '%s', test case: '%s'", expected, txt.String())
	}
}

func TestWriteTag(t *testing.T) {
	txt := NewTxtBuilder()
	label := "Key"
	value := "Value"
	expected := "Key: Value"
	txt.WriteTag(label, value)

	if txt.String() != expected {
		t.Errorf("Expected: '%s', test case: '%s'", expected, txt.String())
	}
}

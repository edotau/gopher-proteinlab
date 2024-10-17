package simpleio

import (
	"log"
	"strings"
)

// CatchError will panic if input error is not nil.
func CatchError(err error) bool {
	if err != nil {
		log.Panic(err)
	}
	return err == nil
}

// ThrowError will output a warning message and return false if err != nil - Warning: err msg
func ThrowError(err error) bool {
	if err != nil {
		log.Printf("Warning: %s", err)
	}
	return err == nil
}

func HandleStrBuilder(buffer *strings.Builder, text string) {
	_, err := buffer.WriteString(text)
	CatchError(err)
}

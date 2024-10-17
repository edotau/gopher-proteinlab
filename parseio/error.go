package parseio

import (
	"log"
	"strings"
)

// ExitOnError will panic if input error is not nil.
func ExitOnError(err error) bool {
	if err != nil {
		log.Panic(err)
	}
	return err == nil
}

// WarningError will output a warning message and return false if err != nil - Warning: err msg
func WarningError(err error) bool {
	if err != nil {
		log.Printf("Warning: %s", err)
	}
	return err == nil
}

func HandleStrBuilder(buffer *strings.Builder, text string) {
	_, err := buffer.WriteString(text)
	ExitOnError(err)
}

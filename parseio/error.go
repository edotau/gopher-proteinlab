package parseio

import (
	"errors"
	"fmt"
	"io"
	"log"
	"runtime/debug"
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

// CatchEofErr will treat the io.EOF as nil and panic on all other errors.
func CatchEofErr(err error) bool {
	if errors.Is(err, io.EOF) {
		return false
	}
	if err != nil {
		log.Panic(err)
	}
	return err == nil
}

// RecoverError recovers from a panic and returns the error message as a string.
func RecoverError(err error) string {
	var errMsg string
	defer func() {
		if r := recover(); r != nil {
			strMsg, ok := r.(string)
			if !ok {
				strMsg = fmt.Sprintf("Panic recovered: %v\n%s", r, debug.Stack())
			}
			errMsg = strMsg
		}
	}()
	if err != nil {
		panic(err)
	}
	return errMsg
}

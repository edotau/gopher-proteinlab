package parseio

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestCatchError(t *testing.T) {
	err := ExitOnError(nil)
	if !err {
		t.Errorf("Expected true for nil error, got false")
	}
	testErr := fmt.Errorf("test error")
	expectedMessage := "test error"

	defer func() {
		recover()
	}()

	ExitOnError(testErr)
	errMsg := RecoverError(testErr)
	if errMsg != expectedMessage {
		t.Errorf("Expected panic message %q, got %q", expectedMessage, errMsg)
	}
}

func TestWarningError(t *testing.T) {
	var logOutput strings.Builder // Capture log output
	log.SetOutput(&logOutput)
	err := WarningError(nil) // Test case where error is nil

	if !err {
		t.Errorf("Expected true for nil error, got false")
	}
	if logOutput.Len() != 0 {
		t.Errorf("Expected no log output for nil error, got %s", logOutput.String())
	}

	logOutput.Reset()
	testErr := fmt.Errorf("test error")
	err = WarningError(testErr)

	if err {
		t.Errorf("Expected false for non-nil error, got true")
	}
	expectedLog := "Warning: test error"

	if !strings.Contains(logOutput.String(), expectedLog) {
		t.Errorf("Expected log output to contain %q, got %s", expectedLog, logOutput.String())
	}
}

func TestRecoverError(t *testing.T) {
	errMsg := RecoverError(nil)
	if errMsg != "" {
		t.Errorf("Expected empty error message, got %q", errMsg)
	}
}

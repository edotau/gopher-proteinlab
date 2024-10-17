package stdio

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestCatchError(t *testing.T) {
	err := CatchError(nil)
	if !err {
		t.Errorf("Expected true for nil error, got false")
	}
	defer func() {
		if r := recover(); r != nil {
			expectedMessage := "test error"
			if r != expectedMessage {
				t.Errorf("Expected panic message %q, got %q", expectedMessage, r)
			}
		} else {
			t.Errorf("Expected panic but did not get one")
		}
	}()

	testErr := fmt.Errorf("test error")
	CatchError(testErr)
	t.Errorf("CatchError should have panicked but did not")
}

func TestThrowError(t *testing.T) {
	// Capture log output
	var logOutput strings.Builder
	log.SetOutput(&logOutput)

	// Test case where error is nil
	err := ThrowError(nil)
	if !err {
		t.Errorf("Expected true for nil error, got false")
	}
	if logOutput.Len() != 0 {
		t.Errorf("Expected no log output for nil error, got %s", logOutput.String())
	}

	logOutput.Reset()

	testErr := fmt.Errorf("test error")
	err = ThrowError(testErr)

	if err {
		t.Errorf("Expected false for non-nil error, got true")
	}
	expectedLog := "Warning: test error"

	if !strings.Contains(logOutput.String(), expectedLog) {
		t.Errorf("Expected log output to contain %q, got %s", expectedLog, logOutput.String())
	}
}

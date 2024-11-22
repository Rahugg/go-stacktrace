package errorhandler

import (
	"errors"
	"strings"
	"testing"
)

func TestWrapError(t *testing.T) {
	err := errors.New("base error")
	wrappedErr := WrapError(err, "payload", "additional message")
	
	if wrappedErr == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
	
	var te *TracedError
	ok := errors.As(wrappedErr, &te)
	if !ok {
		t.Fatalf("expected *TracedError, got %T", wrappedErr)
	}
	
	if te.Payload != "payload" {
		t.Errorf("expected payload to be 'payload', got %s", te.Payload)
	}
	
	if te.UserMessage != "additional message" {
		t.Errorf("expected user message to be 'additional message', got %s", te.UserMessage)
	}
	
	if te.OriginalError.Error() != "base error" {
		t.Errorf("expected original error message to be 'base error', got %s", te.OriginalError.Error())
	}
}

func TestFailOnError(t *testing.T) {
	// Test simple error
	err := errors.New("simple error")
	result := FailOnError(err)
	if result != "simple error" {
		t.Errorf("unexpected result: %s", result)
	}
	
	// Test nil error
	result = FailOnError(nil)
	if result != "" {
		t.Errorf("expected empty result for nil error, got %s", result)
	}
	
	// Test TracedError
	tracedErr := WrapError(err, "test payload", "test message")
	result = FailOnError(tracedErr)
	if !strings.Contains(result, "test message") {
		t.Errorf("expected result to contain 'test message', got %s", result)
	}
}

func TestEnableColors(t *testing.T) {
	// Save original state
	originalState := EnableColors
	
	// Test with colors enabled
	SetEnableColors(true)
	if !EnableColors {
		t.Fatalf("expected EnableColors to be true, got false")
	}
	
	// Test with colors disabled
	SetEnableColors(false)
	if EnableColors {
		t.Fatalf("expected EnableColors to be false, got true")
	}
	
	// Restore original state
	SetEnableColors(originalState)
}

func TestFormatWithColor(t *testing.T) {
	// Test with colors enabled
	SetEnableColors(true)
	coloredText := formatWithColor(Red, "test")
	expected := "\033[31mtest\033[0m"
	if coloredText != expected {
		t.Errorf("expected colored text '%s', got '%s'", expected, coloredText)
	}
	
	// Test with colors disabled
	SetEnableColors(false)
	coloredText = formatWithColor(Red, "test")
	expected = "test"
	if coloredText != expected {
		t.Errorf("expected plain text 'test', got '%s'", coloredText)
	}
}

func TestTracedErrorStackTrace(t *testing.T) {
	err := errors.New("stack trace error")
	wrappedErr := WrapError(err, "", "")
	var te *TracedError
	if !errors.As(wrappedErr, &te) {
		t.Fatalf("expected *TracedError, got %T", wrappedErr)
	}
	
	if len(te.StackTrace) == 0 {
		t.Error("expected non-empty stack trace, got empty stack trace")
	}
}

func TestWrapErrorWithExistingTracedError(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := WrapError(originalErr, "payload", "message")
	doubleWrappedErr := WrapError(wrappedErr, "new payload", "new message")
	
	var te *TracedError
	if !errors.As(doubleWrappedErr, &te) {
		t.Fatalf("expected *TracedError, got %T", doubleWrappedErr)
	}
	
	if te.Payload != "new payload" {
		t.Errorf("expected payload to be 'new payload', got %s", te.Payload)
	}
	
	if !strings.Contains(te.UserMessage, "new message") {
		t.Errorf("expected user message to contain 'new message', got %s", te.UserMessage)
	}
}

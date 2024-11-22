package errorhandler

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// EnableColors determines whether colored output is enabled.
var EnableColors = true

// TracedError represents an error with additional context, user-friendly messages, and stack trace information.
type TracedError struct {
	OriginalError error
	UserMessage   string
	Payload       string
	StackTrace    []uintptr
}

// Error implements the error interface for TracedError, returning a formatted error message.
func (te *TracedError) Error() string {
	return fmt.Sprintf("%s\nUser Message: %s", te.OriginalError.Error(), te.UserMessage)
}

// FailOnError formats the error message for any given error.
// If the error is a TracedError, it includes additional context.
func FailOnError(err error) string {
	if err == nil {
		return ""
	}
	
	var tracedErr *TracedError
	if errors.As(err, &tracedErr) {
		return tracedErr.formatError()
	}
	
	return err.Error()
}

// captureCallers retrieves the call stack up to a certain depth.
func captureCallers() []uintptr {
	const stackDepth = 64
	var pcs [stackDepth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[:n]
}

func formatWithColor(color, text string) string {
	if EnableColors {
		return fmt.Sprintf("%s%s%s", color, text, Reset)
	}
	return text
}

// formatError formats the TracedError for detailed display, including payload, user message, and stack trace.
// optionally using colors for better readability.
func (te *TracedError) formatError() string {
	var sb strings.Builder
	
	// Apply payload if available
	if te.Payload != "" {
		// Using BrightYellow for better visibility of metadata
		sb.WriteString(fmt.Sprintf("%s: %s\n", formatWithColor(BrightYellow, "Payload"), te.Payload))
	}
	
	// Add error details header (using white for headers)
	sb.WriteString(fmt.Sprintf("%s\n", formatWithColor(BrightWhite, "Error Details")))
	
	// Original error in bright red for high visibility
	sb.WriteString(
		fmt.Sprintf(
			"%s: %s\n",
			formatWithColor(BrightRed, "Error"),
			formatWithColor(Red, te.OriginalError.Error()),
		),
	)
	
	// Add user message if available (using cyan for info)
	if te.UserMessage != "" {
		sb.WriteString(
			fmt.Sprintf(
				"%s: %s\n",
				formatWithColor(BrightCyan, "User Message"),
				formatWithColor(Cyan, te.UserMessage),
			),
		)
	}
	
	// Add stack trace header
	sb.WriteString(fmt.Sprintf("\n%s\n", formatWithColor(BrightWhite, "Stack Trace")))
	frames := runtime.CallersFrames(te.StackTrace)
	
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		
		// Skip runtime-related functions
		if strings.Contains(frame.Function, "runtime.") {
			continue
		}
		
		// Function name in bright blue for better visibility
		sb.WriteString(fmt.Sprintf("%s\n", formatWithColor(BrightBlue, frame.Function)))
		// File and line in dimmed white for less emphasis
		sb.WriteString(
			fmt.Sprintf(
				"\t%s:%d\n",
				formatWithColor(White, frame.File),
				frame.Line,
			),
		)
	}
	
	return sb.String()
}

// WrapError wraps an existing error with additional context, including payload and a user-friendly message.
func WrapError(err error, payload, message string) error {
	if err == nil {
		return nil
	}
	
	var existingTracedErr *TracedError
	if errors.As(err, &existingTracedErr) {
		// Merge existing TracedError context
		if payload == "" {
			payload = existingTracedErr.Payload
		}
		if message != "" && existingTracedErr.UserMessage != "" {
			message = fmt.Sprintf("%s\n%s", existingTracedErr.UserMessage, message)
		} else if message == "" {
			message = existingTracedErr.UserMessage
		}
		return &TracedError{
			OriginalError: existingTracedErr.OriginalError,
			UserMessage:   message,
			Payload:       payload,
			StackTrace:    captureCallers(),
		}
	}
	
	// Create a new TracedError
	return &TracedError{
		OriginalError: err,
		UserMessage:   message,
		Payload:       payload,
		StackTrace:    captureCallers(),
	}
}

// SetEnableColors sets the global EnableColors flag for controlling colored output.
func SetEnableColors(enable bool) {
	EnableColors = enable
}

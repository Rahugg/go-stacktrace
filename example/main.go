package main

import (
	"errors"
	"fmt"
	
	"github.com/Rahugg/go-stacktrace/errorhandler"
)

func main() {
	// Example usage
	err := errors.New("an example error")
	wrappedErr := errorhandler.WrapError(err, "some payload", "A user-friendly message")
	fmt.Println(errorhandler.FailOnError(wrappedErr))
	
	// Disable colors
	errorhandler.SetEnableColors(false)
	fmt.Println("With colors disabled:")
	fmt.Println(errorhandler.FailOnError(wrappedErr))
}

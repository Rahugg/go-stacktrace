# go-stacktrace

`go-stacktrace` is a robust Go package for enhanced error handling, providing stack traces, error wrapping, and user-friendly error messages with optional color formatting.

## Features

- 🔍 Detailed stack traces for improved debugging
- 📝 Custom user messages and payload support
- 🎨 Colorized error output (configurable)
- 🔄 Error wrapping with context preservation
- 🪶 Lightweight with zero external dependencies
- 🧪 Comprehensive test coverage

## Installation

```bash
go get github.com/Rahugg/go-stacktrace
```

## Usage

### Basic Error Wrapping

```go
import "github.com/Rahugg/go-stacktrace/errorhandler"

func main() {
    err := someFunction()
    if err != nil {
        wrappedErr := errorhandler.WrapError(
            err,
            "operation-failed",  // payload
            "Failed to process request"  // user message
        )
        fmt.Println(errorhandler.FailOnError(wrappedErr))
    }
}
```

### Controlling Color Output

```go
// Disable colored output
errorhandler.SetEnableColors(false)

// Enable colored output (default)
errorhandler.SetEnableColors(true)
```

### Error Information

The wrapped error includes:
- Original error message
- User-friendly message
- Custom payload (optional)
- Stack trace
- Color-coded output (optional)

## Example Output

```
Payload: operation-failed
Error Details
Original Error: file not found
User Message: Failed to process request

Stack Trace
main.someFunction
    /path/to/file.go:25
main.main
    /path/to/main.go:12
```

## Features in Detail

### TracedError Structure
```go
type TracedError struct {
    OriginalError error
    UserMessage   string
    Payload       string
    StackTrace    []uintptr
}
```

### Available Colors
- Red: Error messages
- Green: User messages
- Yellow: Payload information
- Blue: Function names in stack trace
- Magenta: Stack trace headers
- Cyan: Section headers

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
Or DM me.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Project Structure
```
.
├── errorhandler
│   ├── colors.go          # Color constants
│   ├── errorhandler.go    # Core functionality
│   └── errorhandler_test.go # Test suite
├── example
│   └── main.go           # Usage examples
├── go.mod
├── go.yml
└── README.md
```
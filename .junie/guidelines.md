# Zonotools Development Guidelines

This document provides essential information for developers working on the Zonotools project.

## Build/Configuration Instructions

### Prerequisites
- Go 1.22 or later

### Installation
1. Add the module to your project:
   ```bash
   go get github.com/isann/zonotools/v3
   ```

2. Import the package in your code:
   ```
   import "github.com/isann/zonotools/v3"
   ```

### Local Development Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/isann/zonotools.git
   cd zonotools
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

## Testing Information

### Running Tests
To run all tests:
```bash
go test ./...
```

To run specific tests:
```bash
go test -v -run TestFunctionName
```

To run tests with coverage:
```bash
go test -cover ./...
```

To generate a detailed coverage report:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Adding New Tests
1. Create a test file with the naming convention `filename_test.go` for the file `filename.go`.
2. Use the standard Go testing package and the testify/assert package for assertions.
3. Organize tests using subtests with `t.Run()` for better organization and readability.

Example:
```
func TestExampleFunction(t *testing.T) {
    t.Run("normal case", func(t *testing.T) {
        result := ExampleFunction("hello", "world", " ")
        assert.Equal(t, "hello world", result)
    })

    t.Run("with different separator", func(t *testing.T) {
        result := ExampleFunction("hello", "world", "-")
        assert.Equal(t, "hello-world", result)
    })

    t.Run("with empty strings", func(t *testing.T) {
        result := ExampleFunction("", "", ",")
        assert.Equal(t, ",", result)
    })
}
```

### Test Patterns
- Use descriptive test names that explain what is being tested
- Test both normal cases and edge cases
- For HTTP tests, create mock requests and test response handling
- For file operations, use temporary files/directories and clean up with defer
- Use table-driven tests for testing multiple inputs with the same logic

## Additional Development Information

### Code Style
- Follow standard Go code style and conventions
- Use meaningful variable and function names
- Add comments for non-obvious code, especially for complex algorithms
- Document exported functions with proper comments

### Project Structure
- The project is organized as a collection of utility functions for various purposes:
  - `date.go`: Date and time manipulation utilities
  - `http.go`: HTTP request handling utilities
  - `image.go`: Image processing utilities
  - `json.go`: JSON handling utilities
  - `serialize.go`: Serialization utilities
  - `snippet.go`: Miscellaneous code snippets

### Error Handling
- Return errors rather than using panic
- For HTTP utilities, properly handle edge cases like nil request bodies
- For file operations, ensure proper cleanup with defer statements

### Performance Considerations
- For file operations, use buffered I/O to handle large files efficiently
- Be mindful of memory usage, especially when processing large data

### Dependencies
- Main dependencies:
  - github.com/stretchr/testify: For test assertions
  - golang.org/x/image: For image processing

### Versioning
- The project follows semantic versioning
- Current major versions:
  - v3: For Go 1.18+
  - v2: For Go <1.18
  - v1: Legacy version

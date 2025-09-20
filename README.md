# Precise Calculator

A high-precision command-line calculator for decimal and hexadecimal arithmetic that maintains exact precision throughout all calculations.

## Features

- **Exact Precision**: No floating-point errors - `0.1 + 0.2` correctly equals `0.3`
- **Arbitrary Precision**: Handle numbers of any size limited only by available memory
- **Dual Number Systems**: Support for both decimal and hexadecimal numbers
- **Mathematical Operators**: Addition (+), subtraction (-), multiplication (x), division (/)
- **Operator Precedence**: Follows standard mathematical order of operations
- **Error Handling**: Clear error messages with proper exit codes
- **Library + CLI**: Core library with command-line interface

## Quick Start

### Build and Install

```bash
# Build the calculator
make build

# Run a calculation
./bin/precise-calc "0.1 + 0.2"
# Output: 0.3

# Install to system (optional)
make install
```

### Basic Usage

```bash
# Decimal arithmetic
./bin/precise-calc "5 + 3"          # Output: 8
./bin/precise-calc "0.1 + 0.2"      # Output: 0.3
./bin/precise-calc "15 / 3"         # Output: 5

# Hexadecimal numbers
./bin/precise-calc "0xFF + 1"       # Output: 256
./bin/precise-calc "0xAB + 0xCD"    # Output: 376

# Operator precedence
./bin/precise-calc "2 + 3 x 4"      # Output: 14

# Negative numbers
./bin/precise-calc "-5 + 3"         # Output: -2
./bin/precise-calc "-0xFF + 256"    # Output: 1

# Large numbers
./bin/precise-calc "999999999999999999999999999999 + 1"
# Output: 1000000000000000000000000000000
```

## Installation

### Prerequisites

- Go 1.19 or later
- Make (optional, for build automation)

### From Source

```bash
# Clone the repository
git clone <repository-url>
cd precise-calc

# Build
go build -o bin/precise-calc cmd/precise-calc/main.go

# Or use make
make build
```

## Usage

### Command Line Interface

```bash
precise-calc "<mathematical_expression>"
```

**Supported Characters:**
- Numbers: `0-9`, `A-F`, `a-f`
- Operators: `+`, `-`, `x`, `/`
- Hex prefix: `0x`, `-0x`
- Whitespace: spaces, tabs, newlines (ignored)

**Examples:**

```bash
# Basic calculations
precise-calc "5 + 3"
precise-calc "10 - 4"
precise-calc "3 x 7"
precise-calc "15 / 3"

# Decimal precision
precise-calc "0.0000000000000001 + 0.1"

# Mixed decimal and hex
precise-calc "0.5 + 0xFF"

# Complex expressions
precise-calc "2 x 3 + 4 x 5"  # Result: 26
```

### Error Handling

The calculator provides clear error messages and appropriate exit codes:

```bash
# Division by zero
precise-calc "5 / 0"
# Error: Division by zero
# Exit code: 1

# Invalid characters
precise-calc "5 + @"
# Error: Invalid character '@' at position 4
# Exit code: 1

# Empty expression
precise-calc ""
# Error: Empty expression provided
# Exit code: 1
```

## Library Usage

The calculator can also be used as a Go library:

```go
package main

import (
    "fmt"
    "precise-calc/pkg/calculator"
)

func main() {
    // Calculate an expression
    result, err := calculator.Calculate("0.1 + 0.2")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Format the result
    output := calculator.FormatRational(result)
    fmt.Printf("Result: %s\n", output) // Output: 0.3
}
```

### Library API

**Core Functions:**
- `Calculate(expression string) (*big.Rat, error)` - Evaluate mathematical expressions
- `ValidateExpression(expression string) error` - Validate expression format
- `FormatRational(result *big.Rat) string` - Format results for display

**Parsing Functions:**
- `ParseDecimal(s string) (*big.Rat, error)` - Parse decimal numbers
- `ParseHexadecimal(s string) (*big.Rat, error)` - Parse hexadecimal numbers
- `Tokenize(expression string) ([]Token, error)` - Tokenize expressions

## Development

### Project Structure

```
precise-calc/
├── cmd/precise-calc/          # CLI application entry point
├── pkg/calculator/            # Core calculator library
│   ├── calculator.go         # Main calculation logic
│   ├── types.go              # Data type definitions
│   ├── errors.go             # Error types
│   ├── tokenizer.go          # Expression tokenization
│   ├── parser.go             # Expression parsing
│   ├── evaluator.go          # Expression evaluation
│   └── number_parser.go      # Number parsing utilities
├── tests/
│   ├── contract/             # Contract tests for API compliance
│   ├── integration/          # End-to-end tests
│   └── unit/                # Unit tests for components
├── bin/                      # Built binaries
├── go.mod                    # Go module definition
├── Makefile                  # Build automation
└── README.md                 # This file
```

### Building and Testing

```bash
# Build the application
make build

# Run all tests
make test

# Run tests with coverage
make test-coverage

# Format code
make format

# Run linting
make vet

# Run all checks (format + vet + test)
make check

# Clean build artifacts
make clean
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run specific test suites
go test ./tests/contract/      # Contract tests
go test ./tests/integration/   # Integration tests
go test ./tests/unit/         # Unit tests

# Run with coverage
go test -cover ./...

# Verbose output
go test -v ./...
```

## Technical Details

### Precision and Accuracy

The calculator uses Go's `math/big.Rat` type for all calculations, ensuring:
- **Exact rational arithmetic** - no floating-point approximations
- **Arbitrary precision** - handle numbers of any size
- **Perfect decimal representation** - 0.1 + 0.2 = 0.3 exactly

### Performance

- **Simple expressions** (< 10 tokens): < 1ms
- **Complex expressions** (10-100 tokens): < 10ms
- **Large numbers** (1000+ digits): < 100ms
- **Memory usage**: Scales with number precision requirements

### Operator Precedence

Following standard mathematical conventions:
1. **Multiplication (x) and Division (/)** - Precedence 2
2. **Addition (+) and Subtraction (-)** - Precedence 1
3. **Left-to-right** evaluation for same precedence

Examples:
- `2 + 3 x 4` = `2 + (3 x 4)` = `2 + 12` = `14`
- `20 / 4 + 1` = `(20 / 4) + 1` = `5 + 1` = `6`

### Input Validation

The calculator strictly validates input:
- **Character set**: Only `[A-Fa-f0-9x+\-\s\t\n/.]` allowed
- **Number formats**: Valid decimal or hexadecimal only
- **Expression structure**: Must be well-formed mathematical expressions

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes following the existing code style
4. Add tests for new functionality
5. Run the test suite (`make test`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Code Style

- Follow Go conventions and idioms
- Use `go fmt` for formatting
- Ensure `go vet` passes without warnings
- Add comprehensive tests for new features
- Document public APIs with Go doc comments

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go's excellent `math/big` package for arbitrary precision arithmetic
- Follows library-first design principles for maximum reusability
- Test-driven development approach ensures reliability and maintainability
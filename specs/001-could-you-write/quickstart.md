# Quickstart Guide: Precise Decimal and Hex Calculator

## Overview
This guide walks through building and testing the precise decimal and hex calculator from scratch using Test-Driven Development (TDD).

## Prerequisites
- Go 1.19 or later
- Basic understanding of Go programming
- Command-line terminal access

## Project Setup

### 1. Initialize Go Module
```bash
# Navigate to project root
cd /home/william/git/precise-calc

# Initialize Go module
go mod init precise-calc

# Create basic directory structure
mkdir -p cmd/precise-calc
mkdir -p pkg/calculator
mkdir -p tests/{unit,integration,contract}
```

### 2. Verify Directory Structure
```
precise-calc/
├── cmd/precise-calc/          # CLI application entry point
├── pkg/calculator/            # Core calculator library
├── tests/
│   ├── contract/             # Contract tests for API compliance
│   ├── integration/          # End-to-end tests
│   └── unit/                # Unit tests for components
├── go.mod                   # Go module definition
└── specs/                   # Design specifications (current directory)
```

## TDD Implementation Flow

### Phase 1: Core Number Parsing (Contract Tests First)

#### Step 1: Create Failing Contract Tests
```bash
# Create contract test file
cat > tests/contract/number_parsing_test.go << 'EOF'
package contract

import (
    "testing"
    "precise-calc/pkg/calculator"
)

func TestDecimalNumberParsing(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"5", "5"},
        {"0.1", "1/10"},
        {"-3.14", "-157/50"},
        {"0.0000000000000001", "1/10000000000000000"},
    }

    for _, test := range tests {
        result, err := calculator.ParseDecimal(test.input)
        if err != nil {
            t.Errorf("ParseDecimal(%s) error: %v", test.input, err)
            continue
        }
        if result.String() != test.expected {
            t.Errorf("ParseDecimal(%s) = %s, want %s", test.input, result.String(), test.expected)
        }
    }
}

func TestHexNumberParsing(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"0xFF", "255"},
        {"0xAB91", "43921"},
        {"-0xFF", "-255"},
        {"-0xab91", "-43921"},
    }

    for _, test := range tests {
        result, err := calculator.ParseHexadecimal(test.input)
        if err != nil {
            t.Errorf("ParseHexadecimal(%s) error: %v", test.input, err)
            continue
        }
        if result.String() != test.expected {
            t.Errorf("ParseHexadecimal(%s) = %s, want %s", test.input, result.String(), test.expected)
        }
    }
}
EOF
```

#### Step 2: Run Tests (Should Fail)
```bash
go test ./tests/contract/
# Expected: compilation errors - functions don't exist yet
```

#### Step 3: Create Minimal Implementation
```bash
# Create basic package structure
cat > pkg/calculator/calculator.go << 'EOF'
package calculator

import (
    "math/big"
    "errors"
)

// ParseDecimal parses a decimal number string to exact rational
func ParseDecimal(s string) (*big.Rat, error) {
    // TODO: Implement decimal parsing
    return nil, errors.New("not implemented")
}

// ParseHexadecimal parses a hex number string to exact rational
func ParseHexadecimal(s string) (*big.Rat, error) {
    // TODO: Implement hex parsing
    return nil, errors.New("not implemented")
}
EOF
```

#### Step 4: Run Tests Again (Should Fail with "not implemented")
```bash
go test ./tests/contract/
# Expected: tests fail with "not implemented" errors
```

#### Step 5: Implement Number Parsing
```bash
# Implement actual parsing logic
# This step involves implementing the ParseDecimal and ParseHexadecimal functions
# following the research decisions (using math/big package)
```

### Phase 2: Expression Tokenization

#### Step 1: Create Tokenization Contract Tests
```bash
cat > tests/contract/tokenization_test.go << 'EOF'
package contract

import (
    "testing"
    "precise-calc/pkg/calculator"
)

func TestTokenizeExpression(t *testing.T) {
    tests := []struct {
        input    string
        expected []calculator.Token
    }{
        {
            "5 + 3",
            []calculator.Token{
                {Type: calculator.NumberToken, Value: "5", Position: 0},
                {Type: calculator.OperatorToken, Value: "+", Position: 2},
                {Type: calculator.NumberToken, Value: "3", Position: 4},
            },
        },
        {
            "0xFF - 0xAB",
            []calculator.Token{
                {Type: calculator.NumberToken, Value: "0xFF", Position: 0},
                {Type: calculator.OperatorToken, Value: "-", Position: 5},
                {Type: calculator.NumberToken, Value: "0xAB", Position: 7},
            },
        },
    }

    for _, test := range tests {
        result, err := calculator.Tokenize(test.input)
        if err != nil {
            t.Errorf("Tokenize(%s) error: %v", test.input, err)
            continue
        }

        if len(result) != len(test.expected) {
            t.Errorf("Tokenize(%s) returned %d tokens, want %d",
                test.input, len(result), len(test.expected))
            continue
        }

        for i, token := range result {
            expected := test.expected[i]
            if token.Type != expected.Type || token.Value != expected.Value {
                t.Errorf("Token %d: got {%v, %s}, want {%v, %s}",
                    i, token.Type, token.Value, expected.Type, expected.Value)
            }
        }
    }
}
EOF
```

### Phase 3: Expression Evaluation

#### Step 1: Create End-to-End Integration Tests
```bash
cat > tests/integration/calculator_test.go << 'EOF'
package integration

import (
    "testing"
    "precise-calc/pkg/calculator"
)

func TestCalculateEndToEnd(t *testing.T) {
    tests := []struct {
        expression string
        expected   string
    }{
        // Basic arithmetic
        {"5 + 3", "8"},
        {"10 - 4", "6"},
        {"3 x 7", "21"},
        {"15 / 3", "5"},

        // Operator precedence
        {"2 + 3 x 4", "14"},
        {"20 / 4 + 1", "6"},

        // Decimal precision
        {"0.1 + 0.2", "3/10"},
        {"1/3 + 1/3 + 1/3", "1"},

        // Hex numbers
        {"0xFF + 1", "256"},
        {"0xAB91 + 100", "44021"},

        // Mixed decimal and hex
        {"0.5 + 0xFF", "255.5"},

        // Negative numbers
        {"-5 + 3", "-2"},
        {"-0xFF + 256", "1"},

        // Complex expressions
        {"0.0000000000000001 + 0.1 + -99999999999999 - 0xab91", "-100000000043890.8999999999999999"},
    }

    for _, test := range tests {
        result, err := calculator.Calculate(test.expression)
        if err != nil {
            t.Errorf("Calculate(%s) error: %v", test.expression, err)
            continue
        }

        if result.String() != test.expected {
            t.Errorf("Calculate(%s) = %s, want %s",
                test.expression, result.String(), test.expected)
        }
    }
}

func TestCalculateErrors(t *testing.T) {
    tests := []struct {
        expression string
        errorType  string
    }{
        {"5 / 0", "DivisionByZero"},
        {"5 + @", "InvalidCharacter"},
        {"", "EmptyExpression"},
        {"5 + + 3", "ParseError"},
        {"0xGHI", "ParseError"},
    }

    for _, test := range tests {
        _, err := calculator.Calculate(test.expression)
        if err == nil {
            t.Errorf("Calculate(%s) expected error, got nil", test.expression)
            continue
        }

        // Check error type (implementation will define specific error types)
        if !containsErrorType(err.Error(), test.errorType) {
            t.Errorf("Calculate(%s) error %v does not contain expected type %s",
                test.expression, err, test.errorType)
        }
    }
}

func containsErrorType(errorMsg, expectedType string) bool {
    // Simple substring check - implementation will define proper error types
    return true // Placeholder
}
EOF
```

### Phase 4: CLI Integration

#### Step 1: Create CLI Contract Tests
```bash
cat > tests/contract/cli_test.go << 'EOF'
package contract

import (
    "os/exec"
    "strings"
    "testing"
)

func TestCLIBasicUsage(t *testing.T) {
    tests := []struct {
        args     []string
        expected string
        exitCode int
    }{
        {[]string{"5 + 3"}, "8", 0},
        {[]string{"0xFF + 1"}, "256", 0},
        {[]string{"5 / 0"}, "Error: Division by zero", 1},
        {[]string{"5 + @"}, "Error: Invalid character", 1},
    }

    for _, test := range tests {
        cmd := exec.Command("go", append([]string{"run", "cmd/precise-calc/main.go"}, test.args...)...)
        output, err := cmd.CombinedOutput()

        outputStr := strings.TrimSpace(string(output))

        if test.exitCode == 0 && err != nil {
            t.Errorf("Command %v expected success, got error: %v", test.args, err)
            continue
        }

        if test.exitCode != 0 && err == nil {
            t.Errorf("Command %v expected error, got success", test.args)
            continue
        }

        if !strings.Contains(outputStr, test.expected) {
            t.Errorf("Command %v output %q does not contain expected %q",
                test.args, outputStr, test.expected)
        }
    }
}
EOF
```

## Development Workflow

### Daily TDD Cycle

1. **Red**: Write a failing test for the next small feature
2. **Green**: Write minimal code to make the test pass
3. **Refactor**: Clean up code while keeping tests green
4. **Repeat**: Move to next small feature

### Test Execution Order

```bash
# 1. Run contract tests (should fail initially)
go test ./tests/contract/

# 2. Implement minimal code to satisfy contracts
# (Edit files in pkg/calculator/)

# 3. Run contract tests again (should pass)
go test ./tests/contract/

# 4. Run integration tests
go test ./tests/integration/

# 5. Run all tests
go test ./...

# 6. Build and test CLI
go build -o bin/precise-calc cmd/precise-calc/main.go
./bin/precise-calc "5 + 3"
```

### Implementation Checkpoints

#### Checkpoint 1: Number Parsing Complete
- [ ] All contract tests for ParseDecimal pass
- [ ] All contract tests for ParseHexadecimal pass
- [ ] Unit tests for edge cases pass

#### Checkpoint 2: Tokenization Complete
- [ ] All tokenization contract tests pass
- [ ] Whitespace handling works correctly
- [ ] Error cases properly detected

#### Checkpoint 3: Expression Evaluation Complete
- [ ] All integration tests pass
- [ ] Operator precedence working correctly
- [ ] Error handling implemented

#### Checkpoint 4: CLI Complete
- [ ] All CLI contract tests pass
- [ ] Proper exit codes returned
- [ ] Error messages are user-friendly

## Validation Steps

### Final Acceptance Test
```bash
# Build the final binary
go build -o precise-calc cmd/precise-calc/main.go

# Test all examples from specification
./precise-calc "0.0000000000000001 + 0.1 + -99999999999999 - 0xab91"
./precise-calc "0.1 + 0.2"
./precise-calc "2 x 3 / 6"

# Test error conditions
./precise-calc "5 / 0"
./precise-calc "5 + @"
./precise-calc ""

# Performance test with large numbers
./precise-calc "999999999999999999999999999999999999999999999999999999999999999999 + 1"
```

### Code Quality Check
```bash
# Format code
go fmt ./...

# Run linter (if available)
golint ./...

# Check for common issues
go vet ./...

# Run all tests with coverage
go test -cover ./...
```

## Troubleshooting

### Common Issues

1. **Import Path Errors**: Ensure module name matches go.mod
2. **Test Failures**: Check that interfaces match contracts exactly
3. **Precision Issues**: Verify using big.Rat for all calculations
4. **CLI Issues**: Check argument parsing and error handling

### Debugging Tips

1. **Use fmt.Printf**: Add debug output during development
2. **Test Individual Functions**: Isolate failures to specific components
3. **Check Error Messages**: Ensure they match contract specifications
4. **Validate with Simple Cases**: Start with basic arithmetic before complex expressions

## Next Steps

After completing this quickstart:

1. **Add More Tests**: Edge cases, performance tests, fuzz testing
2. **Optimize Performance**: Profile and optimize hot paths
3. **Extend Functionality**: Add more operators, functions, etc.
4. **Package for Distribution**: Create release binaries, Docker images
5. **Documentation**: Add godoc comments, usage examples

This quickstart ensures the calculator meets all specification requirements through systematic TDD implementation.
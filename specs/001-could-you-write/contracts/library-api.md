# Library API Contract

## Overview
Core library interface for the precise decimal and hex calculator. Provides programmatic access to calculation functionality for CLI and potential future integrations.

## Package Structure

### Main Package: `calculator`

#### Core Functions

##### `Calculate(expression string) (*big.Rat, error)`
**Purpose**: Evaluate a mathematical expression and return the exact result

**Input**:
- `expression`: String containing mathematical expression with decimal/hex numbers and operators

**Output**:
- `*big.Rat`: Exact rational result of the calculation
- `error`: Detailed error information if calculation fails

**Error Types**:
- `ParseError`: Invalid expression format
- `DivisionByZeroError`: Attempt to divide by zero
- `InvalidCharacterError`: Character outside allowed set
- `EmptyExpressionError`: Empty or whitespace-only input

##### `ValidateExpression(expression string) error`
**Purpose**: Validate expression format without performing calculation

**Input**:
- `expression`: String to validate

**Output**:
- `error`: Validation error details, or nil if valid

##### `FormatResult(result *big.Rat, precision int) string`
**Purpose**: Format calculation result for display

**Input**:
- `result`: Rational number result from calculation
- `precision`: Maximum decimal places (0 = exact rational display)

**Output**:
- `string`: Formatted result string

## Type Definitions

### Core Types

```go
package calculator

import (
    "math/big"
    "fmt"
)

// Number represents a parsed numeric value
type Number struct {
    Value    *big.Rat
    Original string
    Type     NumberType
}

// Operator represents a mathematical operation
type Operator struct {
    Symbol       rune
    Precedence   int
    Associativity Associativity
}

// Token represents a parsed element from input
type Token struct {
    Type     TokenType
    Value    string
    Position int
}

// Expression represents a complete mathematical expression
type Expression struct {
    Original      string
    Tokens        []Token
    PostfixTokens []Token
    Result        *big.Rat
}
```

### Enumerations

```go
type NumberType int
const (
    Decimal NumberType = iota
    Hexadecimal
)

type TokenType int
const (
    NumberToken TokenType = iota
    OperatorToken
    WhitespaceToken
)

type Associativity int
const (
    Left Associativity = iota
    Right
)
```

### Error Types

```go
// ParseError represents expression parsing failures
type ParseError struct {
    Message  string
    Position int
    Context  string
}

func (e ParseError) Error() string {
    return fmt.Sprintf("Parse error at position %d: %s", e.Position, e.Message)
}

// DivisionByZeroError represents division by zero attempts
type DivisionByZeroError struct {
    Position int
}

func (e DivisionByZeroError) Error() string {
    return "Division by zero"
}

// InvalidCharacterError represents forbidden characters in input
type InvalidCharacterError struct {
    Character rune
    Position  int
}

func (e InvalidCharacterError) Error() string {
    return fmt.Sprintf("Invalid character '%c' at position %d", e.Character, e.Position)
}

// EmptyExpressionError represents empty input
type EmptyExpressionError struct{}

func (e EmptyExpressionError) Error() string {
    return "Empty expression provided"
}
```

## API Usage Examples

### Basic Calculation
```go
package main

import (
    "fmt"
    "math/big"
    "calculator"
)

func main() {
    result, err := calculator.Calculate("0.1 + 0.2")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // result is *big.Rat with exact value 3/10
    fmt.Printf("Result: %s\n", calculator.FormatResult(result, 10))
    // Output: Result: 0.3
}
```

### Error Handling
```go
result, err := calculator.Calculate("5 / 0")
if err != nil {
    switch e := err.(type) {
    case calculator.DivisionByZeroError:
        fmt.Println("Cannot divide by zero")
    case calculator.ParseError:
        fmt.Printf("Parse error: %s\n", e.Message)
    case calculator.InvalidCharacterError:
        fmt.Printf("Invalid character '%c' at position %d\n", e.Character, e.Position)
    default:
        fmt.Printf("Unknown error: %v\n", err)
    }
}
```

### Validation Only
```go
err := calculator.ValidateExpression("0xAB + 0xCD")
if err != nil {
    fmt.Printf("Invalid expression: %v\n", err)
} else {
    fmt.Println("Expression is valid")
}
```

### Hex Number Handling
```go
result, err := calculator.Calculate("-0xFF + 256")
if err != nil {
    panic(err)
}

// Convert result to different formats
decimal := calculator.FormatResult(result, 0)  // "1"
fraction := result.String()                    // "1/1"
float, _ := result.Float64()                   // 1.0 (may lose precision)
```

## Internal APIs

### Parsing Components

#### `Tokenize(expression string) ([]Token, error)`
**Purpose**: Convert input string into sequence of tokens
**Internal Use**: Called by Calculate function

#### `ParseExpression(tokens []Token) (*Expression, error)`
**Purpose**: Convert tokens into validated expression structure
**Internal Use**: Called after tokenization

#### `InfixToPostfix(tokens []Token) ([]Token, error)`
**Purpose**: Convert infix notation to postfix for evaluation
**Internal Use**: Handles operator precedence

#### `EvaluatePostfix(tokens []Token) (*big.Rat, error)`
**Purpose**: Evaluate postfix expression to get final result
**Internal Use**: Performs actual arithmetic

### Number Parsing

#### `ParseDecimal(s string) (*big.Rat, error)`
**Purpose**: Parse decimal number string to exact rational
**Internal Use**: Handles decimal number conversion

#### `ParseHexadecimal(s string) (*big.Rat, error)`
**Purpose**: Parse hexadecimal number string to exact rational
**Internal Use**: Handles hex number conversion with sign

## Contract Tests

### API Function Tests

#### Calculate Function
```go
func TestCalculateBasicArithmetic(t *testing.T) {
    tests := []struct{
        input    string
        expected string
    }{
        {"5 + 3", "8"},
        {"10 - 7", "3"},
        {"4 x 6", "24"},
        {"15 / 3", "5"},
    }

    for _, test := range tests {
        result, err := calculator.Calculate(test.input)
        assert.NoError(t, err)
        assert.Equal(t, test.expected, result.String())
    }
}
```

#### Error Handling Tests
```go
func TestCalculateErrors(t *testing.T) {
    _, err := calculator.Calculate("5 / 0")
    assert.IsType(t, calculator.DivisionByZeroError{}, err)

    _, err = calculator.Calculate("5 + @")
    assert.IsType(t, calculator.InvalidCharacterError{}, err)

    _, err = calculator.Calculate("")
    assert.IsType(t, calculator.EmptyExpressionError{}, err)
}
```

#### Precision Tests
```go
func TestPrecisionMaintenance(t *testing.T) {
    // Test that 0.1 + 0.2 = 0.3 exactly
    result, err := calculator.Calculate("0.1 + 0.2")
    assert.NoError(t, err)

    expected := big.NewRat(3, 10)
    assert.Equal(t, 0, result.Cmp(expected))
}
```

#### Hex Number Tests
```go
func TestHexNumbers(t *testing.T) {
    result, err := calculator.Calculate("0xFF + 1")
    assert.NoError(t, err)
    assert.Equal(t, "256", result.String())

    result, err = calculator.Calculate("-0xA + 5")
    assert.NoError(t, err)
    assert.Equal(t, "-5", result.String())
}
```

## Performance Contracts

### Time Complexity
- **Tokenization**: O(n) where n is input length
- **Parsing**: O(n) for expression validation
- **Evaluation**: O(k) where k is number of operations
- **Overall**: O(n + k) for typical expressions

### Space Complexity
- **Token storage**: O(n) for input representation
- **Number precision**: O(p) where p is precision requirements
- **Call stack**: O(d) where d is expression depth

### Memory Management
- All `*big.Rat` allocations managed by caller
- No internal memory leaks
- Proper cleanup of intermediate calculations

## Thread Safety

### Concurrent Access
- All public functions are **thread-safe**
- No shared mutable state between calls
- Each calculation is independent

### Resource Sharing
- Read-only operator definitions are safe to share
- Token parsing creates new instances
- No global state modifications
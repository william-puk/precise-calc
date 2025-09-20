package calculator

import "fmt"

// ParseError represents expression parsing failures
type ParseError struct {
	Message  string
	Position int
	Context  string
}

func (e ParseError) Error() string {
	if e.Position >= 0 {
		return fmt.Sprintf("Parse error at position %d: %s", e.Position, e.Message)
	}
	return fmt.Sprintf("Parse error: %s", e.Message)
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

package calculator

import (
	"regexp"
	"strings"
	"unicode"
)

// ValidCharacterSet defines allowed characters for input validation
var ValidCharacterSet = regexp.MustCompile(`^[A-Fa-f0-9x+\-\s\t\n/.]*$`)

// Tokenize converts input string into sequence of tokens
func Tokenize(expression string) ([]Token, error) {
	if strings.TrimSpace(expression) == "" {
		return nil, EmptyExpressionError{}
	}

	// Validate character set
	if !ValidCharacterSet.MatchString(expression) {
		// Find first invalid character
		for i, ch := range expression {
			if !isValidCharacter(ch) {
				return nil, InvalidCharacterError{Character: ch, Position: i}
			}
		}
	}

	tokens := []Token{}
	i := 0
	runes := []rune(expression)

	for i < len(runes) {
		ch := runes[i]

		// Skip whitespace
		if unicode.IsSpace(ch) {
			i++
			continue
		}

		// Handle numbers (decimal or hex) FIRST - including negative numbers
		if isDigit(ch) || ch == '.' ||
			(ch == '0' && i+1 < len(runes) && (runes[i+1] == 'x' || runes[i+1] == 'X')) ||
			(ch == '-' && isStartOfNumber(runes, i, tokens)) {
			start := i
			value, newPos := parseNumberToken(runes, i)
			tokens = append(tokens, Token{
				Type:     NumberToken,
				Value:    value,
				Position: start,
			})
			i = newPos
			continue
		}

		// Handle operators AFTER checking for negative numbers
		if isOperator(ch) {
			tokens = append(tokens, Token{
				Type:     OperatorToken,
				Value:    string(ch),
				Position: i,
			})
			i++
			continue
		}

		// If we get here, it's an invalid character
		return nil, InvalidCharacterError{Character: ch, Position: i}
	}

	return tokens, nil
}

// isValidCharacter checks if character is in allowed set
func isValidCharacter(ch rune) bool {
	return (ch >= 'A' && ch <= 'F') ||
		(ch >= 'a' && ch <= 'f') ||
		(ch >= '0' && ch <= '9') ||
		ch == 'x' || ch == 'X' ||
		ch == '+' || ch == '-' ||
		ch == '.' || ch == '/' ||
		unicode.IsSpace(ch)
}

// isOperator checks if character is a mathematical operator
func isOperator(ch rune) bool {
	_, exists := OperatorMap[ch]
	return exists
}

// isDigit checks if character is a digit
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// isHexDigit checks if character is a hexadecimal digit
func isHexDigit(ch rune) bool {
	return isDigit(ch) || (ch >= 'A' && ch <= 'F') || (ch >= 'a' && ch <= 'f')
}

// isStartOfNumber checks if a minus sign should be treated as part of a number
func isStartOfNumber(runes []rune, i int, tokens []Token) bool {
	// Must be a minus sign
	if i >= len(runes) || runes[i] != '-' {
		return false
	}

	// Must have something after the minus
	if i+1 >= len(runes) {
		return false
	}

	// Next character must be a digit or start of hex number
	next := runes[i+1]
	if isDigit(next) || next == '.' {
		// This is negative decimal number if we're at start or after operator
		return len(tokens) == 0 || tokens[len(tokens)-1].Type == OperatorToken
	}

	// Check for negative hex number: -0x
	if next == '0' && i+2 < len(runes) && (runes[i+2] == 'x' || runes[i+2] == 'X') {
		return len(tokens) == 0 || tokens[len(tokens)-1].Type == OperatorToken
	}

	return false
}

// parseNumberToken parses a number token starting at position i
func parseNumberToken(runes []rune, i int) (string, int) {
	start := i

	// Handle negative sign
	if i < len(runes) && runes[i] == '-' {
		i++
	}

	// Check for hex number
	if i+1 < len(runes) && runes[i] == '0' && (runes[i+1] == 'x' || runes[i+1] == 'X') {
		i += 2 // Skip 0x
		// Parse hex digits
		for i < len(runes) && isHexDigit(runes[i]) {
			i++
		}
		return string(runes[start:i]), i
	}

	// Parse decimal number
	// Integer part
	for i < len(runes) && isDigit(runes[i]) {
		i++
	}

	// Decimal point and fractional part
	if i < len(runes) && runes[i] == '.' {
		i++
		for i < len(runes) && isDigit(runes[i]) {
			i++
		}
	}

	return string(runes[start:i]), i
}

package unit

import (
	"precise-calc/pkg/calculator"
	"testing"
)

func TestTokenizeEdgeCases(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
		description string
	}{
		{"", true, "empty string"},
		{"   ", true, "whitespace only"},
		{"5 @ 3", true, "invalid character"},
		{"5 + ", false, "trailing operator (caught later in parsing)"},
		{"+ 5", false, "leading operator (caught later in parsing)"},
		{"5++3", false, "double operator (caught later in parsing)"},
		{"5  +  3", false, "multiple spaces"},
		{"5\t+\n3", false, "mixed whitespace"},
		{"0xFF+0xAB", false, "no spaces between hex"},
		{"-5+-3", false, "negative numbers with operators"},
		{"123.456x789.012", false, "decimals with multiplication"},
	}

	for _, test := range tests {
		_, err := calculator.Tokenize(test.input)
		hasError := err != nil

		if hasError != test.expectError {
			if test.expectError {
				t.Errorf("Tokenize(%s) expected error for %s, got none", test.input, test.description)
			} else {
				t.Errorf("Tokenize(%s) unexpected error for %s: %v", test.input, test.description, err)
			}
		}
	}
}

func TestTokenizePositionAccuracy(t *testing.T) {
	input := "  5  +  0xFF  "
	tokens, err := calculator.Tokenize(input)
	if err != nil {
		t.Fatalf("Tokenize failed: %v", err)
	}

	expectedPositions := []int{2, 5, 8}
	if len(tokens) != len(expectedPositions) {
		t.Fatalf("Expected %d tokens, got %d", len(expectedPositions), len(tokens))
	}

	for i, token := range tokens {
		if token.Position != expectedPositions[i] {
			t.Errorf("Token %d position = %d, want %d", i, token.Position, expectedPositions[i])
		}
	}
}

func TestTokenizeInvalidCharacters(t *testing.T) {
	invalidChars := []string{"@", "#", "$", "%", "^", "&", "*", "(", ")", "=", "!", "~", "`"}

	for _, char := range invalidChars {
		input := "5 + " + char
		_, err := calculator.Tokenize(input)
		if err == nil {
			t.Errorf("Tokenize(%s) expected error for invalid character %s", input, char)
		}

		// Check if it's the right type of error
		if invalidCharErr, ok := err.(calculator.InvalidCharacterError); ok {
			if invalidCharErr.Character != rune(char[0]) {
				t.Errorf("Expected error for character '%s', got error for '%c'", char, invalidCharErr.Character)
			}
		} else {
			t.Errorf("Expected InvalidCharacterError, got %T", err)
		}
	}
}

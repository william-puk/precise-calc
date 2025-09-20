package contract

import (
	"precise-calc/pkg/calculator"
	"testing"
)

func TestCalculateBasicArithmetic(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5 + 3", "8"},
		{"10 - 4", "6"},
		{"3 x 7", "21"},
		{"15 / 3", "5"},
	}

	for _, test := range tests {
		result, err := calculator.Calculate(test.input)
		if err != nil {
			t.Errorf("Calculate(%s) error: %v", test.input, err)
			continue
		}
		formatted := calculator.FormatRational(result)
		if formatted != test.expected {
			t.Errorf("Calculate(%s) = %s, want %s", test.input, formatted, test.expected)
		}
	}
}

func TestCalculateOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2 + 3 x 4", "14"},
		{"20 / 4 + 1", "6"},
		{"2 x 3 + 4 x 5", "26"},
		{"100 - 10 x 5", "50"},
	}

	for _, test := range tests {
		result, err := calculator.Calculate(test.input)
		if err != nil {
			t.Errorf("Calculate(%s) error: %v", test.input, err)
			continue
		}
		formatted := calculator.FormatRational(result)
		if formatted != test.expected {
			t.Errorf("Calculate(%s) = %s, want %s", test.input, formatted, test.expected)
		}
	}
}

func TestCalculateDecimalPrecision(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0.1 + 0.2", "3/10"},
		{"0.0000000000000001 + 0.1", "1000000000000001/10000000000000000"},
		{"1/3 + 1/3 + 1/3", "1"},
	}

	for _, test := range tests {
		result, err := calculator.Calculate(test.input)
		if err != nil {
			t.Errorf("Calculate(%s) error: %v", test.input, err)
			continue
		}
		formatted := calculator.FormatRational(result)
		if formatted != test.expected {
			t.Errorf("Calculate(%s) = %s, want %s", test.input, formatted, test.expected)
		}
	}
}

func TestCalculateHexNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0xFF + 1", "256"},
		{"0xAB91 + 100", "44021"},
		{"-0xFF + 256", "1"},
		{"0x0 + 5", "5"},
	}

	for _, test := range tests {
		result, err := calculator.Calculate(test.input)
		if err != nil {
			t.Errorf("Calculate(%s) error: %v", test.input, err)
			continue
		}
		formatted := calculator.FormatRational(result)
		if formatted != test.expected {
			t.Errorf("Calculate(%s) = %s, want %s", test.input, formatted, test.expected)
		}
	}
}

func TestCalculateErrorHandling(t *testing.T) {
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

		// Note: actual error type checking will be implemented with specific error types
		if err.Error() == "" {
			t.Errorf("Calculate(%s) error message is empty", test.expression)
		}
	}
}

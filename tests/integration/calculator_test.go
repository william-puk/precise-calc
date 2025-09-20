package integration

import (
	"precise-calc/pkg/calculator"
	"testing"
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
		{"0.5 + 0xFF", "511/2"},

		// Negative numbers
		{"-5 + 3", "-2"},
		{"-0xFF + 256", "1"},

		// Complex expressions
		{"0.0000000000000001 + 0.1 + -99999999999999 - 0xab91", "-1000000000439198999999999999999/10000000000000000"},
	}

	for _, test := range tests {
		result, err := calculator.Calculate(test.expression)
		if err != nil {
			t.Errorf("Calculate(%s) error: %v", test.expression, err)
			continue
		}

		formatted := calculator.FormatRational(result)
		if formatted != test.expected {
			t.Errorf("Calculate(%s) = %s, want %s",
				test.expression, formatted, test.expected)
		}
	}
}

func TestCalculateComplexPrecedence(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
	}{
		{"2 + 3 x 4 - 1", "13"},
		{"100 / 10 + 2 x 5", "20"},
		{"1 + 2 x 3 + 4 x 5", "27"},
		{"10 - 2 x 3 + 1", "5"},
	}

	for _, test := range tests {
		result, err := calculator.Calculate(test.expression)
		if err != nil {
			t.Errorf("Calculate(%s) error: %v", test.expression, err)
			continue
		}

		formatted := calculator.FormatRational(result)
		if formatted != test.expected {
			t.Errorf("Calculate(%s) = %s, want %s",
				test.expression, formatted, test.expected)
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
		{"5 +", "ParseError"},
		{"+ 5", "ParseError"},
	}

	for _, test := range tests {
		_, err := calculator.Calculate(test.expression)
		if err == nil {
			t.Errorf("Calculate(%s) expected error, got nil", test.expression)
			continue
		}

		// Basic error existence check - specific error types will be validated later
		if err.Error() == "" {
			t.Errorf("Calculate(%s) error message is empty", test.expression)
		}
	}
}

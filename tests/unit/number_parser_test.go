package unit

import (
	"precise-calc/pkg/calculator"
	"testing"
)

func TestParseDecimalEdgeCases(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
		description string
	}{
		{"", true, "empty string"},
		{"   ", true, "whitespace only"},
		{"abc", true, "non-numeric"},
		{"123.456.789", true, "multiple decimal points"},
		{".", true, "decimal point only"},
		{"123.", false, "trailing decimal point"},
		{".123", false, "leading decimal point"},
		{"0", false, "zero"},
		{"-0", false, "negative zero"},
		{"123e10", true, "scientific notation not supported"},
		{"123.456789012345678901234567890", false, "very long decimal"},
	}

	for _, test := range tests {
		_, err := calculator.ParseDecimal(test.input)
		hasError := err != nil

		if hasError != test.expectError {
			if test.expectError {
				t.Errorf("ParseDecimal(%s) expected error for %s, got none", test.input, test.description)
			} else {
				t.Errorf("ParseDecimal(%s) unexpected error for %s: %v", test.input, test.description, err)
			}
		}
	}
}

func TestParseHexadecimalEdgeCases(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
		description string
	}{
		{"", true, "empty string"},
		{"   ", true, "whitespace only"},
		{"0x", true, "0x without digits"},
		{"x123", true, "missing 0x prefix"},
		{"0xGHI", true, "invalid hex digits"},
		{"0x0", false, "hex zero"},
		{"-0x0", false, "negative hex zero"},
		{"0XABC", false, "uppercase X"},
		{"-0XDEF", false, "negative uppercase"},
		{"0xFFFFFFFFFFFFFFFF", false, "large hex number"},
		{"0x1234567890ABCDEF", false, "mixed case hex"},
	}

	for _, test := range tests {
		_, err := calculator.ParseHexadecimal(test.input)
		hasError := err != nil

		if hasError != test.expectError {
			if test.expectError {
				t.Errorf("ParseHexadecimal(%s) expected error for %s, got none", test.input, test.description)
			} else {
				t.Errorf("ParseHexadecimal(%s) unexpected error for %s: %v", test.input, test.description, err)
			}
		}
	}
}

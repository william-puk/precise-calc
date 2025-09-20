package contract

import (
	"precise-calc/pkg/calculator"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5", "5"},
		{"0.1", "1/10"},
		{"-3.14", "-157/50"},
		{"0.0000000000000001", "1/10000000000000000"},
		{"123.456", "15432/125"},
		{"-0.001", "-1/1000"},
	}

	for _, test := range tests {
		result, err := calculator.ParseDecimal(test.input)
		if err != nil {
			t.Errorf("ParseDecimal(%s) error: %v", test.input, err)
			continue
		}
		formatted := calculator.FormatRational(result)
		if formatted != test.expected {
			t.Errorf("ParseDecimal(%s) = %s, want %s", test.input, formatted, test.expected)
		}
	}
}

func TestParseHexadecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0xFF", "255"},
		{"0xAB91", "43921"},
		{"-0xFF", "-255"},
		{"-0xab91", "-43921"},
		{"0x0", "0"},
		{"0xA", "10"},
		{"-0xA", "-10"},
	}

	for _, test := range tests {
		result, err := calculator.ParseHexadecimal(test.input)
		if err != nil {
			t.Errorf("ParseHexadecimal(%s) error: %v", test.input, err)
			continue
		}
		formatted := calculator.FormatRational(result)
		if formatted != test.expected {
			t.Errorf("ParseHexadecimal(%s) = %s, want %s", test.input, formatted, test.expected)
		}
	}
}

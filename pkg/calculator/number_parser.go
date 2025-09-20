package calculator

import (
	"math/big"
	"strings"
)

// ParseDecimal parses a decimal number string to exact rational
func ParseDecimal(s string) (*big.Rat, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, ParseError{Message: "Empty decimal number", Position: 0}
	}

	// Check for scientific notation (not supported)
	if strings.Contains(s, "e") || strings.Contains(s, "E") {
		return nil, ParseError{Message: "Scientific notation not supported", Position: 0}
	}

	// Use big.Rat to parse decimal numbers with exact precision
	rat := new(big.Rat)
	_, ok := rat.SetString(s)
	if !ok {
		return nil, ParseError{Message: "Invalid decimal format", Position: 0}
	}

	return rat, nil
}

// ParseHexadecimal parses a hexadecimal number string to exact rational
func ParseHexadecimal(s string) (*big.Rat, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, ParseError{Message: "Empty hex number", Position: 0}
	}

	// Handle negative sign
	negative := false
	if strings.HasPrefix(s, "-") {
		negative = true
		s = s[1:]
	}

	// Check for hex prefix
	if !strings.HasPrefix(s, "0x") && !strings.HasPrefix(s, "0X") {
		return nil, ParseError{Message: "Hex number must start with 0x", Position: 0}
	}

	// Remove 0x prefix
	hexDigits := s[2:]
	if hexDigits == "" {
		return nil, ParseError{Message: "No hex digits after 0x", Position: 2}
	}

	// Parse as base 16 using big.Int to handle large numbers
	bigInt := new(big.Int)
	_, ok := bigInt.SetString(hexDigits, 16)
	if !ok {
		return nil, ParseError{Message: "Invalid hex digits", Position: 2}
	}

	if negative {
		bigInt.Neg(bigInt)
	}

	// Convert to rational number
	rat := new(big.Rat)
	rat.SetInt(bigInt)
	return rat, nil
}

package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"precise-calc/pkg/calculator"
)

func main() {
	// Check command line arguments
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s \"<expression>\"\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s \"0.1 + 0.2\"\n", os.Args[0])
		os.Exit(1)
	}

	// Get the expression from command line arguments
	expression := os.Args[1]

	// Calculate the result
	result, err := calculator.Calculate(expression)
	if err != nil {
		handleError(err)
		os.Exit(1)
	}

	// Format and output the result
	output := formatOutput(result)
	fmt.Println(output)
}

// handleError formats and outputs error messages
func handleError(err error) {
	switch e := err.(type) {
	case calculator.DivisionByZeroError:
		fmt.Fprintf(os.Stderr, "Error: Division by zero\n")
	case calculator.InvalidCharacterError:
		fmt.Fprintf(os.Stderr, "Error: Invalid character '%c' at position %d\n", e.Character, e.Position)
	case calculator.ParseError:
		if e.Position >= 0 {
			fmt.Fprintf(os.Stderr, "Error: %s at position %d\n", e.Message, e.Position)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %s\n", e.Message)
		}
	case calculator.EmptyExpressionError:
		fmt.Fprintf(os.Stderr, "Error: Empty expression provided\n")
	default:
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}

// formatOutput formats the result for display
func formatOutput(result *big.Rat) string {
	// Try to display as a simple decimal if possible
	if result.IsInt() {
		return result.Num().String()
	}

	// Check if it's a simple fraction that can be displayed as decimal
	float, exact := result.Float64()
	if exact && isSimpleDecimal(float) {
		return formatDecimal(float)
	}

	// For fractions that are simple decimals (like 3/10 = 0.3), convert to decimal
	if canDisplayAsDecimal(result) {
		return convertToDecimal(result)
	}

	// For complex fractions, display as exact rational
	return result.String()
}

// isSimpleDecimal checks if a float can be displayed as a simple decimal
func isSimpleDecimal(f float64) bool {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	return !strings.Contains(str, "e") && len(str) <= 20
}

// formatDecimal formats a float as a clean decimal string
func formatDecimal(f float64) string {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	// Remove trailing zeros
	if strings.Contains(str, ".") {
		str = strings.TrimRight(str, "0")
		str = strings.TrimRight(str, ".")
	}
	return str
}

// canDisplayAsDecimal checks if a rational can be displayed as a simple decimal
func canDisplayAsDecimal(r *big.Rat) bool {
	// Check if denominator has only factors of 2 and 5 (powers of 10)
	denom := new(big.Int).Set(r.Denom())

	// Remove factors of 2
	two := big.NewInt(2)
	for {
		q, rem := new(big.Int).DivMod(denom, two, new(big.Int))
		if rem.Sign() != 0 {
			break
		}
		denom = q
	}

	// Remove factors of 5
	five := big.NewInt(5)
	for {
		q, rem := new(big.Int).DivMod(denom, five, new(big.Int))
		if rem.Sign() != 0 {
			break
		}
		denom = q
	}

	// If what's left is 1, then it can be displayed as decimal
	return denom.Cmp(big.NewInt(1)) == 0
}

// convertToDecimal converts a rational to decimal representation
func convertToDecimal(r *big.Rat) string {
	float, _ := r.Float64()
	return formatDecimal(float)
}

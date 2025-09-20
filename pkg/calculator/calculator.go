package calculator

import (
	"math/big"
	"strconv"
	"strings"
)

// Calculate evaluates a mathematical expression and returns the exact result
func Calculate(expression string) (*big.Rat, error) {
	// Store original for error reporting
	original := expression

	// Trim whitespace
	expression = strings.TrimSpace(expression)
	if expression == "" {
		return nil, EmptyExpressionError{}
	}

	// Tokenize the expression
	tokens, err := Tokenize(expression)
	if err != nil {
		return nil, err
	}

	// Parse and validate expression structure
	expr, err := ParseExpression(tokens)
	if err != nil {
		return nil, err
	}

	// Store original in expression
	expr.Original = original

	// Evaluate the postfix expression
	result, err := EvaluatePostfix(expr.PostfixTokens)
	if err != nil {
		return nil, err
	}

	// Store result in expression
	expr.Result = result

	return result, nil
}

// FormatResult formats calculation result for display
func FormatResult(result *big.Rat, precision int) string {
	if precision == 0 {
		// Return simplified representation
		if result.IsInt() {
			return result.Num().String()
		}
		return result.String()
	}

	// Convert to decimal with specified precision
	float, _ := result.Float64()
	return strconv.FormatFloat(float, 'f', precision, 64)
}

// FormatRational formats a rational number in the expected test format
func FormatRational(result *big.Rat) string {
	if result.IsInt() {
		return result.Num().String()
	}
	return result.String()
}

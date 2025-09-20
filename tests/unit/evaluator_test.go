package unit

import (
	"math/big"
	"precise-calc/pkg/calculator"
	"testing"
)

func TestEvaluatePostfixEdgeCases(t *testing.T) {
	tests := []struct {
		description string
		tokens      []calculator.Token
		expectError bool
	}{
		{
			"single number",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "5", Position: 0},
			},
			false,
		},
		{
			"missing operand",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "5", Position: 0},
				{Type: calculator.OperatorToken, Value: "+", Position: 2},
			},
			true,
		},
		{
			"too many operands",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "5", Position: 0},
				{Type: calculator.NumberToken, Value: "3", Position: 2},
				{Type: calculator.NumberToken, Value: "2", Position: 4},
				{Type: calculator.OperatorToken, Value: "+", Position: 6},
			},
			true,
		},
		{
			"division by zero",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "5", Position: 0},
				{Type: calculator.NumberToken, Value: "0", Position: 2},
				{Type: calculator.OperatorToken, Value: "/", Position: 4},
			},
			true,
		},
		{
			"valid expression",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "5", Position: 0},
				{Type: calculator.NumberToken, Value: "3", Position: 2},
				{Type: calculator.OperatorToken, Value: "+", Position: 4},
			},
			false,
		},
	}

	for _, test := range tests {
		_, err := calculator.EvaluatePostfix(test.tokens)
		hasError := err != nil

		if hasError != test.expectError {
			if test.expectError {
				t.Errorf("%s: expected error, got none", test.description)
			} else {
				t.Errorf("%s: unexpected error: %v", test.description, err)
			}
		}
	}
}

func TestEvaluatePostfixPrecision(t *testing.T) {
	// Test that precision is maintained through operations
	tokens := []calculator.Token{
		{Type: calculator.NumberToken, Value: "0.1", Position: 0},
		{Type: calculator.NumberToken, Value: "0.2", Position: 4},
		{Type: calculator.OperatorToken, Value: "+", Position: 8},
	}

	result, err := calculator.EvaluatePostfix(tokens)
	if err != nil {
		t.Fatalf("EvaluatePostfix failed: %v", err)
	}

	// 0.1 + 0.2 should equal exactly 3/10
	expected := big.NewRat(3, 10)
	if result.Cmp(expected) != 0 {
		t.Errorf("0.1 + 0.2 = %s, want %s", result, expected)
	}
}

func TestOperatorPrecedenceInEvaluation(t *testing.T) {
	// Test expression: 2 + 3 x 4 = 2 + 12 = 14
	// In postfix: 2 3 4 x +
	tokens := []calculator.Token{
		{Type: calculator.NumberToken, Value: "2", Position: 0},
		{Type: calculator.NumberToken, Value: "3", Position: 2},
		{Type: calculator.NumberToken, Value: "4", Position: 4},
		{Type: calculator.OperatorToken, Value: "x", Position: 6},
		{Type: calculator.OperatorToken, Value: "+", Position: 8},
	}

	result, err := calculator.EvaluatePostfix(tokens)
	if err != nil {
		t.Fatalf("EvaluatePostfix failed: %v", err)
	}

	expected := big.NewRat(14, 1)
	if result.Cmp(expected) != 0 {
		t.Errorf("2 + 3 x 4 = %s, want %s", result, expected)
	}
}

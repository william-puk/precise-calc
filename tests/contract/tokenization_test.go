package contract

import (
	"precise-calc/pkg/calculator"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []calculator.Token
	}{
		{
			"5 + 3",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "5", Position: 0},
				{Type: calculator.OperatorToken, Value: "+", Position: 2},
				{Type: calculator.NumberToken, Value: "3", Position: 4},
			},
		},
		{
			"0xFF - 0xAB",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "0xFF", Position: 0},
				{Type: calculator.OperatorToken, Value: "-", Position: 5},
				{Type: calculator.NumberToken, Value: "0xAB", Position: 7},
			},
		},
		{
			"  2  x  3  ",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "2", Position: 2},
				{Type: calculator.OperatorToken, Value: "x", Position: 5},
				{Type: calculator.NumberToken, Value: "3", Position: 8},
			},
		},
		{
			"15/3",
			[]calculator.Token{
				{Type: calculator.NumberToken, Value: "15", Position: 0},
				{Type: calculator.OperatorToken, Value: "/", Position: 2},
				{Type: calculator.NumberToken, Value: "3", Position: 3},
			},
		},
	}

	for _, test := range tests {
		result, err := calculator.Tokenize(test.input)
		if err != nil {
			t.Errorf("Tokenize(%s) error: %v", test.input, err)
			continue
		}

		if len(result) != len(test.expected) {
			t.Errorf("Tokenize(%s) returned %d tokens, want %d",
				test.input, len(result), len(test.expected))
			continue
		}

		for i, token := range result {
			expected := test.expected[i]
			if token.Type != expected.Type || token.Value != expected.Value || token.Position != expected.Position {
				t.Errorf("Token %d: got {%v, %s, %d}, want {%v, %s, %d}",
					i, token.Type, token.Value, token.Position, expected.Type, expected.Value, expected.Position)
			}
		}
	}
}

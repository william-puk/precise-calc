package calculator

import (
	"math/big"
)

// EvaluatePostfix evaluates postfix expression to get final result
func EvaluatePostfix(tokens []Token) (*big.Rat, error) {
	stack := []*big.Rat{}

	for _, token := range tokens {
		switch token.Type {
		case NumberToken:
			// Parse the number based on its format
			var num *big.Rat
			var err error

			if len(token.Value) >= 2 && token.Value[:2] == "0x" {
				num, err = ParseHexadecimal(token.Value)
			} else if len(token.Value) >= 3 && token.Value[:3] == "-0x" {
				num, err = ParseHexadecimal(token.Value)
			} else {
				num, err = ParseDecimal(token.Value)
			}

			if err != nil {
				return nil, ParseError{Message: "Invalid number format: " + token.Value, Position: token.Position}
			}

			stack = append(stack, num)

		case OperatorToken:
			if len(stack) < 2 {
				return nil, ParseError{Message: "Insufficient operands for operator", Position: token.Position}
			}

			// Pop two operands
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Perform operation
			result, err := performOperation(left, right, rune(token.Value[0]), token.Position)
			if err != nil {
				return nil, err
			}

			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return nil, ParseError{Message: "Invalid expression structure", Position: 0}
	}

	return stack[0], nil
}

// performOperation performs a single arithmetic operation
func performOperation(left, right *big.Rat, operator rune, position int) (*big.Rat, error) {
	result := new(big.Rat)

	switch operator {
	case '+':
		result.Add(left, right)
	case '-':
		result.Sub(left, right)
	case 'x':
		result.Mul(left, right)
	case '/':
		// Check for division by zero
		if right.Sign() == 0 {
			return nil, DivisionByZeroError{Position: position}
		}
		result.Quo(left, right)
	default:
		return nil, ParseError{Message: "Unknown operator: " + string(operator), Position: position}
	}

	return result, nil
}

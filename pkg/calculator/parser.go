package calculator

// ParseExpression converts tokens into validated expression structure
func ParseExpression(tokens []Token) (*Expression, error) {
	if len(tokens) == 0 {
		return nil, EmptyExpressionError{}
	}

	// Basic validation: must start and end with numbers
	if tokens[0].Type != NumberToken {
		return nil, ParseError{Message: "Expression must start with a number", Position: tokens[0].Position}
	}
	if tokens[len(tokens)-1].Type != NumberToken {
		return nil, ParseError{Message: "Expression must end with a number", Position: tokens[len(tokens)-1].Position}
	}

	// Validate alternating pattern: number op number op number...
	for i, token := range tokens {
		if i%2 == 0 { // Even positions should be numbers
			if token.Type != NumberToken {
				return nil, ParseError{Message: "Expected number", Position: token.Position}
			}
		} else { // Odd positions should be operators
			if token.Type != OperatorToken {
				return nil, ParseError{Message: "Expected operator", Position: token.Position}
			}
		}
	}

	// Convert to postfix notation for evaluation
	postfixTokens, err := InfixToPostfix(tokens)
	if err != nil {
		return nil, err
	}

	return &Expression{
		Tokens:        tokens,
		PostfixTokens: postfixTokens,
		Result:        nil,
	}, nil
}

// InfixToPostfix converts infix notation to postfix using Shunting Yard algorithm
func InfixToPostfix(tokens []Token) ([]Token, error) {
	output := []Token{}
	operatorStack := []Token{}

	for _, token := range tokens {
		switch token.Type {
		case NumberToken:
			output = append(output, token)

		case OperatorToken:
			op := OperatorMap[rune(token.Value[0])]

			// Pop operators with higher or equal precedence
			for len(operatorStack) > 0 {
				stackTop := operatorStack[len(operatorStack)-1]
				stackOp := OperatorMap[rune(stackTop.Value[0])]

				if stackOp.Precedence >= op.Precedence {
					output = append(output, stackTop)
					operatorStack = operatorStack[:len(operatorStack)-1]
				} else {
					break
				}
			}

			operatorStack = append(operatorStack, token)
		}
	}

	// Pop remaining operators
	for len(operatorStack) > 0 {
		output = append(output, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return output, nil
}

// ValidateExpression validates expression format without performing calculation
func ValidateExpression(expression string) error {
	tokens, err := Tokenize(expression)
	if err != nil {
		return err
	}

	_, err = ParseExpression(tokens)
	return err
}

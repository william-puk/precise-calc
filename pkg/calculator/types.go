package calculator

import "math/big"

// NumberType represents the format of a parsed number
type NumberType int

const (
	Decimal NumberType = iota
	Hexadecimal
)

// TokenType represents the type of a parsed token
type TokenType int

const (
	NumberToken TokenType = iota
	OperatorToken
	WhitespaceToken
)

// Associativity represents operator associativity
type Associativity int

const (
	Left Associativity = iota
	Right
)

// Number represents a parsed numeric value
type Number struct {
	Value    *big.Rat
	Original string
	Type     NumberType
}

// Operator represents a mathematical operation
type Operator struct {
	Symbol        rune
	Precedence    int
	Associativity Associativity
}

// Token represents a parsed element from input
type Token struct {
	Type     TokenType
	Value    string
	Position int
}

// Expression represents a complete mathematical expression
type Expression struct {
	Original      string
	Tokens        []Token
	PostfixTokens []Token
	Result        *big.Rat
}

// Predefined operators with precedence
var (
	AdditionOp       = Operator{Symbol: '+', Precedence: 1, Associativity: Left}
	SubtractionOp    = Operator{Symbol: '-', Precedence: 1, Associativity: Left}
	MultiplicationOp = Operator{Symbol: 'x', Precedence: 2, Associativity: Left}
	DivisionOp       = Operator{Symbol: '/', Precedence: 2, Associativity: Left}
)

// OperatorMap maps operator symbols to their definitions
var OperatorMap = map[rune]Operator{
	'+': AdditionOp,
	'-': SubtractionOp,
	'x': MultiplicationOp,
	'/': DivisionOp,
}

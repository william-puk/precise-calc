# Data Model: Precise Decimal and Hex Calculator

## Core Entities

### Number
**Purpose**: Represents a parsed numeric value in the calculation system

**Fields**:
- `Value`: `*big.Rat` - The exact numeric value using arbitrary precision
- `Original`: `string` - The original string representation as provided by user
- `Type`: `NumberType` - Enum indicating decimal or hexadecimal origin

**Validation Rules**:
- Original must contain only characters from set [A-Fa-f0-9x.-]
- For hex numbers: must start with "0x" or "-0x"
- For decimal numbers: must be valid decimal format (including scientific notation support)
- Value must not be nil after successful parsing

**State Transitions**:
1. Raw string → Tokenized → Parsed → Validated → Ready for calculation

### Operator
**Purpose**: Represents a mathematical operation in the expression

**Fields**:
- `Symbol`: `rune` - The operator character (+, -, x, /)
- `Precedence`: `int` - Numeric precedence for order of operations
- `Associativity`: `Associativity` - Left or right associativity

**Validation Rules**:
- Symbol must be one of: '+', '-', 'x', '/'
- Precedence: Addition/Subtraction = 1, Multiplication/Division = 2
- All operators are left-associative

**Predefined Values**:
```
Addition:       {Symbol: '+', Precedence: 1, Associativity: Left}
Subtraction:    {Symbol: '-', Precedence: 1, Associativity: Left}
Multiplication: {Symbol: 'x', Precedence: 2, Associativity: Left}
Division:       {Symbol: '/', Precedence: 2, Associativity: Left}
```

### Token
**Purpose**: Represents a parsed element from the input expression

**Fields**:
- `Type`: `TokenType` - Enum: Number, Operator, Whitespace
- `Value`: `string` - The raw token value
- `Position`: `int` - Character position in original input (for error reporting)

**Validation Rules**:
- Type must correspond to Value content
- Position must be >= 0 and within input string bounds
- Value must not be empty for Number and Operator types

**State Transitions**:
1. Character stream → Tokenized → Classified → Ready for parsing

### Expression
**Purpose**: Represents a complete mathematical expression to be evaluated

**Fields**:
- `Original`: `string` - The original input string
- `Tokens`: `[]Token` - Parsed tokens in order
- `PostfixTokens`: `[]Token` - Tokens arranged in postfix notation for evaluation
- `Result`: `*big.Rat` - The calculated result (nil until evaluated)

**Validation Rules**:
- Original must not be empty
- Tokens must represent a valid mathematical expression
- PostfixTokens must be generated from valid infix expression
- Result must be computable (no division by zero)

**State Transitions**:
1. Raw input → Tokenized → Validated → Converted to postfix → Evaluated → Complete

### CalculationError
**Purpose**: Represents specific error conditions during calculation

**Fields**:
- `Type`: `ErrorType` - Enum: ParseError, DivisionByZero, InvalidCharacter, etc.
- `Message`: `string` - Human-readable error description
- `Position`: `int` - Character position where error occurred (if applicable)
- `Context`: `string` - Additional context for debugging

**Error Types**:
- `ParseError`: Invalid number format
- `DivisionByZero`: Attempt to divide by zero
- `InvalidCharacter`: Character outside allowed set
- `InvalidExpression`: Malformed mathematical expression
- `EmptyInput`: No input provided

## Type Definitions

### Enums
```go
type NumberType int
const (
    Decimal NumberType = iota
    Hexadecimal
)

type TokenType int
const (
    NumberToken TokenType = iota
    OperatorToken
    WhitespaceToken
)

type Associativity int
const (
    Left Associativity = iota
    Right
)

type ErrorType int
const (
    ParseError ErrorType = iota
    DivisionByZero
    InvalidCharacter
    InvalidExpression
    EmptyInput
)
```

## Relationships

### Number ↔ Token
- One-to-one: Each Number token contains one Number entity
- Number tokens are created during tokenization phase

### Operator ↔ Token
- One-to-one: Each Operator token contains one Operator entity
- Operator tokens are created during tokenization phase

### Expression ↔ Token
- One-to-many: Expression contains multiple Token entities
- Expression manages token lifecycle and ordering

### Expression ↔ Number
- One-to-many: Expression can contain multiple Number entities
- Numbers are extracted from Number tokens during evaluation

### CalculationError ↔ Expression
- Many-to-one: Multiple error types can occur during expression processing
- Errors provide context about specific Expression that failed

## Data Flow

```
Input String
    ↓
Tokenization → []Token
    ↓
Validation → Valid Expression
    ↓
Infix to Postfix → []Token (postfix order)
    ↓
Evaluation → *big.Rat (result)
    ↓
Output Formatting → String (result)
```

## Validation Rules Summary

### Input Level
- Must contain only characters: [A-Fa-f0-9x+\-\s\t\n/]
- Must not be empty
- Must represent valid mathematical expression

### Token Level
- Number tokens must parse to valid decimal or hex format
- Operator tokens must be valid mathematical operators
- Positions must be accurate for error reporting

### Expression Level
- Must have balanced operators and operands
- Must not have division by zero
- Must be evaluable to a finite result

### Output Level
- Result must be representable as string
- Error messages must be clear and actionable
- Position information must be accurate for debugging
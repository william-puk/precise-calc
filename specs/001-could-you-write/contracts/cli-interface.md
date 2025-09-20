# CLI Interface Contract

## Overview
Command-line interface for the precise decimal and hex calculator. Accepts mathematical expressions as input and returns calculated results with arbitrary precision.

## Command Specification

### Basic Usage
```bash
precise-calc "<expression>"
```

### Input Format
- **Type**: String argument containing mathematical expression
- **Character Set**: [A-Fa-f0-9x+\-\s\t\n/]
- **Operators**: `+` (addition), `-` (subtraction), `x` (multiplication), `/` (division)
- **Numbers**:
  - Decimal: Standard format including scientific notation (e.g., `123.456`, `-0.001`, `1e-10`)
  - Hexadecimal: Prefixed with `0x` or `-0x` (e.g., `0xAB91`, `-0xFF`)
- **Whitespace**: Ignored (spaces, tabs, newlines)

### Output Format

#### Success Response
```
<calculated_result>
```
- **Format**: Decimal representation of the result
- **Precision**: Exact rational representation (no rounding errors)
- **Exit Code**: 0

#### Error Response
```
Error: <error_message>
Position: <character_position> (if applicable)
```
- **Format**: Human-readable error description
- **Exit Code**: 1

## Input Examples

### Valid Expressions
```bash
# Basic decimal arithmetic
precise-calc "0.1 + 0.2"
# Output: 0.3

# Hexadecimal arithmetic
precise-calc "0xAB91 + 100"
# Output: 43921

# Mixed decimal and hex with precedence
precise-calc "0.0000000000000001 + 0.1 + -99999999999999 - 0xab91"
# Output: -100000000043890.8999999999999999

# Whitespace handling
precise-calc "  2  x  3  /  6  "
# Output: 1

# Operator precedence
precise-calc "2 + 3 x 4"
# Output: 14

# Negative numbers
precise-calc "-5 + 3"
# Output: -2

# Negative hex numbers
precise-calc "-0xFF + 256"
# Output: 1
```

### Error Cases
```bash
# Division by zero
precise-calc "5 / 0"
# Output: Error: Division by zero
# Exit Code: 1

# Invalid characters
precise-calc "5 + G"
# Output: Error: Invalid character 'G' at position 4
# Exit Code: 1

# Malformed hex number
precise-calc "0xGHI + 5"
# Output: Error: Invalid hexadecimal format at position 0
# Exit Code: 1

# Empty input
precise-calc ""
# Output: Error: Empty expression provided
# Exit Code: 1

# Invalid expression structure
precise-calc "5 + + 3"
# Output: Error: Invalid expression structure at position 4
# Exit Code: 1
```

## Contract Tests

### Functional Requirements Mapping

#### FR-001: Accept mathematical expression as string input
```bash
Test: precise-calc "5 + 3"
Expected: "8"
Exit Code: 0
```

#### FR-002: Parse decimal numbers with arbitrary precision
```bash
Test: precise-calc "0.0000000000000001 + 0.1"
Expected: "0.1000000000000001"
Exit Code: 0
```

#### FR-003: Parse hexadecimal numbers prefixed with "0x" or "-0x"
```bash
Test: precise-calc "0xAB + 0xCD"
Expected: "376"
Exit Code: 0
```

#### FR-004: Support negative numbers for both formats
```bash
Test: precise-calc "-5 + -0xFF"
Expected: "-260"
Exit Code: 0
```

#### FR-005: Support four basic mathematical operators
```bash
Test: precise-calc "2 + 3 - 1 x 4 / 2"
Expected: "3"
Exit Code: 0
```

#### FR-006: Follow standard mathematical operator precedence
```bash
Test: precise-calc "2 + 3 x 4"
Expected: "14"
Exit Code: 0
```

#### FR-007: Ignore whitespace characters
```bash
Test: precise-calc "  2  +  3  "
Expected: "5"
Exit Code: 0
```

#### FR-008: Preserve precision throughout calculations
```bash
Test: precise-calc "1/3 + 1/3 + 1/3"
Expected: "1"
Exit Code: 0
```

#### FR-009: Validate character set [A-Fa-f0-9x]
```bash
Test: precise-calc "5 + Z"
Expected: Error message
Exit Code: 1
```

#### FR-010: Output final calculated result
```bash
Test: precise-calc "7 x 6"
Expected: "42"
Exit Code: 0
```

#### FR-011: Handle division by zero
```bash
Test: precise-calc "1 / 0"
Expected: Error message containing "division by zero"
Exit Code: 1
```

#### FR-012: Reject invalid characters
```bash
Test: precise-calc "5 + @"
Expected: Error message containing "invalid character"
Exit Code: 1
```

## Performance Expectations

### Response Time
- **Simple expressions** (< 10 tokens): < 1ms
- **Complex expressions** (10-100 tokens): < 10ms
- **Large numbers** (> 1000 digits): < 100ms

### Memory Usage
- **Baseline**: < 1MB for simple calculations
- **Large numbers**: Scales with precision requirements
- **No memory leaks**: Proper cleanup of big.Rat allocations

### Resource Limits
- **Expression length**: Maximum 10,000 characters
- **Number precision**: Limited only by available memory
- **Recursion depth**: Maximum 1000 levels (for complex precedence)

## Security Considerations

### Input Validation
- Character set strictly enforced
- No code execution capabilities
- No file system access
- No network operations

### Resource Protection
- Expression complexity limits prevent DoS
- Memory allocation monitoring
- Execution time limits for very large calculations

## Error Handling Specification

### Error Categories
1. **Input Errors**: Invalid format, empty input, forbidden characters
2. **Parse Errors**: Malformed expressions, invalid number formats
3. **Calculation Errors**: Division by zero, overflow conditions
4. **System Errors**: Memory allocation failures, internal errors

### Error Message Format
```
Error: <category>: <specific_description>
Position: <character_position> (when applicable)
Context: <surrounding_text> (when helpful)
```

### Recovery Behavior
- **No recovery**: Each invocation is independent
- **Clean termination**: Always exit gracefully with appropriate code
- **Resource cleanup**: Free all allocated memory before exit
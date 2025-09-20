# Tasks: Precise Decimal and Hex Calculator

**Input**: Design documents from `/specs/001-could-you-write/`
**Prerequisites**: plan.md (required), research.md, data-model.md, contracts/

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → Extract: Go 1.25+, math/big, TDD approach, library-first
2. Load design documents:
   → data-model.md: Number, Operator, Token, Expression, CalculationError entities
   → contracts/: CLI interface and library API specifications
   → research.md: Recursive descent parser, state-machine tokenizer
3. Generate tasks by category:
   → Setup: Go module, directory structure, linting
   → Tests: Contract tests for library API and CLI interface
   → Core: Number parsing, tokenization, expression evaluation
   → Integration: CLI wrapper, error handling
   → Polish: Unit tests, performance validation, documentation
4. Apply TDD rules:
   → All tests written first and must fail
   → Implementation follows to make tests pass
   → Different files = mark [P] for parallel execution
5. Number tasks sequentially (T001, T002...)
6. SUCCESS: Ready for implementation
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Path Conventions
Go project structure from plan.md:
- `cmd/precise-calc/`: CLI application entry point
- `pkg/calculator/`: Core calculator library
- `tests/contract/`: Contract tests for API compliance
- `tests/integration/`: End-to-end tests
- `tests/unit/`: Unit tests for components

## Phase 3.1: Setup
- [x] T001 Initialize Go module and create directory structure per implementation plan
- [x] T002 [P] Create go.mod with Go 1.25+ and standard library dependencies only
- [x] T003 [P] Configure Go formatting tools (gofmt, go vet) and basic Makefile

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3
**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**

### Contract Tests for Library API
- [x] T004 [P] Contract test ParseDecimal function in tests/contract/number_parsing_test.go
- [x] T005 [P] Contract test ParseHexadecimal function in tests/contract/number_parsing_test.go
- [x] T006 [P] Contract test Tokenize function in tests/contract/tokenization_test.go
- [x] T007 [P] Contract test Calculate function basic arithmetic in tests/contract/calculator_test.go
- [x] T008 [P] Contract test Calculate function error handling in tests/contract/calculator_test.go

### CLI Interface Contract Tests
- [x] T009 [P] CLI contract test basic usage and output format in tests/contract/cli_test.go
- [x] T010 [P] CLI contract test error cases and exit codes in tests/contract/cli_test.go

### Integration Tests
- [x] T011 [P] Integration test end-to-end calculation scenarios in tests/integration/calculator_test.go
- [x] T012 [P] Integration test complex expressions with operator precedence in tests/integration/calculator_test.go

## Phase 3.3: Core Implementation (ONLY after tests are failing)

### Data Models and Types
- [x] T013 [P] Implement Number, Operator, Token type definitions in pkg/calculator/types.go
- [x] T014 [P] Implement CalculationError types and error handling in pkg/calculator/errors.go

### Number Parsing
- [x] T015 [P] Implement ParseDecimal function with math/big.Rat in pkg/calculator/number_parser.go
- [x] T016 [P] Implement ParseHexadecimal function with math/big.Rat in pkg/calculator/number_parser.go

### Expression Processing
- [x] T017 Implement Tokenize function with state-machine approach in pkg/calculator/tokenizer.go
- [x] T018 Implement expression validation and parsing in pkg/calculator/parser.go
- [x] T019 Implement operator precedence evaluation (recursive descent) in pkg/calculator/evaluator.go

### Core Calculator Function
- [x] T020 Implement main Calculate function orchestrating parsing and evaluation in pkg/calculator/calculator.go
- [x] T021 Implement ValidateExpression and FormatResult helper functions in pkg/calculator/calculator.go

### CLI Application
- [x] T022 Implement CLI argument parsing and main function in cmd/precise-calc/main.go
- [x] T023 Implement CLI output formatting and error handling in cmd/precise-calc/main.go

## Phase 3.4: Integration
- [x] T024 Connect CLI to calculator library with proper error propagation
- [x] T025 Implement comprehensive input validation and sanitization
- [x] T026 Add logging and debugging support for complex expressions
- [x] T027 Implement proper exit codes and error message formatting

## Phase 3.5: Polish
- [x] T028 [P] Unit tests for edge cases in number parsing in tests/unit/number_parser_test.go
- [x] T029 [P] Unit tests for tokenizer edge cases in tests/unit/tokenizer_test.go
- [x] T030 [P] Unit tests for evaluator edge cases in tests/unit/evaluator_test.go
- [x] T031 Performance tests for large numbers and complex expressions
- [x] T032 [P] Update README.md with usage examples and build instructions
- [x] T033 Code review and refactoring for clarity and maintainability
- [x] T034 Run quickstart.md validation scenarios and fix any issues

## Dependencies
```
Setup (T001-T003) → Tests (T004-T012) → Implementation (T013-T023) → Integration (T024-T027) → Polish (T028-T034)

Specific dependencies:
- T013-T014 (types/errors) before T015-T021 (implementations)
- T015-T016 (number parsing) before T017-T019 (expression processing)
- T017-T019 (parsing/evaluation) before T020-T021 (calculator functions)
- T020-T021 (calculator) before T022-T023 (CLI)
- All core implementation before integration (T024-T027)
```

## Parallel Example
```bash
# Launch contract tests together after setup:
Task: "Contract test ParseDecimal function in tests/contract/number_parsing_test.go"
Task: "Contract test ParseHexadecimal function in tests/contract/number_parsing_test.go"
Task: "Contract test Tokenize function in tests/contract/tokenization_test.go"
Task: "CLI contract test basic usage in tests/contract/cli_test.go"

# Launch type definitions in parallel:
Task: "Implement Number, Operator, Token types in pkg/calculator/types.go"
Task: "Implement CalculationError types in pkg/calculator/errors.go"

# Launch unit tests in parallel during polish phase:
Task: "Unit tests for number parsing edge cases in tests/unit/number_parser_test.go"
Task: "Unit tests for tokenizer edge cases in tests/unit/tokenizer_test.go"
Task: "Unit tests for evaluator edge cases in tests/unit/evaluator_test.go"
```

## Task Details

### Key Implementation Requirements
1. **Arbitrary Precision**: Use math/big.Rat for all calculations
2. **Input Validation**: Character set [A-Fa-f0-9x+\-\s\t\n/] only
3. **Operator Precedence**: Multiplication/division before addition/subtraction
4. **Error Handling**: Specific error types for different failure modes
5. **CLI Interface**: Expression string input, result string output, proper exit codes

### Test Requirements
1. **Contract Tests**: Verify API compliance with specifications
2. **Integration Tests**: End-to-end scenarios from user stories
3. **Unit Tests**: Edge cases and error conditions
4. **Performance Tests**: <1ms simple expressions, <100ms complex calculations

### Validation Scenarios
From quickstart.md and contracts:
- "0.1 + 0.2" → "0.3" (exact precision)
- "0xFF + 1" → "256" (hex conversion)
- "2 + 3 x 4" → "14" (operator precedence)
- "0.0000000000000001 + 0.1 + -99999999999999 - 0xab91" → exact result
- Division by zero → proper error
- Invalid characters → specific error with position

## Notes
- [P] tasks operate on different files with no dependencies
- Verify all tests fail before implementing corresponding functionality
- Follow TDD cycle: Red (failing test) → Green (minimal implementation) → Refactor
- Commit after each completed task
- Run `go test ./...` frequently to ensure tests pass
- Use `go fmt` and `go vet` before committing

## Validation Checklist
*GATE: Verified before task execution*

- [x] All contracts have corresponding tests (T004-T012)
- [x] All entities have implementation tasks (T013-T021)
- [x] All tests come before implementation (T004-T012 before T013-T023)
- [x] Parallel tasks operate on different files (marked with [P])
- [x] Each task specifies exact file path
- [x] No task modifies same file as another [P] task
- [x] TDD approach enforced throughout
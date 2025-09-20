# precise-calc Development Guidelines

Auto-generated from all feature plans. Last updated: 2025-09-21

## Active Technologies
- Go 1.19+ with math/big package for arbitrary precision arithmetic (001-could-you-write)
- Standard library only (no external dependencies) (001-could-you-write)

## Project Structure
```
cmd/precise-calc/          # CLI application entry point
pkg/calculator/            # Core calculator library
tests/
├── contract/             # Contract tests for API compliance
├── integration/          # End-to-end tests
└── unit/                # Unit tests for components
specs/001-could-you-write/ # Feature specifications and plans
```

## Commands
# Build and run calculator
go build -o precise-calc cmd/precise-calc/main.go
./precise-calc "0.1 + 0.2"

# Run tests
go test ./...
go test -cover ./...

## Code Style
- Use TDD: Tests first, then implementation
- Library-first: Core logic in pkg/calculator, CLI wrapper in cmd/
- Exact precision: Use math/big.Rat for all calculations
- Clear error handling: Specific error types for different failures

## Recent Changes
- 001-could-you-write: Added Go calculator with decimal/hex support and arbitrary precision

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
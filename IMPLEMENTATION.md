# Chorlang Reference Implementation

This is the reference implementation of the Chorlang programming language, built in Go.

## Project Structure

```
Chorlang/
├── compiler/          # Compiler implementation
│   ├── lexer/        # Tokenization
│   ├── parser/       # AST generation  
│   ├── ast/          # Abstract syntax tree definitions
│   └── codegen/      # Go code generation
├── cmd/
│   └── chorelang/    # CLI tool
├── examples/         # Example Chorlang programs
└── tests/           # Test suite
```

## Building

```bash
make build
```

## Usage

```bash
# Generate Go code
./chorelang hello.chore

# Compile to binary
./chorelang -c hello.chore

# Run immediately
./chorelang -r hello.chore

# Compile with custom output name
./chorelang -c -o myapp hello.chore
```

## Implementation Status

### Completed Features

- **Lexer**: Full tokenization of Chorlang syntax including dance-inspired keywords
- **Parser**: Recursive descent parser building complete AST
- **Code Generator**: Transpiles to Go code with proper scoping
- **Core Language Features**:
  - Variable declarations (`dance`)
  - For loops (`sway`)
  - Function calls (`spin`)
  - Conditionals (`if/else`)
  - Basic types (int, float, string, bool)
  - Pattern matching (`match/when`)
  - Goroutines (`start`)
  - Channels (`flow`, `send`)

### Example Programs

See the `examples/` directory for sample Chorlang programs:
- `hello_world.chore` - Basic output
- `fibonacci.chore` - Loop and variable usage
- `conditions.chore` - Conditionals and pattern matching
- `concurrent.chore` - Goroutines and channels

## Testing

Run all tests:
```bash
make test
```

Run examples:
```bash
make run-example-hello
make run-example-fibonacci
```

## Next Steps

Future enhancements could include:
- Function definitions
- Type system and type checking
- Import system
- Standard library implementation
- Error handling improvements
- Optimization passes
- Dance diagram generation
- REPL implementation
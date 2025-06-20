# Chorlang User Manual

Welcome to Chorlang - a graceful programming language that brings the elegance of dance to code!

## Table of Contents
1. [Quickstart Guide](#quickstart-guide)
2. [Language Basics](#language-basics)
3. [Control Flow](#control-flow)
4. [Concurrency](#concurrency)
5. [Pattern Matching](#pattern-matching)
6. [Compiler Usage](#compiler-usage)
7. [Troubleshooting](#troubleshooting)

## Quickstart Guide

### Installation

1. **Prerequisites**: Ensure you have Go 1.21+ installed on your system.

2. **Clone and Build**:
```bash
git clone https://github.com/yourusername/chorlang.git
cd chorlang
make build
```

3. **Verify Installation**:
```bash
./chorelang -h
```

### Your First Chorlang Program

Create a file named `hello.chore`:

```chorelang
// My first Chorlang program
spin print("Hello, Chorlang!")
```

Run it:
```bash
./chorelang -r hello.chore
```

### Quick Examples

**Variables and Math**:
```chorelang
dance x = 10
dance y = 20
dance sum = x + y
spin print("Sum is:", sum)
```

**Loops**:
```chorelang
sway i from 1 to 5 {
    spin print("Step", i)
}
```

**Concurrent Hello**:
```chorelang
start spin print("Hello from a goroutine!")
spin print("Hello from main!")
```

## Language Basics

### Variables - `dance`

Declare variables with the `dance` keyword:

```chorelang
dance name = "Chorlang"
dance version = 1.0
dance isAwesome = true
```

Variables can be reassigned:
```chorelang
dance count = 0
count = count + 1  // No 'dance' needed for reassignment
```

### Data Types

Chorlang supports these basic types:
- **Integers**: `dance age = 25`
- **Floats**: `dance pi = 3.14159`
- **Strings**: `dance msg = "Hello, dancer!"`
- **Booleans**: `dance ready = true`

### Function Calls - `spin`

Call functions with the `spin` keyword:

```chorelang
spin print("Hello")
spin print("Value:", 42)
spin print("Multiple", "arguments", "work!")
```

### Comments

```chorelang
// This is a single-line comment
spin print("Code here")  // Comments can go after code too
```

## Control Flow

### Loops - `sway`

The `sway` keyword creates loops that iterate from a start to end value:

```chorelang
// Count from 0 to 10 (inclusive)
sway i from 0 to 10 {
    spin print(i)
}

// Use variables for bounds
dance start = 5
dance end = 15
sway j from start to end {
    spin print("Number:", j)
}
```

### Conditionals - `if/else`

Standard conditional statements:

```chorelang
dance age = 18

if age >= 18 {
    spin print("You can vote!")
} else {
    spin print("Too young to vote")
}

// Multiple conditions
dance score = 85
if score >= 90 {
    spin print("A grade")
} else if score >= 80 {
    spin print("B grade")
} else {
    spin print("Keep practicing!")
}
```

### Comparison Operators

- `==` - Equal to
- `!=` - Not equal to
- `<` - Less than
- `>` - Greater than
- `<=` - Less than or equal
- `>=` - Greater than or equal
- `=~` - Pattern match (for regex)

## Concurrency

### Goroutines - `start`

Launch concurrent tasks with `start`:

```chorelang
// Start a concurrent task
start {
    sway i from 0 to 5 {
        spin print("Background:", i)
    }
}

// Start multiple goroutines
sway i from 0 to 3 {
    start spin print("Dancer", i, "is performing!")
}
```

### Channels - `flow` and `send`

Channels enable communication between goroutines:

```chorelang
// Declare a channel
flow messages = flow channel<string>

// Send to channel in a goroutine
start {
    send messages <- "Hello from goroutine"
}

// Receive from channel
dance msg = <-messages
spin print("Received:", msg)
```

**Producer-Consumer Example**:
```chorelang
flow numbers = flow channel<int>

// Producer
start {
    sway i from 1 to 5 {
        send numbers <- i
    }
}

// Consumer
start {
    sway i from 1 to 5 {
        dance num = <-numbers
        spin print("Got:", num)
    }
}
```

## Pattern Matching

### Match Expressions - `match/when`

Pattern matching for elegant control flow:

```chorelang
dance userType = "admin"

dance access = match userType {
    when "admin": flow "full"
    when "user": flow "limited"
    when "guest": flow "readonly"
}

spin print("Access level:", access)
```

**With Pattern Destructuring** (planned feature):
```chorelang
dance result = match response {
    when Success(data): flow processData(data)
    when Error(msg): flow handleError(msg)
    when Loading(): flow showSpinner()
}
```

## Compiler Usage

### Basic Commands

**Generate Go Code**:
```bash
./chorelang myprogram.chore
# Creates: myprogram.go
```

**Compile to Binary**:
```bash
./chorelang -c myprogram.chore
# Creates: myprogram (executable)
```

**Run Immediately**:
```bash
./chorelang -r myprogram.chore
# Compiles and runs in one step
```

**Custom Output Name**:
```bash
./chorelang -c -o myapp myprogram.chore
# Creates: myapp (executable)
```

### Development Workflow

1. **Write** your Chorlang code in `.chore` files
2. **Test** quickly with `-r` flag
3. **Debug** by generating Go code to inspect
4. **Deploy** by compiling to binary with `-c`

### Makefile Targets

If using the provided Makefile:

```bash
make build              # Build the compiler
make test              # Run all tests
make run-example-hello # Run hello world example
make compile-examples  # Compile all examples
```

## Troubleshooting

### Common Issues

**"Parser errors" when compiling**:
- Check syntax - remember to use dance keywords
- Ensure proper block structure with `{` and `}`
- Verify string quotes are closed

**"Undefined variable" errors in generated Go**:
- Variables must be declared with `dance` before use
- Check variable scoping in loops and conditions

**Runtime panics with channels**:
- Ensure channels are created with `flow channel<type>`
- Match sends and receives to avoid deadlocks

### Best Practices

1. **Name Convention**: Use descriptive names for clarity
   ```chorelang
   dance userAge = 25  // Good
   dance x = 25        // Less clear
   ```

2. **Scope Management**: Declare variables in appropriate scope
   ```chorelang
   dance total = 0
   sway i from 1 to 10 {
       total = total + i  // Reuse outer variable
   }
   ```

3. **Concurrent Safety**: Be careful with shared variables
   ```chorelang
   flow safe = flow channel<int>
   start send safe <- computeValue()
   dance result = <-safe  // Safe communication
   ```

### Getting Help

- Check the examples in `examples/` directory
- Review generated Go code for debugging
- Consult the technical specification in `TECH_SPEC.md`

## Next Steps

Now that you've learned the basics:

1. Try modifying the example programs
2. Build a small project using concurrency
3. Experiment with pattern matching
4. Share your Chorlang creations!

Remember: In Chorlang, every program is a performance, and every function call is a dance move. Code with grace! 💃🕺
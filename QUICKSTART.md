# Chorlang Quickstart for Developers

Get dancing with code in 5 minutes! 🕺

## 1. Build the Compiler (30 seconds)

```bash
git clone https://github.com/yourusername/chorlang.git
cd chorlang
make build
```

## 2. Hello World (1 minute)

Create `hello.chore`:
```chorelang
spin print("Hello, Chorlang!")
```

Run it:
```bash
./chorelang -r hello.chore
```

## 3. Learn by Example (3 minutes)

### Variables & Loops
```chorelang
dance count = 10
sway i from 1 to count {
    spin print("Step", i)
}
```

### Conditionals
```chorelang
dance age = 21
if age >= 18 {
    spin print("Adult")
} else {
    spin print("Minor")
}
```

### Concurrency (Go-style)
```chorelang
// Launch goroutine
start spin print("I'm concurrent!")

// Channels
flow msgs = flow channel<string>
start send msgs <- "Hello"
dance received = <-msgs
```

### Pattern Matching
```chorelang
dance result = match status {
    when "success": flow "✓"
    when "error": flow "✗"
    when "pending": flow "..."
}
```

## 4. Compiler Commands

```bash
./chorelang file.chore          # Generate Go code
./chorelang -c file.chore       # Compile to binary
./chorelang -r file.chore       # Run immediately
./chorelang -c -o app file.chore # Custom output name
```

## 5. Complete Example: Concurrent Counter

Create `counter.chore`:
```chorelang
flow counts = flow channel<int>

// Producer: send numbers
start {
    sway i from 1 to 5 {
        send counts <- i
        spin print("Sent:", i)
    }
}

// Consumer: receive and sum
dance total = 0
sway i from 1 to 5 {
    dance num = <-counts
    total = total + num
    spin print("Received:", num, "Total:", total)
}

spin print("Final sum:", total)
```

Run it:
```bash
./chorelang -r counter.chore
```

## Language Cheatsheet

| Chorlang | Purpose | Go Equivalent |
|----------|---------|---------------|
| `dance` | Declare variable | `var` or `:=` |
| `sway` | For loop | `for` |
| `spin` | Function call | direct call |
| `flow` | Channel type | `chan` |
| `start` | Launch goroutine | `go` |
| `send` | Channel send | `<-` |
| `match/when` | Pattern match | `switch/case` |

## What's Next?

- 📁 Check `examples/` for more code samples
- 📖 Read `docs/USER_MANUAL.md` for detailed guide
- 🧪 Run `make test` to see tests in action
- 🚀 Build something awesome!

**Pro tip**: Generate Go code first to debug:
```bash
./chorelang mycode.chore  # Creates mycode.go
cat mycode.go            # See the generated Go code
```

Happy dancing with code! 💃
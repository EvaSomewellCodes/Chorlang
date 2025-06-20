# Chorlang Quick Reference Card

## Basic Syntax

```chorelang
// Comments
dance x = 42              // Variable declaration
x = 100                   // Reassignment (no 'dance')
spin print(x)             // Function call
```

## Data Types
```chorelang
dance num = 42            // Integer
dance pi = 3.14          // Float  
dance name = "Chorlang"  // String
dance ready = true       // Boolean
```

## Control Flow

### Loops
```chorelang
sway i from 0 to 10 {    // 0,1,2...10 (inclusive)
    spin print(i)
}
```

### Conditionals
```chorelang
if condition {
    // true branch
} else if other {
    // else if branch
} else {
    // else branch
}
```

## Operators

### Arithmetic
```chorelang
+ - * /                  // Basic math
dance sum = a + b
```

### Comparison
```chorelang
==  !=                   // Equal, not equal
<   >                    // Less, greater
<=  >=                   // Less/greater or equal
=~                       // Pattern match
```

### Logical
```chorelang
&&  ||  !                // AND, OR, NOT (planned)
```

## Concurrency

### Goroutines
```chorelang
start {                  // Launch goroutine
    spin print("Async")
}

start functionCall()     // Launch single statement
```

### Channels
```chorelang
flow ch = flow channel<int>    // Create channel
send ch <- 42                  // Send value
dance val = <-ch              // Receive value
```

## Pattern Matching
```chorelang
dance result = match expr {
    when pattern1: flow value1
    when pattern2: flow value2
    when pattern3: flow value3
}
```

## Common Patterns

### Producer-Consumer
```chorelang
flow queue = flow channel<int>

// Producer
start {
    sway i from 1 to 10 {
        send queue <- i
    }
}

// Consumer  
start {
    sway i from 1 to 10 {
        dance item = <-queue
        spin print("Processing:", item)
    }
}
```

### Error Handling (Pattern)
```chorelang
dance status = checkOperation()
dance result = match status {
    when "ok": flow processSuccess()
    when "error": flow handleError()
    when "retry": flow attemptRetry()
}
```

### Concurrent Workers
```chorelang
flow jobs = flow channel<int>
flow results = flow channel<int>

// Start workers
sway w from 1 to 3 {
    start {
        dance job = <-jobs
        send results <- job * 2
    }
}

// Send jobs
sway j from 1 to 5 {
    send jobs <- j
}

// Collect results
sway r from 1 to 5 {
    dance result = <-results
    spin print("Result:", result)
}
```

## Compiler Flags

```bash
chorelang file.chore         # Generate Go code
chorelang -r file.chore      # Run immediately  
chorelang -c file.chore      # Compile to binary
chorelang -o name file.chore # Custom output
chorelang -h                 # Show help
```

## Tips & Tricks

1. **Variable Scoping**: Variables in blocks are scoped
   ```chorelang
   dance x = 1
   if true {
       dance y = 2    // Only exists in this block
       x = 3          // Updates outer x
   }
   ```

2. **Loop Variables**: Are scoped to the loop
   ```chorelang
   sway i from 0 to 5 {
       // i only exists here
   }
   ```

3. **Debug with Go**: Generate Go to understand issues
   ```bash
   ./chorelang problem.chore
   cat problem.go  # See what went wrong
   ```

4. **Channel Deadlocks**: Match sends with receives
   ```chorelang
   // BAD: Deadlock!
   flow ch = flow channel<int>
   send ch <- 1  // Blocks forever
   
   // GOOD: Use goroutine
   flow ch = flow channel<int>
   start send ch <- 1
   dance val = <-ch
   ```

---
*Keep dancing with code!* 💃🕺
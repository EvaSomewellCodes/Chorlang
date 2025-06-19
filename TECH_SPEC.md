# ChoreLang Technical Specification

This document describes the proposed design of ChoreLang. The language draws inspiration from interpretive dance, aiming for expressive syntax and rhythmic concurrency. Implementation details may evolve, but the goals remain: clarity, elegance, and first-class AI capabilities.

## 1. Language Goals

1. Produce small, cross-platform binaries.
2. Embrace concurrency as a core construct using Go-like goroutines and channels.
3. Offer Python-like ease of use with a statically typed foundation.
4. Encourage readable code with minimal boilerplate, following Go's style.
5. Provide native access to OpenAI and Ollama APIs for AI-driven applications.

## 2. Syntax Overview

ChoreLang uses dance-inspired keywords to make code feel fluid:

```chorelang
sway i from 0 to 10 {
    spin print(i)
}
```

- `sway` acts like a `for` loop.
- `spin` calls a function.

Pattern matching and generics borrow from Scala:

```chorelang
dance result = match item {
    when Note(n): flow process_note(n)
    when Rest(): flow handle_rest()
}
```

## 3. Concurrency Model

ChoreLang compiles down to Go and retains goroutines and channels:

```chorelang
flow channel<int> steps
start sway i from 0 to 3 {
    send steps <- i
}
start sway step from steps {
    spin print("step", step)
}
```

- `start` launches a concurrent dancer (goroutine).
- `flow` declares a channel.
- `send` transmits values.

## 4. Standard Library AI Support

ChoreLang bundles clients for the OpenAI and Ollama APIs. Authentication keys can be provided through environment variables or configuration files. Sample usage:

```chorelang
flow ai = openai.client(api_key)
dance resp = ai.complete("Write a short poem about dance")
```

Calling these APIs should feel as seamless as any other standard library call.

## 5. Compiler and Tooling

The compiler is planned to be written in Go. It will translate ChoreLang source to Go code, then leverage the Go toolchain for optimized binaries. The toolchain includes:

- **chore build**: Compile sources to a native executable.
- **chore test**: Run tests.
- **chore deploy**: Package and distribute your application.

## 6. Implementation Roadmap

1. **Prototype Parser and Lexer** – Define the syntax and create a translator to Go.
2. **Concurrency Primitives** – Implement goroutine and channel mappings.
3. **AI API Modules** – Integrate OpenAI and Ollama libraries with idiomatic wrappers.
4. **Standard Library** – Provide utilities for file IO, networking, and basic data structures.
5. **IDE Support and Visual Studio** – Develop the "Studio" environment to visualize programs as choreography.

## 7. Future Ideas

- Visual debugging through dance diagrams.
- Built-in linting that encourages elegant flows.
- A package repository called "Stage" for sharing dances (libraries).

ChoreLang invites developers to craft applications as if directing a performance. By blending practical tooling with artistic inspiration, it aims to make coding a graceful and enjoyable experience.


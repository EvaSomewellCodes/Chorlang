# ChoreLang Technical Specification

This document describes the proposed design of ChoreLang. The language draws inspiration from interpretive dance, aiming for expressive syntax and rhythmic concurrency. Implementation details may evolve, but the goals remain: clarity, elegance, and first-class AI capabilities.

## 1. Language Goals

1. Produce small, cross-platform binaries.
2. Embrace concurrency as a core construct using Go-like goroutines and channels.
3. Offer Python-like ease of use with a statically typed foundation.
4. Encourage readable code with minimal boilerplate, following Go's style.
5. Provide native access to OpenAI and Ollama APIs for AI-driven applications.
6. Include a unified file system API for cross-platform development and secure sandboxing.
7. Fail fast with clear diagnostics so developers and agents quickly see where problems occur.
8. Ship a native key-value store reminiscent of Lua tables for flexible, lightweight data structures.
9. Provide an advanced regular expression and pattern matching engine in the standard library.

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

## 5. File System and API Discovery

ChoreLang includes a cross-platform file system module that abstracts away OS differences. By default, file access is sandboxed so agents can operate safely. The same module provides a discovery mechanism that scans available APIs (both local services and network endpoints) and generates bindings automatically.

```chorelang
flow fs = chore.fs()
dance content = fs.read("poem.txt")
```

Discovered APIs can be imported using a single command:

```shell
$ chore discover http://localhost:3000
```

The compiler will generate stubs so you can call these APIs just like any other library.

## 6. Key-Value Store and Pattern Matching

ChoreLang's standard library includes a native key-value store inspired by Lua tables. These tables can hold mixed types and are optimized for performance so small data sets can be stored without external dependencies.

Powerful regular expression handling is integrated directly into the language runtime. Patterns compile at build time when possible, and matching functions provide expressive captures and replacements.

```chorelang
dance t = chore.table()
t["name"] = "step"
if t["name"] =~ /st.+/ {
    spin print("matched")
}
```

The goal is to make common data manipulation and searching simple and fast without sacrificing elegance.
## 7. Dance Diagrams

ChoreLang treats every program as a choreography that can be visualized. The CLI
provides a `chore chart` command (or `file.dance --chart`) that outputs a Mermai
d diagram outlining each step of execution. These diagrams help developers and a
gents quickly grasp how data flows through the program.

## 8. Compiler and Tooling

The compiler is planned to be written in Go. It will translate ChoreLang source to Go code, then leverage the Go toolchain for optimized binaries. Compilation errors are reported with full file and line information so problems can be corrected quickly. The toolchain includes:

- **chore build**: Compile sources to a native executable.
- **chore test**: Run tests.
- **chore deploy**: Package and distribute your application.
- **chore chart**: Generate Mermaid diagrams from `.dance` files to visualize execution.
- **chore lint**: Enforce elegant style and highlight potential missteps.
- **chore stage**: Fetch libraries and plugins from the shared Stage repository.

## 9. Implementation Roadmap

1. **Prototype Parser and Lexer** – Define the syntax and create a translator to Go.
2. **Concurrency Primitives** – Implement goroutine and channel mappings.
3. **AI API Modules** – Integrate OpenAI and Ollama libraries with idiomatic wrappers.
4. **Standard Library** – Provide utilities for file IO, networking, and basic data structures.
5. **IDE Support and Visual Studio** – Develop the "Studio" environment to visualize programs as choreography.

## 10. Future Ideas

- Live-coded choreography for interactive performances.
- Advanced visualization tools that render real-time dance animations.
- Community-driven extensions distributed through the Stage repository.

ChoreLang invites developers to craft applications as if directing a performance. By blending practical tooling with artistic inspiration, it aims to make coding a graceful and enjoyable experience.


# ChoreLang

ChoreLang is a proposal for a new compiled programming language inspired by the elegance of interpretive dance. Programs in ChoreLang aim to be clear, expressive performances that transform data as gracefully as dancers move to music.

## Highlights

* **Native Compilation** - Generates small binaries that run on multiple targets with minimal runtime overhead.
* **Concurrent by Design** - Leverages Go's concurrency primitives for effortless parallelism.
* **Python-like Flexibility & Scala-like Precision** - Provides dynamic-style constructs with a robust type system and pattern matching.
* **Go-style Clarity** - Keeps code readable and the toolchain lightweight.
* **First-Class AI APIs** - OpenAI and Ollama integrations are part of the standard library, allowing AI calls without extra packages.
* **Loud, Fast Failures** - Errors surface immediately with detailed diagnostics so both humans and agents know exactly what went wrong.
* **Unified File System API** - Cross-platform file operations are built in, including secure sandboxing for agent-oriented programs.
* **Native Key-Value Store** - Table-like data structures reminiscent of Lua tables come built in for lightweight storage.
* **Fast API Discovery** - Tooling can auto-discover available APIs (local or remote) and generate stubs for quick prototyping.
* **Advanced Regex Engine** - Pattern matching with unreasonably powerful regular expressions is part of the standard library.
* **Rich Standard Library** - Includes JSON parsing, HTTP utilities, and process control for everyday scripting.
* **Dance Diagrams** - Use `chore chart file.dance` to generate mermaid charts that illustrate program flow.
* **Built-In Linter** - `chore lint` keeps code elegant and consistent.
* **Stage Package Manager** - Install plugins and libraries with `chore stage`.
* **Scripting & REPL** - `chore run` executes scripts directly and `chore repl` provides an interactive shell.

Detailed specifications for each highlight are available in the
[docs](docs/) directory.

## Getting Started

ChoreLang is currently in the design phase. The toolchain is planned to be written in Go for ease of distribution. When ready, you will be able to install the compiler and runtime tools with a single command and begin dancing with code immediately.

The language favors immediate feedback. Compilation and runtime errors are presented clearly and highlight the exact file and line so you can fix missteps without missing a beat.

```
# install chorelang (placeholder)
$ chore install

# run a script
$ chore run hello.dance
```

### Example

```chorelang
sway i from 0 to 3 {
    spin print("step", i)
}
```

The example above shows a simple loop using dance-inspired keywords. Concurrency constructs let you orchestrate parallel movements easily, while the standard library grants direct access to OpenAI and Ollama for language model interactions. File operations and API discovery follow the same pattern, so every part of your program feels consistent and choreographed.

## Technical Specification

A detailed specification describing syntax, concurrency, and AI integrations lives in [TECH_SPEC.md](TECH_SPEC.md). The spec outlines how programs flow like choreography and how the compiler will be implemented.

Stay tuned and join the conversation as we build a language where code truly dances.

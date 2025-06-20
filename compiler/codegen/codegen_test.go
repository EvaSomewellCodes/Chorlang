package codegen

import (
	"strings"
	"testing"
	
	"github.com/chorlang/chorlang/compiler/lexer"
	"github.com/chorlang/chorlang/compiler/parser"
)

func TestGenerateSimpleProgram(t *testing.T) {
	input := `
dance x = 5
dance y = 10
spin print(x + y)
`
	
	expected := `package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 10
	fmt.Println((x + y))
}`
	
	result := generateAndCompare(t, input, expected)
	if result != expected {
		// Debug: show length comparison
		if len(result) != len(expected) {
			t.Errorf("Length mismatch: got %d, expected %d", len(result), len(expected))
		}
		// Debug: show byte-by-byte comparison
		for i := 0; i < len(result) && i < len(expected); i++ {
			if result[i] != expected[i] {
				t.Errorf("First difference at position %d: got %q, expected %q", i, result[i], expected[i])
				break
			}
		}
		t.Errorf("Generated code does not match expected.\nGot:\n%s\n\nExpected:\n%s", result, expected)
	}
}

func TestGenerateSwayLoop(t *testing.T) {
	input := `
sway i from 0 to 10 {
    spin print(i)
}
`
	
	expected := `package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}`
	
	result := generateAndCompare(t, input, expected)
	if result != expected {
		// Debug: show length comparison
		if len(result) != len(expected) {
			t.Errorf("Length mismatch: got %d, expected %d", len(result), len(expected))
		}
		// Debug: show byte-by-byte comparison
		for i := 0; i < len(result) && i < len(expected); i++ {
			if result[i] != expected[i] {
				t.Errorf("First difference at position %d: got %q, expected %q", i, result[i], expected[i])
				break
			}
		}
		t.Errorf("Generated code does not match expected.\nGot:\n%s\n\nExpected:\n%s", result, expected)
	}
}

func TestGenerateStartStatement(t *testing.T) {
	input := `
start sway i from 0 to 3 {
    spin print(i)
}
`
	
	expected := `package main

import (
	"fmt"
)

func main() {
	go func() {
		for i := 0; i <= 3; i++ {
			fmt.Println(i)
		}
	}()
}`
	
	result := generateAndCompare(t, input, expected)
	if result != expected {
		// Debug: show length comparison
		if len(result) != len(expected) {
			t.Errorf("Length mismatch: got %d, expected %d", len(result), len(expected))
		}
		// Debug: show byte-by-byte comparison
		for i := 0; i < len(result) && i < len(expected); i++ {
			if result[i] != expected[i] {
				t.Errorf("First difference at position %d: got %q, expected %q", i, result[i], expected[i])
				break
			}
		}
		t.Errorf("Generated code does not match expected.\nGot:\n%s\n\nExpected:\n%s", result, expected)
	}
}

func TestGenerateIfStatement(t *testing.T) {
	input := `
dance x = 5
dance y = 10
if x < y {
    spin print("x is less than y")
} else {
    spin print("x is not less than y")
}
`
	
	expected := `package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 10
	if (x < y) {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is not less than y")
	}
}`
	
	result := generateAndCompare(t, input, expected)
	if result != expected {
		// Debug: show length comparison
		if len(result) != len(expected) {
			t.Errorf("Length mismatch: got %d, expected %d", len(result), len(expected))
		}
		// Debug: show byte-by-byte comparison
		for i := 0; i < len(result) && i < len(expected); i++ {
			if result[i] != expected[i] {
				t.Errorf("First difference at position %d: got %q, expected %q", i, result[i], expected[i])
				break
			}
		}
		t.Errorf("Generated code does not match expected.\nGot:\n%s\n\nExpected:\n%s", result, expected)
	}
}

func generateAndCompare(t *testing.T, input, expected string) string {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) != 0 {
		t.Fatalf("Parser errors: %v", p.Errors())
	}
	
	g := New()
	result, err := g.Generate(program)
	if err != nil {
		t.Fatalf("Code generation error: %v", err)
	}
	
	// Normalize whitespace for comparison
	result = normalizeWhitespace(result)
	expected = normalizeWhitespace(expected)
	
	return result
}

func normalizeWhitespace(s string) string {
	// Split by newline and remove trailing spaces
	lines := strings.Split(s, "\n")
	var normalized []string
	for _, line := range lines {
		normalized = append(normalized, strings.TrimRight(line, " \t"))
	}
	// Join and trim
	result := strings.TrimSpace(strings.Join(normalized, "\n"))
	return result
}
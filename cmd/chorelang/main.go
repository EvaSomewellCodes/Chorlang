package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	
	"github.com/chorlang/chorlang/compiler/codegen"
	"github.com/chorlang/chorlang/compiler/lexer"
	"github.com/chorlang/chorlang/compiler/parser"
)

func main() {
	var (
		output  = flag.String("o", "", "output file name")
		compile = flag.Bool("c", false, "compile to binary")
		run     = flag.Bool("r", false, "run the program after compilation")
		help    = flag.Bool("h", false, "show help")
	)
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ChoreLang Compiler\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <input.chore>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s hello.chore                  # Generate Go code\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -c hello.chore               # Compile to binary\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -r hello.chore               # Run immediately\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -c -o myapp hello.chore      # Compile with custom output name\n", os.Args[0])
	}
	
	flag.Parse()
	
	if *help || flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}
	
	inputFile := flag.Arg(0)
	
	// Read input file
	source, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", inputFile, err)
		os.Exit(1)
	}
	
	// Lexing
	l := lexer.New(string(source))
	
	// Parsing
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		fmt.Fprintf(os.Stderr, "Parser errors:\n")
		for _, err := range p.Errors() {
			fmt.Fprintf(os.Stderr, "  %s\n", err)
		}
		os.Exit(1)
	}
	
	// Code generation
	g := codegen.New()
	goCode, err := g.Generate(program)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Code generation error: %v\n", err)
		os.Exit(1)
	}
	
	// Determine output file name
	baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
	goFile := baseName + ".go"
	
	if !*compile && !*run {
		// Just output the Go code
		if *output != "" {
			goFile = *output
		}
		err = ioutil.WriteFile(goFile, []byte(goCode), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing Go file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Generated %s\n", goFile)
		return
	}
	
	// Write temporary Go file for compilation
	tmpGoFile := filepath.Join(os.TempDir(), baseName+"_chorelang.go")
	err = ioutil.WriteFile(tmpGoFile, []byte(goCode), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing temporary Go file: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tmpGoFile)
	
	if *compile || *run {
		// Compile to binary
		outputBinary := baseName
		if *output != "" {
			outputBinary = *output
		}
		
		cmd := exec.Command("go", "build", "-o", outputBinary, tmpGoFile)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		
		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Compilation error: %v\n", err)
			os.Exit(1)
		}
		
		if *compile && !*run {
			fmt.Printf("Compiled to %s\n", outputBinary)
		}
		
		if *run {
			// Run the compiled binary
			cmd = exec.Command("./" + outputBinary)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			
			err = cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
				os.Exit(1)
			}
			
			// Clean up binary after running
			os.Remove(outputBinary)
		}
	}
}
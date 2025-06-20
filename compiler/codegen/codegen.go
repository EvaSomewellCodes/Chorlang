package codegen

import (
	"bytes"
	"fmt"
	"strings"
	
	"github.com/chorlang/chorlang/compiler/ast"
)

type CodeGenerator struct {
	output      bytes.Buffer
	indent      int
	errors      []string
	hasMain     bool
	imports     map[string]bool
	declaredVars map[string]bool
	scopeStack  []map[string]bool
}

func New() *CodeGenerator {
	g := &CodeGenerator{
		imports:      make(map[string]bool),
		declaredVars: make(map[string]bool),
		scopeStack:   []map[string]bool{},
	}
	// Push initial scope
	g.pushScope()
	return g
}

func (g *CodeGenerator) Generate(program *ast.Program) (string, error) {
	// Generate package declaration
	g.write("package main\n\n")
	
	// Process all statements to collect imports
	for _, stmt := range program.Statements {
		g.collectImports(stmt)
	}
	
	// Generate imports
	if len(g.imports) > 0 {
		g.write("import (\n")
		g.indent++
		for imp := range g.imports {
			g.writeLine(fmt.Sprintf(`"%s"`, imp))
		}
		g.indent--
		g.write(")\n\n")
	}
	
	// Generate main function with all statements
	g.write("func main() {\n")
	g.indent++
	
	// Generate statements inside main
	for _, stmt := range program.Statements {
		if err := g.generateStatement(stmt); err != nil {
			return "", err
		}
	}
	
	g.indent--
	g.write("}\n")
	
	return g.output.String(), nil
}

func (g *CodeGenerator) collectImports(stmt ast.Statement) {
	switch s := stmt.(type) {
	case *ast.ExpressionStatement:
		if spin, ok := s.Expression.(*ast.SpinExpression); ok {
			if ident, ok := spin.Function.(*ast.Identifier); ok {
				if ident.Value == "print" || ident.Value == "println" {
					g.imports["fmt"] = true
				}
			}
		}
	case *ast.SwayStatement:
		g.collectImportsFromBlock(s.Body)
	case *ast.StartStatement:
		g.collectImports(s.Statement)
	case *ast.IfStatement:
		g.collectImportsFromBlock(s.Consequence)
		if s.Alternative != nil {
			g.collectImportsFromBlock(s.Alternative)
		}
	}
}

func (g *CodeGenerator) collectImportsFromBlock(block *ast.BlockStatement) {
	for _, stmt := range block.Statements {
		g.collectImports(stmt)
	}
}

func (g *CodeGenerator) generateStatement(stmt ast.Statement) error {
	switch s := stmt.(type) {
	case *ast.DanceStatement:
		return g.generateDanceStatement(s)
	case *ast.ExpressionStatement:
		return g.generateExpressionStatement(s)
	case *ast.SwayStatement:
		return g.generateSwayStatement(s)
	case *ast.StartStatement:
		return g.generateStartStatement(s)
	case *ast.SendStatement:
		return g.generateSendStatement(s)
	case *ast.IfStatement:
		return g.generateIfStatement(s)
	default:
		return fmt.Errorf("unknown statement type: %T", stmt)
	}
}

func (g *CodeGenerator) generateDanceStatement(stmt *ast.DanceStatement) error {
	g.writeIndent()
	
	// Check if variable is already declared in current scope
	varName := stmt.Name.Value
	if g.isVarDeclared(varName) {
		// Use = for reassignment
		g.write(varName)
		g.write(" = ")
	} else {
		// Use := for new declaration
		g.write(varName)
		g.write(" := ")
		g.declareVar(varName)
	}
	
	if err := g.generateExpression(stmt.Value); err != nil {
		return err
	}
	
	g.write("\n")
	return nil
}

func (g *CodeGenerator) generateExpressionStatement(stmt *ast.ExpressionStatement) error {
	g.writeIndent()
	if err := g.generateExpression(stmt.Expression); err != nil {
		return err
	}
	g.write("\n")
	return nil
}

func (g *CodeGenerator) generateSwayStatement(stmt *ast.SwayStatement) error {
	g.writeIndent()
	g.write("for ")
	g.write(stmt.Variable.Value)
	g.write(" := ")
	
	if err := g.generateExpression(stmt.From); err != nil {
		return err
	}
	
	g.write("; ")
	g.write(stmt.Variable.Value)
	g.write(" <= ")
	
	if err := g.generateExpression(stmt.To); err != nil {
		return err
	}
	
	g.write("; ")
	g.write(stmt.Variable.Value)
	g.write("++ {\n")
	
	// Push new scope for loop body
	g.pushScope()
	// Declare loop variable in new scope
	g.declareVar(stmt.Variable.Value)
	
	g.indent++
	for _, s := range stmt.Body.Statements {
		if err := g.generateStatement(s); err != nil {
			return err
		}
	}
	g.indent--
	
	// Pop scope
	g.popScope()
	
	g.writeIndent()
	g.write("}\n")
	
	return nil
}

func (g *CodeGenerator) generateStartStatement(stmt *ast.StartStatement) error {
	g.writeIndent()
	g.write("go func() {\n")
	
	// Push new scope for goroutine
	g.pushScope()
	g.indent++
	
	if err := g.generateStatement(stmt.Statement); err != nil {
		return err
	}
	
	g.indent--
	g.popScope()
	
	g.writeIndent()
	g.write("}()\n")
	
	return nil
}

func (g *CodeGenerator) generateSendStatement(stmt *ast.SendStatement) error {
	g.writeIndent()
	
	if err := g.generateExpression(stmt.Channel); err != nil {
		return err
	}
	
	g.write(" <- ")
	
	if err := g.generateExpression(stmt.Value); err != nil {
		return err
	}
	
	g.write("\n")
	return nil
}

func (g *CodeGenerator) generateIfStatement(stmt *ast.IfStatement) error {
	g.writeIndent()
	g.write("if ")
	
	if err := g.generateExpression(stmt.Condition); err != nil {
		return err
	}
	
	g.write(" {\n")
	
	// Push scope for consequence
	g.pushScope()
	g.indent++
	for _, s := range stmt.Consequence.Statements {
		if err := g.generateStatement(s); err != nil {
			return err
		}
	}
	g.indent--
	g.popScope()
	
	g.writeIndent()
	g.write("}")
	
	if stmt.Alternative != nil {
		g.write(" else {\n")
		// Push scope for alternative
		g.pushScope()
		g.indent++
		for _, s := range stmt.Alternative.Statements {
			if err := g.generateStatement(s); err != nil {
				return err
			}
		}
		g.indent--
		g.popScope()
		g.writeIndent()
		g.write("}")
	}
	
	g.write("\n")
	return nil
}

func (g *CodeGenerator) generateExpression(exp ast.Expression) error {
	switch e := exp.(type) {
	case *ast.Identifier:
		g.write(e.Value)
	case *ast.IntegerLiteral:
		g.write(fmt.Sprintf("%d", e.Value))
	case *ast.FloatLiteral:
		g.write(fmt.Sprintf("%f", e.Value))
	case *ast.StringLiteral:
		g.write(fmt.Sprintf(`"%s"`, e.Value))
	case *ast.Boolean:
		g.write(fmt.Sprintf("%t", e.Value))
	case *ast.InfixExpression:
		g.write("(")
		if err := g.generateExpression(e.Left); err != nil {
			return err
		}
		g.write(fmt.Sprintf(" %s ", e.Operator))
		if err := g.generateExpression(e.Right); err != nil {
			return err
		}
		g.write(")")
	case *ast.SpinExpression:
		return g.generateSpinExpression(e)
	case *ast.FlowExpression:
		return g.generateFlowExpression(e)
	case *ast.MatchExpression:
		return g.generateMatchExpression(e)
	default:
		return fmt.Errorf("unknown expression type: %T", exp)
	}
	return nil
}

func (g *CodeGenerator) generateSpinExpression(exp *ast.SpinExpression) error {
	// Handle built-in functions
	if ident, ok := exp.Function.(*ast.Identifier); ok {
		switch ident.Value {
		case "print", "println":
			g.write("fmt.Println(")
			for i, arg := range exp.Arguments {
				if i > 0 {
					g.write(", ")
				}
				if err := g.generateExpression(arg); err != nil {
					return err
				}
			}
			g.write(")")
			return nil
		}
	}
	
	// Regular function call
	if err := g.generateExpression(exp.Function); err != nil {
		return err
	}
	
	g.write("(")
	for i, arg := range exp.Arguments {
		if i > 0 {
			g.write(", ")
		}
		if err := g.generateExpression(arg); err != nil {
			return err
		}
	}
	g.write(")")
	
	return nil
}

func (g *CodeGenerator) generateFlowExpression(exp *ast.FlowExpression) error {
	// For now, generate as a channel type
	// flow channel<int> becomes chan int
	if ident, ok := exp.ChannelType.(*ast.Identifier); ok && ident.Value == "channel" {
		g.write("make(chan ")
		// TODO: Parse the type parameter properly
		g.write("interface{})")
	} else {
		g.write("make(")
		if err := g.generateExpression(exp.ChannelType); err != nil {
			return err
		}
		g.write(")")
	}
	return nil
}

func (g *CodeGenerator) generateMatchExpression(exp *ast.MatchExpression) error {
	// Generate match as a switch statement
	g.write("func() interface{} {\n")
	g.indent++
	g.writeIndent()
	g.write("switch ")
	if err := g.generateExpression(exp.Expression); err != nil {
		return err
	}
	g.write(" {\n")
	
	for _, c := range exp.Cases {
		g.writeIndent()
		g.write("case ")
		if err := g.generateExpression(c.Pattern); err != nil {
			return err
		}
		g.write(":\n")
		g.indent++
		g.writeIndent()
		g.write("return ")
		if err := g.generateExpression(c.Consequence); err != nil {
			return err
		}
		g.write("\n")
		g.indent--
	}
	
	g.writeIndent()
	g.write("}\n")
	g.writeIndent()
	g.write("return nil\n")
	g.indent--
	g.writeIndent()
	g.write("}()")
	
	return nil
}

func (g *CodeGenerator) write(s string) {
	g.output.WriteString(s)
}

func (g *CodeGenerator) writeLine(s string) {
	g.writeIndent()
	g.output.WriteString(s)
	g.output.WriteString("\n")
}

func (g *CodeGenerator) writeIndent() {
	g.output.WriteString(strings.Repeat("\t", g.indent))
}

func (g *CodeGenerator) pushScope() {
	newScope := make(map[string]bool)
	// Copy parent scope variables
	if len(g.scopeStack) > 0 {
		parentScope := g.scopeStack[len(g.scopeStack)-1]
		for k, v := range parentScope {
			newScope[k] = v
		}
	}
	g.scopeStack = append(g.scopeStack, newScope)
}

func (g *CodeGenerator) popScope() {
	if len(g.scopeStack) > 0 {
		g.scopeStack = g.scopeStack[:len(g.scopeStack)-1]
	}
}

func (g *CodeGenerator) isVarDeclared(name string) bool {
	if len(g.scopeStack) > 0 {
		currentScope := g.scopeStack[len(g.scopeStack)-1]
		return currentScope[name]
	}
	return false
}

func (g *CodeGenerator) declareVar(name string) {
	if len(g.scopeStack) > 0 {
		currentScope := g.scopeStack[len(g.scopeStack)-1]
		currentScope[name] = true
	}
}
package parser

import (
	"testing"
	
	"github.com/chorlang/chorlang/compiler/ast"
	"github.com/chorlang/chorlang/compiler/lexer"
)

func TestDanceStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"dance x = 5;", "x", 5},
		{"dance y = true;", "y", true},
		{"dance foobar = y;", "foobar", "y"},
		{"dance str = \"hello world\";", "str", "hello world"},
		{"dance f = 10.5;", "f", 10.5},
	}
	
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}
		
		stmt := program.Statements[0]
		if !testDanceStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
		
		val := stmt.(*ast.DanceStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func TestSwayStatement(t *testing.T) {
	input := `
sway i from 0 to 10 {
    spin print(i)
}`
	
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.SwayStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.SwayStatement. got=%T",
			program.Statements[0])
	}
	
	if stmt.Variable.Value != "i" {
		t.Fatalf("stmt.Variable.Value not 'i'. got=%s", stmt.Variable.Value)
	}
	
	if !testLiteralExpression(t, stmt.From, 0) {
		return
	}
	
	if !testLiteralExpression(t, stmt.To, 10) {
		return
	}
	
	if len(stmt.Body.Statements) != 1 {
		t.Fatalf("stmt.Body.Statements does not contain 1 statements. got=%d",
			len(stmt.Body.Statements))
	}
}

func TestSpinExpression(t *testing.T) {
	input := "spin print(1, 2 * 3, 4 + 5);"
	
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	
	exp, ok := stmt.Expression.(*ast.SpinExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.SpinExpression. got=%T",
			stmt.Expression)
	}
	
	if !testIdentifier(t, exp.Function, "print") {
		return
	}
	
	if len(exp.Arguments) != 3 {
		t.Fatalf("wrong length of arguments. got=%d", len(exp.Arguments))
	}
	
	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func TestStartStatement(t *testing.T) {
	input := `start sway i from 0 to 3 {
    send steps <- i
}`
	
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.StartStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.StartStatement. got=%T",
			program.Statements[0])
	}
	
	swayStmt, ok := stmt.Statement.(*ast.SwayStatement)
	if !ok {
		t.Fatalf("stmt.Statement is not ast.SwayStatement. got=%T",
			stmt.Statement)
	}
	
	if swayStmt.Variable.Value != "i" {
		t.Fatalf("swayStmt.Variable.Value not 'i'. got=%s", swayStmt.Variable.Value)
	}
}

func TestIfStatement(t *testing.T) {
	input := `if x < y {
    dance z = x
} else {
    dance z = y
}`
	
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.IfStatement. got=%T",
			program.Statements[0])
	}
	
	if !testInfixExpression(t, stmt.Condition, "x", "<", "y") {
		return
	}
	
	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n",
			len(stmt.Consequence.Statements))
	}
	
	if stmt.Alternative == nil {
		t.Errorf("stmt.Alternative was nil")
	}
	
	if len(stmt.Alternative.Statements) != 1 {
		t.Errorf("alternative is not 1 statements. got=%d\n",
			len(stmt.Alternative.Statements))
	}
}

func TestMatchExpression(t *testing.T) {
	input := `dance result = match item {
    when Note(n): flow process_note(n)
    when Rest(): flow handle_rest()
}`
	
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.DanceStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.DanceStatement. got=%T",
			program.Statements[0])
	}
	
	match, ok := stmt.Value.(*ast.MatchExpression)
	if !ok {
		t.Fatalf("stmt.Value is not ast.MatchExpression. got=%T", stmt.Value)
	}
	
	if !testIdentifier(t, match.Expression, "item") {
		return
	}
	
	if len(match.Cases) != 2 {
		t.Fatalf("match.Cases does not contain 2 cases. got=%d", len(match.Cases))
	}
}

func testDanceStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "dance" {
		t.Errorf("s.TokenLiteral not 'dance'. got=%q", s.TokenLiteral())
		return false
	}
	
	danceStmt, ok := s.(*ast.DanceStatement)
	if !ok {
		t.Errorf("s not *ast.DanceStatement. got=%T", s)
		return false
	}
	
	if danceStmt.Name.Value != name {
		t.Errorf("danceStmt.Name.Value not '%s'. got=%s", name, danceStmt.Name.Value)
		return false
	}
	
	if danceStmt.Name.TokenLiteral() != name {
		t.Errorf("danceStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, danceStmt.Name.TokenLiteral())
		return false
	}
	
	return true
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
	operator string, right interface{}) bool {
	
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}
	
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}
	
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	
	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case float64:
		return testFloatLiteral(t, exp, v)
	case string:
		if _, ok := exp.(*ast.StringLiteral); ok {
			return testStringLiteral(t, exp, v)
		}
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	
	return true
}

func testFloatLiteral(t *testing.T, fl ast.Expression, value float64) bool {
	float, ok := fl.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("fl not *ast.FloatLiteral. got=%T", fl)
		return false
	}
	
	if float.Value != value {
		t.Errorf("float.Value not %f. got=%f", value, float.Value)
		return false
	}
	
	return true
}

func testStringLiteral(t *testing.T, sl ast.Expression, value string) bool {
	str, ok := sl.(*ast.StringLiteral)
	if !ok {
		t.Errorf("sl not *ast.StringLiteral. got=%T", sl)
		return false
	}
	
	if str.Value != value {
		t.Errorf("str.Value not %s. got=%s", value, str.Value)
		return false
	}
	
	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}
	
	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}
	
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
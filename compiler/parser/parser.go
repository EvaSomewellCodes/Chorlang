package parser

import (
	"fmt"
	"strconv"
	
	"github.com/chorlang/chorlang/compiler/ast"
	"github.com/chorlang/chorlang/compiler/lexer"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string
	
	curToken  lexer.Token
	peekToken lexer.Token
	
	prefixParseFns map[lexer.TokenType]prefixParseFn
	infixParseFns  map[lexer.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	
	p.prefixParseFns = make(map[lexer.TokenType]prefixParseFn)
	p.registerPrefix(lexer.IDENT, p.parseIdentifier)
	p.registerPrefix(lexer.INT, p.parseIntegerLiteral)
	p.registerPrefix(lexer.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(lexer.STRING, p.parseStringLiteral)
	p.registerPrefix(lexer.TRUE, p.parseBoolean)
	p.registerPrefix(lexer.FALSE, p.parseBoolean)
	p.registerPrefix(lexer.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(lexer.SPIN, p.parseSpinExpression)
	p.registerPrefix(lexer.FLOW, p.parseFlowExpression)
	p.registerPrefix(lexer.MATCH, p.parseMatchExpression)
	
	p.infixParseFns = make(map[lexer.TokenType]infixParseFn)
	p.registerInfix(lexer.PLUS, p.parseInfixExpression)
	p.registerInfix(lexer.MINUS, p.parseInfixExpression)
	p.registerInfix(lexer.SLASH, p.parseInfixExpression)
	p.registerInfix(lexer.ASTERISK, p.parseInfixExpression)
	p.registerInfix(lexer.EQ, p.parseInfixExpression)
	p.registerInfix(lexer.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(lexer.LT, p.parseInfixExpression)
	p.registerInfix(lexer.GT, p.parseInfixExpression)
	p.registerInfix(lexer.LTE, p.parseInfixExpression)
	p.registerInfix(lexer.GTE, p.parseInfixExpression)
	p.registerInfix(lexer.MATCH_OP, p.parseInfixExpression)
	
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t lexer.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType lexer.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	
	for !p.curTokenIs(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case lexer.DANCE:
		return p.parseDanceStatement()
	case lexer.SWAY:
		return p.parseSwayStatement()
	case lexer.START:
		return p.parseStartStatement()
	case lexer.SEND_KW:
		return p.parseSendStatement()
	case lexer.IF:
		return p.parseIfStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseDanceStatement() *ast.DanceStatement {
	stmt := &ast.DanceStatement{Token: p.curToken}
	
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	
	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}
	
	p.nextToken()
	
	stmt.Value = p.parseExpression(LOWEST)
	
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	
	return stmt
}

func (p *Parser) parseSwayStatement() *ast.SwayStatement {
	stmt := &ast.SwayStatement{Token: p.curToken}
	
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	
	stmt.Variable = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	
	if !p.expectPeek(lexer.FROM) {
		return nil
	}
	
	p.nextToken()
	stmt.From = p.parseExpression(LOWEST)
	
	if !p.expectPeek(lexer.TO) {
		return nil
	}
	
	p.nextToken()
	stmt.To = p.parseExpression(LOWEST)
	
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	
	stmt.Body = p.parseBlockStatement()
	
	return stmt
}

func (p *Parser) parseStartStatement() *ast.StartStatement {
	stmt := &ast.StartStatement{Token: p.curToken}
	
	p.nextToken()
	
	stmt.Statement = p.parseStatement()
	
	return stmt
}

func (p *Parser) parseSendStatement() *ast.SendStatement {
	stmt := &ast.SendStatement{Token: p.curToken}
	
	p.nextToken()
	stmt.Channel = p.parseExpression(LOWEST)
	
	if !p.expectPeek(lexer.SEND) {
		return nil
	}
	
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	
	return stmt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}
	
	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)
	
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	
	stmt.Consequence = p.parseBlockStatement()
	
	if p.peekTokenIs(lexer.ELSE) {
		p.nextToken()
		
		if !p.expectPeek(lexer.LBRACE) {
			return nil
		}
		
		stmt.Alternative = p.parseBlockStatement()
	}
	
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	
	stmt.Expression = p.parseExpression(LOWEST)
	
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	
	for !p.peekTokenIs(lexer.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		
		p.nextToken()
		
		leftExp = infix(leftExp)
	}
	
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	
	lit.Value = value
	
	return lit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.curToken}
	
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	
	lit.Value = value
	
	return lit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(lexer.TRUE)}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	
	exp := p.parseExpression(LOWEST)
	
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	
	return exp
}

func (p *Parser) parseSpinExpression() ast.Expression {
	exp := &ast.SpinExpression{Token: p.curToken}
	
	p.nextToken()
	exp.Function = p.parseExpression(LOWEST)
	
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	
	exp.Arguments = p.parseExpressionList(lexer.RPAREN)
	
	return exp
}

func (p *Parser) parseFlowExpression() ast.Expression {
	exp := &ast.FlowExpression{Token: p.curToken}
	
	p.nextToken()
	exp.ChannelType = p.parseExpression(LOWEST)
	
	return exp
}

func (p *Parser) parseMatchExpression() ast.Expression {
	exp := &ast.MatchExpression{Token: p.curToken}
	
	p.nextToken()
	exp.Expression = p.parseExpression(LOWEST)
	
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	
	exp.Cases = []*ast.WhenCase{}
	
	p.nextToken()
	
	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		if p.curTokenIs(lexer.WHEN) {
			whenCase := p.parseWhenCase()
			if whenCase != nil {
				exp.Cases = append(exp.Cases, whenCase)
			}
		}
		p.nextToken()
	}
	
	return exp
}

func (p *Parser) parseWhenCase() *ast.WhenCase {
	whenCase := &ast.WhenCase{Token: p.curToken}
	
	p.nextToken()
	whenCase.Pattern = p.parsePattern()
	
	if !p.expectPeek(lexer.COLON) {
		return nil
	}
	
	p.nextToken()
	whenCase.Consequence = p.parseExpression(LOWEST)
	
	return whenCase
}

func (p *Parser) parsePattern() ast.Expression {
	// For now, we'll parse patterns as expressions
	// This handles simple patterns like Note(n), Rest(), etc.
	pattern := p.parseExpression(CALL)
	
	// Skip optional parentheses for pattern parameters
	if p.peekTokenIs(lexer.LPAREN) {
		p.nextToken()
		// Parse pattern parameters if any
		for !p.curTokenIs(lexer.RPAREN) && !p.curTokenIs(lexer.EOF) {
			p.nextToken()
		}
	}
	
	return pattern
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	
	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}
	
	p.nextToken()
	
	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	
	return block
}

func (p *Parser) parseExpressionList(end lexer.TokenType) []ast.Expression {
	list := []ast.Expression{}
	
	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}
	
	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))
	
	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}
	
	if !p.expectPeek(end) {
		return nil
	}
	
	return list
}

var precedences = map[lexer.TokenType]int{
	lexer.EQ:       EQUALS,
	lexer.NOT_EQ:   EQUALS,
	lexer.LT:       LESSGREATER,
	lexer.GT:       LESSGREATER,
	lexer.LTE:      LESSGREATER,
	lexer.GTE:      LESSGREATER,
	lexer.MATCH_OP: EQUALS,
	lexer.PLUS:     SUM,
	lexer.MINUS:    SUM,
	lexer.SLASH:    PRODUCT,
	lexer.ASTERISK: PRODUCT,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	
	return LOWEST
}

func (p *Parser) noPrefixParseFnError(t lexer.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
package ast

import (
	"bytes"
	"github.com/chorlang/chorlang/compiler/lexer"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	
	return out.String()
}

// Identifier
type Identifier struct {
	Token lexer.Token // the IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// Integer Literal
type IntegerLiteral struct {
	Token lexer.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// Float Literal
type FloatLiteral struct {
	Token lexer.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }

// String Literal
type StringLiteral struct {
	Token lexer.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" }

// Boolean
type Boolean struct {
	Token lexer.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// Dance Statement (variable declaration)
type DanceStatement struct {
	Token lexer.Token // the DANCE token
	Name  *Identifier
	Value Expression
}

func (ds *DanceStatement) statementNode()       {}
func (ds *DanceStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DanceStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(ds.TokenLiteral() + " ")
	out.WriteString(ds.Name.String())
	out.WriteString(" = ")
	
	if ds.Value != nil {
		out.WriteString(ds.Value.String())
	}
	
	return out.String()
}

// Expression Statement
type ExpressionStatement struct {
	Token      lexer.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Infix Expression
type InfixExpression struct {
	Token    lexer.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	
	return out.String()
}

// Spin Expression (function call)
type SpinExpression struct {
	Token     lexer.Token // The SPIN token
	Function  Expression   // Identifier or function literal
	Arguments []Expression
}

func (se *SpinExpression) expressionNode()      {}
func (se *SpinExpression) TokenLiteral() string { return se.Token.Literal }
func (se *SpinExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString(se.TokenLiteral() + " ")
	out.WriteString(se.Function.String())
	out.WriteString("(")
	
	for i, arg := range se.Arguments {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(arg.String())
	}
	
	out.WriteString(")")
	
	return out.String()
}

// Sway Statement (for loop)
type SwayStatement struct {
	Token    lexer.Token // The SWAY token
	Variable *Identifier
	From     Expression
	To       Expression
	Body     *BlockStatement
}

func (ss *SwayStatement) statementNode()       {}
func (ss *SwayStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SwayStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.Variable.String())
	out.WriteString(" from ")
	out.WriteString(ss.From.String())
	out.WriteString(" to ")
	out.WriteString(ss.To.String())
	out.WriteString(" ")
	out.WriteString(ss.Body.String())
	
	return out.String()
}

// Block Statement
type BlockStatement struct {
	Token      lexer.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString("{\n")
	
	for _, s := range bs.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	
	out.WriteString("}")
	
	return out.String()
}

// Flow Expression (channel declaration)
type FlowExpression struct {
	Token       lexer.Token // The FLOW token
	ChannelType Expression  // e.g., channel<int>
}

func (fe *FlowExpression) expressionNode()      {}
func (fe *FlowExpression) TokenLiteral() string { return fe.Token.Literal }
func (fe *FlowExpression) String() string {
	return fe.TokenLiteral() + " " + fe.ChannelType.String()
}

// Start Statement (goroutine)
type StartStatement struct {
	Token     lexer.Token // The START token
	Statement Statement
}

func (ss *StartStatement) statementNode()       {}
func (ss *StartStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *StartStatement) String() string {
	return ss.TokenLiteral() + " " + ss.Statement.String()
}

// Send Statement (channel send)
type SendStatement struct {
	Token   lexer.Token // The SEND token
	Channel Expression
	Value   Expression
}

func (ss *SendStatement) statementNode()       {}
func (ss *SendStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SendStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.Channel.String())
	out.WriteString(" <- ")
	out.WriteString(ss.Value.String())
	
	return out.String()
}

// If Statement
type IfStatement struct {
	Token       lexer.Token // The IF token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString("if ")
	out.WriteString(is.Condition.String())
	out.WriteString(" ")
	out.WriteString(is.Consequence.String())
	
	if is.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(is.Alternative.String())
	}
	
	return out.String()
}

// Match Expression (pattern matching)
type MatchExpression struct {
	Token      lexer.Token // The MATCH token
	Expression Expression
	Cases      []*WhenCase
}

func (me *MatchExpression) expressionNode()      {}
func (me *MatchExpression) TokenLiteral() string { return me.Token.Literal }
func (me *MatchExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString(me.TokenLiteral() + " ")
	out.WriteString(me.Expression.String())
	out.WriteString(" {\n")
	
	for _, c := range me.Cases {
		out.WriteString(c.String())
		out.WriteString("\n")
	}
	
	out.WriteString("}")
	
	return out.String()
}

// When Case (for pattern matching)
type WhenCase struct {
	Token      lexer.Token // The WHEN token
	Pattern    Expression
	Consequence Expression
}

func (wc *WhenCase) String() string {
	var out bytes.Buffer
	
	out.WriteString(wc.Token.Literal + " ")
	out.WriteString(wc.Pattern.String())
	out.WriteString(": ")
	out.WriteString(wc.Consequence.String())
	
	return out.String()
}
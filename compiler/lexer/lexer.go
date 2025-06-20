package lexer

import (
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		r, size := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.ch = r
		l.position = l.readPosition
		l.readPosition += size
		
		if l.ch == '\n' {
			l.line++
			l.column = 0
		} else {
			l.column++
		}
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
		return r
	}
}

func (l *Lexer) NextToken() Token {
	var tok Token
	
	l.skipWhitespace()
	
	tok.Line = l.line
	tok.Column = l.column
	
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.makeToken(EQ, string(ch)+string(l.ch))
		} else if l.peekChar() == '~' {
			ch := l.ch
			l.readChar()
			tok = l.makeToken(MATCH_OP, string(ch)+string(l.ch))
		} else {
			tok = l.makeToken(ASSIGN, string(l.ch))
		}
	case '+':
		tok = l.makeToken(PLUS, string(l.ch))
	case '-':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = l.makeToken(ARROW, string(ch)+string(l.ch))
		} else {
			tok = l.makeToken(MINUS, string(l.ch))
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.makeToken(NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = l.makeToken(BANG, string(l.ch))
		}
	case '/':
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken()
		} else {
			tok = l.makeToken(SLASH, string(l.ch))
		}
	case '*':
		tok = l.makeToken(ASTERISK, string(l.ch))
	case '<':
		if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			tok = l.makeToken(SEND, string(ch)+string(l.ch))
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.makeToken(LTE, string(ch)+string(l.ch))
		} else {
			tok = l.makeToken(LT, string(l.ch))
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.makeToken(GTE, string(ch)+string(l.ch))
		} else {
			tok = l.makeToken(GT, string(l.ch))
		}
	case ',':
		tok = l.makeToken(COMMA, string(l.ch))
	case ';':
		tok = l.makeToken(SEMICOLON, string(l.ch))
	case ':':
		tok = l.makeToken(COLON, string(l.ch))
	case '(':
		tok = l.makeToken(LPAREN, string(l.ch))
	case ')':
		tok = l.makeToken(RPAREN, string(l.ch))
	case '{':
		tok = l.makeToken(LBRACE, string(l.ch))
	case '}':
		tok = l.makeToken(RBRACE, string(l.ch))
	case '[':
		tok = l.makeToken(LBRACKET, string(l.ch))
	case ']':
		tok = l.makeToken(RBRACKET, string(l.ch))
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
		tok.Line = l.line
		tok.Column = l.column
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type, tok.Literal = l.readNumber()
			return tok
		} else {
			tok = l.makeToken(ILLEGAL, string(l.ch))
		}
	}
	
	l.readChar()
	return tok
}

func (l *Lexer) makeToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal, Line: l.line, Column: l.column}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (TokenType, string) {
	position := l.position
	tokenType := INT
	
	for isDigit(l.ch) {
		l.readChar()
	}
	
	// Check for float
	if l.ch == '.' && isDigit(l.peekChar()) {
		tokenType = FLOAT
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	
	return tokenType, l.input[position:l.position]
}

func (l *Lexer) readString() string {
	// Current position is at the opening quote
	// Move past the opening quote
	l.readChar()
	position := l.position
	
	for l.ch != '"' && l.ch != 0 {
		if l.ch == '\\' {
			l.readChar()
			if l.ch == 0 {
				break
			}
		}
		l.readChar()
	}
	
	result := l.input[position:l.position]
	
	// Consume the closing quote
	if l.ch == '"' {
		l.readChar()
	}
	
	return result
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
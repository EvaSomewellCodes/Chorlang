package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `dance x = 5
dance y = 10.5
sway i from 0 to 10 {
    spin print(i)
}
flow channel<int> steps
start sway i from 0 to 3 {
    send steps <- i
}
if x =~ /pattern/ {
    dance result = match item {
        when Note(n): flow process_note(n)
        when Rest(): flow handle_rest()
    }
}
// This is a comment
"hello world"
== != <= >= ->`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{DANCE, "dance"},
		{IDENT, "x"},
		{ASSIGN, "="},
		{INT, "5"},
		{DANCE, "dance"},
		{IDENT, "y"},
		{ASSIGN, "="},
		{FLOAT, "10.5"},
		{SWAY, "sway"},
		{IDENT, "i"},
		{FROM, "from"},
		{INT, "0"},
		{TO, "to"},
		{INT, "10"},
		{LBRACE, "{"},
		{SPIN, "spin"},
		{IDENT, "print"},
		{LPAREN, "("},
		{IDENT, "i"},
		{RPAREN, ")"},
		{RBRACE, "}"},
		{FLOW, "flow"},
		{IDENT, "channel"},
		{LT, "<"},
		{IDENT, "int"},
		{GT, ">"},
		{IDENT, "steps"},
		{START, "start"},
		{SWAY, "sway"},
		{IDENT, "i"},
		{FROM, "from"},
		{INT, "0"},
		{TO, "to"},
		{INT, "3"},
		{LBRACE, "{"},
		{SEND_KW, "send"},
		{IDENT, "steps"},
		{SEND, "<-"},
		{IDENT, "i"},
		{RBRACE, "}"},
		{IF, "if"},
		{IDENT, "x"},
		{MATCH_OP, "=~"},
		{SLASH, "/"},
		{IDENT, "pattern"},
		{SLASH, "/"},
		{LBRACE, "{"},
		{DANCE, "dance"},
		{IDENT, "result"},
		{ASSIGN, "="},
		{MATCH, "match"},
		{IDENT, "item"},
		{LBRACE, "{"},
		{WHEN, "when"},
		{IDENT, "Note"},
		{LPAREN, "("},
		{IDENT, "n"},
		{RPAREN, ")"},
		{COLON, ":"},
		{FLOW, "flow"},
		{IDENT, "process_note"},
		{LPAREN, "("},
		{IDENT, "n"},
		{RPAREN, ")"},
		{WHEN, "when"},
		{IDENT, "Rest"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{COLON, ":"},
		{FLOW, "flow"},
		{IDENT, "handle_rest"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{RBRACE, "}"},
		{RBRACE, "}"},
		{STRING, "hello world"},
		{EQ, "=="},
		{NOT_EQ, "!="},
		{LTE, "<="},
		{GTE, ">="},
		{ARROW, "->"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q (literal=%q)",
				i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
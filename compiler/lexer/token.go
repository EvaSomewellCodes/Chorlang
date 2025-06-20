package lexer

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	
	// Literals
	IDENT
	INT
	FLOAT
	STRING
	
	// Operators
	ASSIGN     // =
	PLUS       // +
	MINUS      // -
	ASTERISK   // *
	SLASH      // /
	BANG       // !
	LT         // <
	GT         // >
	EQ         // ==
	NOT_EQ     // !=
	LTE        // <=
	GTE        // >=
	ARROW      // ->
	SEND       // <-
	MATCH_OP   // =~
	
	// Delimiters
	COMMA
	SEMICOLON
	COLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET
	
	// Keywords (dance-inspired)
	DANCE      // dance (variable declaration)
	SWAY       // sway (for loop)
	SPIN       // spin (function call)
	FLOW       // flow (channel declaration)
	START      // start (goroutine)
	SEND_KW    // send (channel send)
	MATCH      // match (pattern matching)
	WHEN       // when (pattern case)
	FROM       // from
	TO         // to
	IF         // if
	ELSE       // else
	RETURN     // return
	FUNCTION   // function
	TRUE       // true
	FALSE      // false
)

var keywords = map[string]TokenType{
	"dance":    DANCE,
	"sway":     SWAY,
	"spin":     SPIN,
	"flow":     FLOW,
	"start":    START,
	"send":     SEND_KW,
	"match":    MATCH,
	"when":     WHEN,
	"from":     FROM,
	"to":       TO,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"function": FUNCTION,
	"true":     TRUE,
	"false":    FALSE,
}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"
	case FLOAT:
		return "FLOAT"
	case STRING:
		return "STRING"
	case ASSIGN:
		return "="
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case ASTERISK:
		return "*"
	case SLASH:
		return "/"
	case BANG:
		return "!"
	case LT:
		return "<"
	case GT:
		return ">"
	case EQ:
		return "=="
	case NOT_EQ:
		return "!="
	case LTE:
		return "<="
	case GTE:
		return ">="
	case ARROW:
		return "->"
	case SEND:
		return "<-"
	case MATCH_OP:
		return "=~"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case COLON:
		return ":"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case LBRACKET:
		return "["
	case RBRACKET:
		return "]"
	case DANCE:
		return "dance"
	case SWAY:
		return "sway"
	case SPIN:
		return "spin"
	case FLOW:
		return "flow"
	case START:
		return "start"
	case SEND_KW:
		return "send"
	case MATCH:
		return "match"
	case WHEN:
		return "when"
	case FROM:
		return "from"
	case TO:
		return "to"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case RETURN:
		return "return"
	case FUNCTION:
		return "function"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	default:
		return "UNKNOWN"
	}
}
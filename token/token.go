package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

var keywords = map[string]TokenType {
    "fn" : FUNCTION,
    "let": LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

const (
    ILLEGAL    = "ILLEGAL"
    EOF        = "EOF"

    IDENT      = "IDENT"
    INT        = "INT"

    ASSIGN     = "="
    PLUS       = "+"

    COMMA      = ","
    SEMICOLON  = ";"

    LPAREN     = "("
    RPAREN     = ")"
    LBRACE     = "{"
    RBRACE     = "}"

    FUNCTION   = "FUNCTION"
    LET        = "LET"
)

type LetStatement struct {
    Token token.Token
    Name *Identifier
    Value Expression
}

func (ls *LetStatement) statementNode()
func (ls* LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
    Token token.TOken
    Value string
}

func (i *Identifier) expressionNode()
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

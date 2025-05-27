package compiler

type TokenType int

const (
	// Keywords
	TokenLet TokenType = iota
	TokenFn
	TokenIf
	TokenElse
	TokenWhile
	TokenReturn
	TokenStruct
	TokenImpl
	TokenTrue
	TokenFalse

	// Literals
	TokenInt
	TokenFloat
	TokenString
	TokenBool
	TokenIdent

	// Operators
	TokenPlus
	TokenMinus
	TokenStar
	TokenSlash
	TokenEqual
	TokenEqualEqual
	TokenNotEqual
	TokenLess
	TokenLessEqual
	TokenGreater
	TokenGreaterEqual

	// Delimiters
	TokenLParen
	TokenRParen
	TokenLBrace
	TokenRBrace
	TokenComma
	TokenSemicolon
	TokenDot
	TokenArrow

	// Special
	TokenEOF
	TokenError
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

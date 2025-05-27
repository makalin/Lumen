package compiler

import (
	"bufio"
	"bytes"
	"unicode"
)

type Lexer struct {
	input   *bufio.Reader
	buffer  []byte
	line    int
	column  int
	current rune
	peek    rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:  bufio.NewReader(bytes.NewReader([]byte(input))),
		buffer: []byte(input),
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	r, _, err := l.input.ReadRune()
	if err != nil {
		l.current = 0
	} else {
		l.current = r
	}
	l.column++
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.current {
	case '+':
		tok = Token{Type: TokenPlus, Literal: "+", Line: l.line, Column: l.column}
	case '-':
		if l.peek == '>' {
			l.readChar()
			tok = Token{Type: TokenArrow, Literal: "->", Line: l.line, Column: l.column}
		} else {
			tok = Token{Type: TokenMinus, Literal: "-", Line: l.line, Column: l.column}
		}
	case '*':
		tok = Token{Type: TokenStar, Literal: "*", Line: l.line, Column: l.column}
	case '/':
		tok = Token{Type: TokenSlash, Literal: "/", Line: l.line, Column: l.column}
	case '=':
		if l.peek == '=' {
			l.readChar()
			tok = Token{Type: TokenEqualEqual, Literal: "==", Line: l.line, Column: l.column}
		} else {
			tok = Token{Type: TokenEqual, Literal: "=", Line: l.line, Column: l.column}
		}
	case '(':
		tok = Token{Type: TokenLParen, Literal: "(", Line: l.line, Column: l.column}
	case ')':
		tok = Token{Type: TokenRParen, Literal: ")", Line: l.line, Column: l.column}
	case '{':
		tok = Token{Type: TokenLBrace, Literal: "{", Line: l.line, Column: l.column}
	case '}':
		tok = Token{Type: TokenRBrace, Literal: "}", Line: l.line, Column: l.column}
	case ',':
		tok = Token{Type: TokenComma, Literal: ",", Line: l.line, Column: l.column}
	case ';':
		tok = Token{Type: TokenSemicolon, Literal: ";", Line: l.line, Column: l.column}
	case '.':
		tok = Token{Type: TokenDot, Literal: ".", Line: l.line, Column: l.column}
	case 0:
		tok = Token{Type: TokenEOF, Literal: "", Line: l.line, Column: l.column}
	default:
		if isLetter(l.current) {
			tok.Literal = l.readIdentifier()
			tok.Type = l.lookupIdent(tok.Literal)
			tok.Line = l.line
			tok.Column = l.column
			return tok
		} else if isDigit(l.current) {
			tok.Type = TokenInt
			tok.Literal = l.readNumber()
			tok.Line = l.line
			tok.Column = l.column
			return tok
		} else {
			tok = Token{Type: TokenError, Literal: string(l.current), Line: l.line, Column: l.column}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.current) {
		if l.current == '\n' {
			l.line++
			l.column = 0
		}
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.column
	for isLetter(l.current) || isDigit(l.current) {
		l.readChar()
	}
	return string(l.buffer[position-1 : l.column-1])
}

func (l *Lexer) readNumber() string {
	position := l.column
	for isDigit(l.current) {
		l.readChar()
	}
	return string(l.buffer[position-1 : l.column-1])
}

func (l *Lexer) lookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"let":    TokenLet,
		"fn":     TokenFn,
		"if":     TokenIf,
		"else":   TokenElse,
		"while":  TokenWhile,
		"return": TokenReturn,
		"struct": TokenStruct,
		"impl":   TokenImpl,
		"true":   TokenBool,
		"false":  TokenBool,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TokenIdent
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

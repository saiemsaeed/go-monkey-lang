package lexer

import (
	"github.com/saiemsaeed/monkey-go/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	currLine     int
	currColumn   int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input, currLine: 1, currColumn: 1}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.EQ, string(ch)+string(l.ch), l.currLine, l.currColumn)
		} else {
			tok = newToken(token.ASSIGN, string(l.ch), l.currLine, l.currColumn)
		}
	case '+':
		tok = newToken(token.PLUS, string(l.ch), l.currLine, l.currColumn)
	case '-':
		tok = newToken(token.MINUS, string(l.ch), l.currLine, l.currColumn)
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch), l.currLine, l.currColumn)
	case '/':
		tok = newToken(token.SLASH, string(l.ch), l.currLine, l.currColumn)
	case ',':
		tok = newToken(token.COMMA, string(l.ch), l.currLine, l.currColumn)
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch), l.currLine, l.currColumn)
	case '(':
		tok = newToken(token.LPAREN, string(l.ch), l.currLine, l.currColumn)
	case ')':
		tok = newToken(token.RPAREN, string(l.ch), l.currLine, l.currColumn)
	case '{':
		tok = newToken(token.LBRACE, string(l.ch), l.currLine, l.currColumn)
	case '<':
		tok = newToken(token.LT, string(l.ch), l.currLine, l.currColumn)
	case '>':
		tok = newToken(token.GT, string(l.ch), l.currLine, l.currColumn)
	case '}':
		tok = newToken(token.RBRACE, string(l.ch), l.currLine, l.currColumn)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.NOT_EQ, string(ch)+string(l.ch), l.currLine, l.currColumn)
		} else {
			tok = newToken(token.BANG, string(l.ch), l.currLine, l.currColumn)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			tok.Line = l.currLine
			tok.Column = l.currColumn - len(tok.Literal) - 1 // Reading one extra characters in readNumber
			return tok
		} else if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Line = l.currLine
			tok.Column = l.currColumn - len(tok.Literal) - 1 // Reading one extra character in readIdentifier
			return tok
		} else {
			tok.Literal = ""
			tok.Type = token.ILLEGAL
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.currColumn = l.currColumn + 1
	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		if l.ch == '\n' {
			l.currLine++
			l.currColumn = 0
		}

		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}

	return false
}

func newToken(tokenType token.TokenType, ch string, line int, column int) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
		Line:    line,
		Column:  column - len(string(ch)),
	}
}

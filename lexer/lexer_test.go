package lexer

import (
	"testing"

	"github.com/saiemsaeed/monkey-go/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.LET, "let", 1, 1},
		{token.IDENT, "five", 1, 5},
		{token.ASSIGN, "=", 1, 10},
		{token.INT, "5", 1, 12},
		{token.SEMICOLON, ";", 1, 13},
		{token.LET, "let", 2, 0},
		{token.IDENT, "ten", 2, 4},
		{token.ASSIGN, "=", 2, 8},
		{token.INT, "10", 2, 10},
		{token.SEMICOLON, ";", 2, 12},
		{token.LET, "let", 4, 0},
		{token.IDENT, "add", 4, 4},
		{token.ASSIGN, "=", 4, 8},
		{token.FUNCTION, "fn", 4, 10},
		{token.LPAREN, "(", 4, 12},
		{token.IDENT, "x", 4, 13},
		{token.COMMA, ",", 4, 14},
		{token.IDENT, "y", 4, 16},
		{token.RPAREN, ")", 4, 17},
		{token.LBRACE, "{", 4, 19},
		{token.IDENT, "x", 5, 2},
		{token.PLUS, "+", 5, 4},
		{token.IDENT, "y", 5, 6},
		{token.SEMICOLON, ";", 5, 7},
		{token.RBRACE, "}", 6, 0},
		{token.SEMICOLON, ";", 6, 1},
		{token.LET, "let", 8, 0},
		{token.IDENT, "result", 8, 4},
		{token.ASSIGN, "=", 8, 11},
		{token.IDENT, "add", 8, 13},
		{token.LPAREN, "(", 8, 16},
		{token.IDENT, "five", 8, 17},
		{token.COMMA, ",", 8, 21},
		{token.IDENT, "ten", 8, 23},
		{token.RPAREN, ")", 8, 26},
		{token.SEMICOLON, ";", 8, 27},
		{token.BANG, "!", 9, 0},
		{token.MINUS, "-", 9, 1},
		{token.SLASH, "/", 9, 2},
		{token.ASTERISK, "*", 9, 3},
		{token.INT, "5", 9, 4},
		{token.SEMICOLON, ";", 9, 5},
		{token.INT, "5", 10, 0},
		{token.LT, "<", 10, 2},
		{token.INT, "10", 10, 4},
		{token.GT, ">", 10, 7},
		{token.INT, "5", 10, 9},
		{token.SEMICOLON, ";", 10, 10},
		{token.IF, "if", 12, 0},
		{token.LPAREN, "(", 12, 3},
		{token.INT, "5", 12, 4},
		{token.LT, "<", 12, 6},
		{token.INT, "10", 12, 8},
		{token.RPAREN, ")", 12, 10},
		{token.LBRACE, "{", 12, 12},
		{token.RETURN, "return", 13, 1},
		{token.TRUE, "true", 13, 8},
		{token.SEMICOLON, ";", 13, 12},
		{token.RBRACE, "}", 14, 0},
		{token.ELSE, "else", 14, 2},
		{token.LBRACE, "{", 14, 7},
		{token.RETURN, "return", 15, 1},
		{token.FALSE, "false", 15, 8},
		{token.SEMICOLON, ";", 15, 13},
		{token.RBRACE, "}", 16, 0},
		{token.INT, "10", 18, 0},
		{token.EQ, "==", 18, 3},
		{token.INT, "10", 18, 6},
		{token.SEMICOLON, ";", 18, 8},
		{token.INT, "10", 19, 0},
		{token.NOT_EQ, "!=", 19, 3},
		{token.INT, "9", 19, 6},
		{token.SEMICOLON, ";", 19, 7},
		{token.EOF, "", 0, 0},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line wrong. expected=%d, got=%d",
				i, tt.expectedLine, tok.Line)
		}

		if tok.Column != tt.expectedColumn {
			t.Fatalf("tests[%d] - column wrong. expected=%d, got=%d",
				i, tt.expectedColumn, tok.Column)
		}
	}
}

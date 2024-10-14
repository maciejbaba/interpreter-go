package lexer

import (
	"testing"
	"my-monkey/token"
)

type TestType struct {
	ExpectedTokenType token.TokenType
	expectedLiteral   string
}

func TestNextToken(t *testing.T) {
	input := `=+():{},;`

	tests := []TestType{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LEFT_PAREN, "("},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_BRACE, "{"},
		{token.RIGHT_BRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.END_OF_FILE, ""},
	}

	lexer := New(input)

	for index, TestType := range tests {
		token := lexer.nextToken()

		if token.TokenType != TestType.ExpectedTokenType {
			t.Fatalf("tests[%d] - TokenType wrong. expected=%q, got=%q", index, TestType.ExpectedTokenType, token.TokenType)
		}

		if token.Literal != TestType.expectedLiteral {
			t.Fatalf("tests[%d] - Token Literal wrong. expected=%q, got=%q", index, TestType.expectedLiteral, token.Literal)
		}
	}
}

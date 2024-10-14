package lexer

import (
	"my-monkey/token"
	"testing"
)

type TestType struct {
	ExpectedTokenType token.TokenType
	expectedLiteral   string
}

func TestNextToken(testing *testing.T) {
	input := `=+():{},;`

	tests := []TestType{
		{token.ASSIGN, "="},
 		{token.PLUS, "+"},
		{token.LEFT_PAREN, "("},
		{token.RIGHT_PAREN, ")"},
    {token.COLON, ":"},
 		{token.LEFT_BRACE, "{"},
		{token.RIGHT_BRACE, "}"},
 		{token.COMMA, ","},
 		{token.SEMICOLON, ";"},
		{token.END_OF_FILE, ""},
	}

	lexer := New(input)

	for index, TestType := range tests {
		token := lexer.nextToken()

		if token.Type != TestType.ExpectedTokenType {
			testing.Fatalf("tests[%d] - TokenType wrong. expected=%q, got=%q", index, TestType.ExpectedTokenType, token.Type)
		}

		if token.Literal != TestType.expectedLiteral {
			testing.Fatalf("tests[%d] - Token Literal wrong. expected=%q, got=%q", index, TestType.expectedLiteral, token.Literal)
		}
	}
}

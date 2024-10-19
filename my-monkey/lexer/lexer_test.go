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
	input := `let five = 5;
  let ten = 10;

  let add = fn(x,y) {
    x + y;
  }

  let result = add(five, ten);
  ` 

	tests := []TestType{
    {token.LET, "let"},
    {token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
    {token.INTEGER, "5"},
    {token.SEMICOLON, ";"},
    {token.LET, "let"},
    {token.IDENTIFIER, "ten"},
    {token.ASSIGN, "="},
    {token.INTEGER, "10"},
    {token.SEMICOLON, ";"},
    {token.LET, "let"},
    {token.IDENTIFIER, "add"},
    {token.ASSIGN, "="},
    {token.FUNCTION, "fn"},
    {token.LEFT_PAREN, "("},
    {token.IDENTIFIER, "x"},
    {token.COMMA, ","},
    {token.IDENTIFIER, "y"},
    {token.RIGHT_PAREN, ")"},
    {token.LEFT_BRACE, "{"},
    {token.IDENTIFIER, "x"},
    {token.PLUS, "+"},
    {token.IDENTIFIER, "y"},
    {token.SEMICOLON, ";"},
    {token.RIGHT_BRACE, "}"},
    {token.LET, "let"},
    {token.IDENTIFIER, "result"},
    {token.ASSIGN, "="},
    {token.IDENTIFIER, "add"},
    {token.LEFT_PAREN, "("},
    {token.IDENTIFIER, "five"},
    {token.COMMA, ","},
    {token.IDENTIFIER, "ten"},
    {token.RIGHT_PAREN, ")"},
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

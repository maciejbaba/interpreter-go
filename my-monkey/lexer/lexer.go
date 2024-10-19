package lexer

import (
	"my-monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readCharacter() // this will initialize all Lexer variables
	return lexer
}

func (lexer *Lexer) readCharacter() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0 // NUL - we haven't read antything or END_OF_FILE
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) nextToken() token.Token {
	var currentToken token.Token

	switch lexer.character {
	case '=':
		currentToken = newToken(token.ASSIGN, lexer.character)
	case ';':
		currentToken = newToken(token.SEMICOLON, lexer.character)
	case '(':
		currentToken = newToken(token.LEFT_PAREN, lexer.character)
	case ')':
		currentToken = newToken(token.RIGHT_PAREN, lexer.character)
	case ',':
		currentToken = newToken(token.COMMA, lexer.character)
	case '+':
		currentToken = newToken(token.PLUS, lexer.character)
	case '{':
		currentToken = newToken(token.LEFT_BRACE, lexer.character)
	case '}':
		currentToken = newToken(token.RIGHT_BRACE, lexer.character)
	case ':':
		currentToken = newToken(token.COLON, lexer.character)
	case 0:
		currentToken.Literal = ""
		currentToken.Type = token.END_OF_FILE
	default:
		if isLetter(lexer.character) {
			currentToken.Literal = lexer.readIdentifier()
			return currentToken
		} else {
			currentToken = newToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readCharacter()
	return currentToken
}

func (lexer *Lexer) readIdentifier() string {
	startIndexOfIdentifier := lexer.position
	for isLetter(lexer.character) { // read until encounter something different than letter
		lexer.readCharacter()
	}
	endIndexOfIdentifier := lexer.position // then we have the end index of indentifier
	return lexer.input[startIndexOfIdentifier:endIndexOfIdentifier]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

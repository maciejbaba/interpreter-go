package lexer

import (
  "fmt"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	return lexer
}

func (lexer *Lexer) readCharacter() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0 // NUL - we haven't read antything or END_OF_FILE
	} else {
		fmt.Printf("Lexer.character %d", lexer.character)
		fmt.Printf("Lexer.readPosition %d", lexer.readPosition)
		lexer.character = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}



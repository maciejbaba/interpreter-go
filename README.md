# Monkey Language Implementation

This repository contains implementations of the Monkey programming language, primarily following the book "Writing An Interpreter In Go" by Thorsten Ball.

## Project Structure

### `waiig_code_1.3/`
Contains the official code from the "Writing An Interpreter In Go" book (version 1.3). This includes:
- Complete chapter-by-chapter implementations
- Final working interpreter and compiler
- All test files and examples from the book

### `my-monkey/`
A custom implementation of the Monkey language lexer, built from scratch. Currently includes:
- **Token System** (`token/`): Defines token types and structures for the Monkey language
- **Lexer** (`lexer/`): Tokenizes Monkey source code into a stream of tokens
- **Tests**: Comprehensive test suite for the lexer functionality

## Getting Started

### Prerequisites
- Go 1.23.2 or later

### Running the Custom Implementation

1. Navigate to the my-monkey directory:
   ```bash
   cd my-monkey
   ```

2. Run the tests:
   ```bash
   go test ./...
   ```

3. Use the lexer in your Go code:
   ```go
   import "my-monkey/lexer"
   import "my-monkey/token"
   
   l := lexer.New("let x = 5;")
   tok := l.NextToken()
   ```

### Using the Book Code

The `waiig_code_1.3/` directory contains the complete implementations from each chapter. Each numbered directory (01/, 02/, etc.) represents the code state at the end of that chapter.

## About the Monkey Language

Monkey is a programming language designed for learning interpreter and compiler construction. It features:

- Variable bindings
- Integers and booleans
- Arithmetic expressions
- Built-in functions
- First-class and higher-order functions
- Closures
- String data structure
- Array data structure
- Hash data structure

### Example Monkey Code

```monkey
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);

let myArray = [1, 2, 3, 4, 5];
let thorsten = {"name": "Thorsten", "age": 28};

let add = fn(a, b) { return a + b; };
let factorial = fn(n) {
  if (n == 0) {
    1
  } else {
    n * factorial(n - 1)
  }
};
```

## Current Implementation Status

- ✅ **Lexer**: Complete tokenization of Monkey source code
- ⏳ **Parser**: Not yet implemented
- ⏳ **Evaluator**: Not yet implemented
- ⏳ **REPL**: Not yet implemented

## Contributing

This is a learning project. Feel free to:
- Add more token types
- Implement the parser
- Add more comprehensive tests
- Improve error handling

## Resources

- [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball
- [Writing A Compiler In Go](https://compilerbook.com/) by Thorsten Ball

## License

The code in `waiig_code_1.3/` is licensed under the MIT license as specified by the book author.
The custom implementation in `my-monkey/` is available for educational purposes.
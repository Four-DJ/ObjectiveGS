package transpiler

import (
	"strings"
)

type TokenType int

const (
	Empty TokenType = iota
	Equals
	EOL
	Namespace
	Identifier
	Class
	CurlyOpen
	CurlyClose
	Function
	BracketOpen
	BracketClose
	If
	Or
	And
	Not
	Slash
	SquareOpen
	SquareClose
	Quote
	Dot
	QuestionMark
	Colen
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) ([]Token, error) {
	tokens := []Token{}
	var textBuffer strings.Builder

	for _, char := range input {
		switch char {
		case ';', '\n':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: EOL})
		case ' ':
			tokens = textTokenizer(tokens, &textBuffer, Token{})
		case '"':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: Quote})
		case '?':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: QuestionMark})
		case '.':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: Dot})
		case ':':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: Colen})
		case '{':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: CurlyOpen})
		case '}':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: CurlyClose})
		case '(':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: BracketOpen})
		case ')':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: BracketClose})
		case '[':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: SquareOpen})
		case ']':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: SquareClose})
		case '|':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: Or})
		case '&':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: And})
		case '!':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: Not})
		case '=':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: Equals})
		case '/':
			tokens = textTokenizer(tokens, &textBuffer, Token{Type: Slash})
		default:
			textBuffer.WriteRune(char)
		}
	}

	tokens = textTokenizer(tokens, &textBuffer, Token{})
	return tokens, nil
}

func textTokenizer(tokens []Token, textBuffer *strings.Builder, token Token) []Token {
	var textToken Token
	switch textBuffer.String() {
	case "namespace":
		textToken = Token{Type: Namespace}
	case "class":
		textToken = Token{Type: Class}
	case "func", "function":
		textToken = Token{Type: Function}
	case "if":
		textToken = Token{Type: If}
	default:
		textToken = Token{Type: Identifier, Value: textBuffer.String()}
	}

	if textBuffer.Len() > 0 {
		tokens = append(tokens, textToken)
	}
	if token.Type != Empty {
		tokens = append(tokens, token)
	}
	textBuffer.Reset()
	return tokens
}

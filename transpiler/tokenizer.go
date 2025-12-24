package transpiler

import (
	"strings"
)

type TokenType int

const (
	Empty TokenType = iota
	EOF
	Namespace
	Identifier
	Class
	CurlyOpen
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
		case ';':
			token := textTokenizer(textBuffer.String())
			if token.Type != Empty {
				tokens = append(tokens, token)
				textBuffer.Reset()
			}
			tokens = append(tokens, Token{Type: EOF})
			continue
		case ' ':
			token := textTokenizer(textBuffer.String())
			tokens = append(tokens, token)
			textBuffer.Reset()
			continue
		case '{':
			tokens = append(tokens, Token{Type: CurlyOpen})
		default:
			textBuffer.WriteRune(char)
		}
	}
	return tokens, nil
}

func textTokenizer(input string) Token {
	if input == "" {
		return Token{Type: Empty}
	}
	if input == "namespace" {
		return Token{Type: Namespace}
	}
	if input == "class" {
		return Token{Type: Class}
	}
	return Token{Type: Identifier, Value: input}
}

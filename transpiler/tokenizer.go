package transpiler

import (
	"strings"
	"unicode"
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
		if unicode.IsLetter(char) {
			textBuffer.WriteRune(char)
			continue
		}
		if textBuffer.Len() > 0 {
			token := textTokenizer(textBuffer.String())
			if token.Type != Empty {
				tokens = append(tokens, token)
				textBuffer.Reset()
			}
		}
		switch char {
		case ';':
			tokens = append(tokens, Token{Type: EOF})
			continue
		case ' ':
			continue
		case '{':
			tokens = append(tokens, Token{Type: CurlyOpen})
		default:
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

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
	Function
	BracketOpen
	BracketClose
	If
	Or
	And
	Not
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
			continue
		case '(':
			tokens = append(tokens, Token{Type: BracketOpen})
			continue
		case ')':
			tokens = append(tokens, Token{Type: BracketClose})
			continue
		case '|':
			tokens = append(tokens, Token{Type: Or})
			continue
		case '&':
			tokens = append(tokens, Token{Type: And})
			continue
		case '!':
			tokens = append(tokens, Token{Type: Not})
			continue
		default:
		}
	}
	return tokens, nil
}

func textTokenizer(input string) Token {
	switch input {
	case "":
		return Token{Type: Empty}
	case "namespace":
		return Token{Type: Namespace}
	case "class":
		return Token{Type: Class}
	case "func", "function":
		return Token{Type: Function}
	case "if":
		return Token{Type: If}
	}
	return Token{Type: Identifier, Value: input}
}

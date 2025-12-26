package transpiler

import (
	"strings"
	"unicode"
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
		if unicode.IsLetter(char) || unicode.IsNumber(char) || char == '_' {
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
		case ';', '\n':
			tokens = append(tokens, Token{Type: EOL})
			continue
		case ' ':
			continue
		case '"':
			tokens = append(tokens, Token{Type: Quote})
			continue
		case '?':
			tokens = append(tokens, Token{Type: QuestionMark})
			continue
		case '.':
			tokens = append(tokens, Token{Type: Dot})
			continue
		case ':':
			tokens = append(tokens, Token{Type: Colen})
			continue
		case '{':
			tokens = append(tokens, Token{Type: CurlyOpen})
			continue
		case '}':
			tokens = append(tokens, Token{Type: CurlyClose})
			continue
		case '(':
			tokens = append(tokens, Token{Type: BracketOpen})
			continue
		case ')':
			tokens = append(tokens, Token{Type: BracketClose})
			continue
		case '[':
			tokens = append(tokens, Token{Type: SquareOpen})
			continue
		case ']':
			tokens = append(tokens, Token{Type: SquareClose})
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
		case '=':
			tokens = append(tokens, Token{Type: Equals})
			continue
		case '/':
			tokens = append(tokens, Token{Type: Slash})
			continue
		default:
		}
	}

	if textBuffer.Len() > 0 {
		token := textTokenizer(textBuffer.String())
		if token.Type != Empty {
			tokens = append(tokens, token)
			textBuffer.Reset()
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

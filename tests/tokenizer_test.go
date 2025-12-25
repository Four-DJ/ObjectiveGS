package test

import (
	"gso/transpiler"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		Case        string
		Input       []string
		Expected    []transpiler.Token
		shouldError bool
	}{
		{
			Case:  "valid Tokenization of array access",
			Input: []string{"params[0]"},
			Expected: []transpiler.Token{
				{Type: transpiler.Identifier, Value: "params"}, {Type: transpiler.SquareOpen}, {Type: transpiler.Identifier, Value: "0"}, {Type: transpiler.SquareClose},
			},
		},
		{
			Case:  "valid Tokenization on Comment",
			Input: []string{"//comment"},
			Expected: []transpiler.Token{
				{Type: transpiler.Slash}, {Type: transpiler.Slash}, {Type: transpiler.Identifier, Value: "comment"},
			},
			shouldError: false,
		},
		{
			Case:  "valid Tokenization of function call",
			Input: []string{"test();"},
			Expected: []transpiler.Token{
				{
					Type:  transpiler.Identifier,
					Value: "test",
				},
				{
					Type: transpiler.BracketOpen,
				},
				{
					Type: transpiler.BracketClose,
				},
				{
					Type: transpiler.EOL,
				},
			},
			shouldError: false,
		},
		{
			Case:  "valid Tokenization of varible declatation",
			Input: []string{"test = 0;"},
			Expected: []transpiler.Token{
				{
					Type:  transpiler.Identifier,
					Value: "test",
				},
				{
					Type: transpiler.Equals,
				},
				{
					Type:  transpiler.Identifier,
					Value: "0",
				},
				{
					Type: transpiler.EOL,
				},
			},
			shouldError: false,
		},
		{
			Case:  "valid Tokenization of semicolon",
			Input: []string{";"},
			Expected: []transpiler.Token{
				{
					Type: transpiler.EOL,
				},
			},
			shouldError: false,
		},
		{
			Case:  "valid Tokenization of namespace",
			Input: []string{"namespace Test;"},
			Expected: []transpiler.Token{
				{
					Type: transpiler.Namespace,
				},
				{
					Type:  transpiler.Identifier,
					Value: "Test",
				},
				{
					Type: transpiler.EOL,
				},
			},
			shouldError: false,
		},
		{
			Case:  "valid Tokenization of clases",
			Input: []string{"class Test {", "class Test{"},
			Expected: []transpiler.Token{
				{
					Type: transpiler.Class,
				},
				{
					Type:  transpiler.Identifier,
					Value: "Test",
				},
				{
					Type: transpiler.CurlyOpen,
				},
			},
			shouldError: false,
		},
		{
			Case:  "valid Tokenization of functions",
			Input: []string{"func test() {", "func test(){", "function test(){", "function test() {"},
			Expected: []transpiler.Token{
				{
					Type: transpiler.Function,
				},
				{
					Type:  transpiler.Identifier,
					Value: "test",
				},
				{
					Type: transpiler.BracketOpen,
				},
				{
					Type: transpiler.BracketClose,
				},
				{
					Type: transpiler.CurlyOpen,
				},
			},
			shouldError: false,
		},
		{
			Case:  "valid Tokenization of if statement",
			Input: []string{"if test || !test && test {"},
			Expected: []transpiler.Token{
				{
					Type: transpiler.If,
				},
				{
					Type:  transpiler.Identifier,
					Value: "test",
				},
				{
					Type: transpiler.Or,
				},
				{
					Type: transpiler.Or,
				},
				{
					Type: transpiler.Not,
				},
				{
					Type:  transpiler.Identifier,
					Value: "test",
				},
				{
					Type: transpiler.And,
				},
				{
					Type: transpiler.And,
				},
				{
					Type:  transpiler.Identifier,
					Value: "test",
				},
				{
					Type: transpiler.CurlyOpen,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Case, func(t *testing.T) {
			for _, input := range test.Input {
				actual, err := transpiler.Tokenize(input)
				if err != nil && !test.shouldError {
					t.Errorf("%s failed because of triggered error: \n%s", test.Case, err)
				}
				if !reflect.DeepEqual(actual, test.Expected) {
					t.Errorf("%s failed because \nexpected:\n%v\nactual:\n%v", test.Case, test.Expected, actual)
				}
			}
		})
	}
}

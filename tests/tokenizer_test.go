package test

import (
	"gso/transpiler"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		Case        string
		Input       string
		Expected    []transpiler.Token
		shouldError bool
	}{
		{
			Case:  "valid Tokenization of semicolon",
			Input: ";",
			Expected: []transpiler.Token{
				{
					Type: transpiler.EOF,
				},
			},
			shouldError: false,
		},
		{
			Case:  "valid Tokenization of namespace",
			Input: "namespace Test;",
			Expected: []transpiler.Token{
				{
					Type: transpiler.Namespace,
				},
				{
					Type:  transpiler.Identifier,
					Value: "Test",
				},
				{
					Type: transpiler.EOF,
				},
			},
			shouldError: false,
		},
	}
	for _, test := range tests {
		t.Run(test.Case, func(t *testing.T) {
			actual, err := transpiler.Tokenize(test.Input)
			if err != nil && !test.shouldError {
				t.Errorf("%s failed because of triggered error: \n%s", test.Case, err)
			}
			if !reflect.DeepEqual(actual, test.Expected) {
				t.Errorf("%s failed because \nexpected:\n%v\nactual:\n%v", test.Case, test.Expected, actual)
			}
		})
	}
}

package ast_test

import (
	"go/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Value: "myVar",
			},
			Value: &Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "myVay"},
				Value: "anotherVar",
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got = %v", program.String())
	}
}

package parser

import "testing"

import "github.com/tokenshift/scry/ast"
import "github.com/tokenshift/scry/util"
import "github.com/dustin/go-parse"

func TestArithExpr(t *testing.T) {
	in := new(parsec.StringVessel)
	in.SetInput("2 + 2")
	expr, err := ParseExpr(in)

	if !util.AssertEquals(t, nil, err) {
		return
	}

	if _, ok := expr.(ast.ArithExpr); !ok {
		t.Errorf("Expected an arithmetic expression, got %v.", expr)
		return
	}
}

package lambda2pi

import (
	"github.com/nekketsuuu/lambda2pi/parser"
	"github.com/nekketsuuu/lambda2pi/syntax"
)

func ConvertFromFile(filename string, mode EvalMode) (syntax.Pi, error) {
	l, err := parser.ParseFile(filename)
	if err != nil {
		return syntax.PNull{}, err
	}

	p, err := Convert(l, mode)
	return p, err
}

func ConvertFromString(str string, mode EvalMode) (syntax.Pi, error) {
	l, err := parser.ParseExpr(str)
	if err != nil {
		return syntax.PNull{}, err
	}

	p, err := Convert(l, mode)
	return p, err
}

package parser

import (
	"errors"
	"io/ioutil"

	"github.com/nekketsuuu/lambda2pi/lib/syntax"
)

// ParseFile parses single file as a lambda term, and return the AST.
//
// Note that this function is different from go/parser.ParseFile.
// ParseFile converts the content of the file into the AST
//
// Note that the interface of this function is different from go/parser.ParseFile
//
func ParseFile(filename string) (syntax.Lambda, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return syntax.LVar{Name: "error"}, err
	}
	l, err := parseBytes(b)
	return l, err
}

// ParseExpr returns the lambda AST of the argument.
func ParseExpr(e string) (syntax.Lambda, error) {
	l, err := parseBytes([]byte(e))
	return l, err
}

func parseBytes(b []byte) (syntax.Lambda, error) {
	lexer := yyLex{line: b, err: nil}

	// yyParse returns 0 if succeed
	if yyParse(&lexer) != 0 {
		if lexer.err == nil {
			lexer.err = errors.New("parser: goyacc doesn't set the error value")
		}
		return syntax.LVar{Name: "error"}, lexer.err
	}
	return lexer.term, nil
}

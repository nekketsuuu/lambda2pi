package parser

import (
	"errors"
	"io/ioutil"

	"github.com/nekketsuuu/lambda2pi"
)

// ParseFile parses single file as a lambda term, and return the AST.
//
// Note that this function is different from go/parser.ParseFile.
// ParseFile converts the content of the file into the AST
//
// Note that the interface of this function is different from go/parser.ParseFile
//
func ParseFile(filename string) (lambda2pi.Lambda, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return lambda2pi.LVar{Name: "error"}, err
	}
	l, err := parseBytes(b)
	return l, err
}

// ParseExpr returns the lambda AST of the argument.
func ParseExpr(e string) (lambda2pi.Lambda, error) {
	l, err := parseBytes([]byte(e))
	return l, err
}

func parseBytes(b []byte) (lambda2pi.Lambda, error) {
	lexer := yyLex{line: b, err: nil}

	// yyParse returns 0 if succeed
	if yyParse(&lexer) != 0 {
		if lexer.err != nil {
			lexer.err = errors.New("parser: goyacc doesn't set the error value")
		}
		return lambda2pi.LVar{Name: "error"}, lexer.err
	}
	return lexer.term, nil
}

package parser

import (
	"errors"

	"github.com/nekketsuuu/lambda2pi"
)

// ParseFile parses single file as a lambda term, and return the AST.
//
// Note that this function is different from go/parser.ParseFile.
//
/*
func ParseFile(filename string) (lambda2pi.Lambda, error) {

}
*/

// ParseExpr returns the lambda AST of the argument.
func ParseExpr(e string) (lambda2pi.Lambda, error) {
	line := []byte(e)
	lexer := yyLex{line: line, err: nil}

	// yyParse returns 0 if succeed
	if yyParse(&lexer) != 0 {
		if lexer.err != nil {
			lexer.err = errors.New("parser: goyacc doesn't set the error value")
		}
		return lambda2pi.LVar{Name: "error"}, lexer.err
	}
	return lexer.term, nil
}

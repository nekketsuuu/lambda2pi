package parser

import "testing"

// lexAll consumes all bytes in lex, and outputs the result of lexing.
func lexAll(lex yyLexer) []int {
	tokens := []int{}
	var lval yySymType
LexingLoop:
	for {
		t := lex.Lex(&lval)
		if t == eof {
			break LexingLoop
		}
		tokens = append(tokens, t)
	}
	return tokens
}

// toToken converts a character value defined in parser.go into a token name.
func toToken(i int) string {
	switch i {
	case LPAR:
		return "LPAR"
	case RPAR:
		return "RPAR"
	case LAMBDA:
		return "LAMBDA"
	case DOT:
		return "DOT"
	case IDENT:
		return "IDENT"
	case eof:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}

// toTokens converts a slice of character values into token names.
func toTokens(is []int) []string {
	ts := make([]string, len(is))
	for i, v := range is {
		ts[i] = toToken(v)
	}
	return ts
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

// tests of some small examples
func TestLexExpr(t *testing.T) {
	positiveExamples := map[string][]int{
		"x":                {IDENT},
		"xyz":              {IDENT},
		"x2":               {IDENT},
		"\\x. x":           {LAMBDA, IDENT, DOT, IDENT},
		"(x)":              {LPAR, IDENT, RPAR},
		"x y":              {IDENT, IDENT},
		"(\\x. \\y. x y)z": {LPAR, LAMBDA, IDENT, DOT, LAMBDA, IDENT, DOT, IDENT, IDENT, RPAR, IDENT},
	}

	for expr, want := range positiveExamples {
		line := []byte(expr)
		lexer := yyLex{line: line, err: nil}
		result := lexAll(&lexer)
		if !equal(result, want) {
			rt := toTokens(result)
			wt := toTokens(want)
			t.Errorf("lexAll(\"%s\") is %v, want %v", expr, rt, wt)
		}
	}
}

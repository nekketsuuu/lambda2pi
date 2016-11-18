package parser

import (
	"path/filepath"
	"testing"

	"github.com/nekketsuuu/lambda2pi/lib/syntax"
)

// tests of some small examples
func TestParseExpr(t *testing.T) {
	// a variable
	e := "x"
	x, err := ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(syntax.LVar); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want syntax.LVar", e, x)
	}

	// a lambda abstraction
	e = "\\x. x"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(syntax.LAbs); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want syntax.LAbs", e, x)
	}

	// parenthesis
	e = "(x)"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(syntax.LVar); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want syntax.LVar", e, x)
	}

	// application 1
	e = "x y"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(syntax.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want syntax.LApp", e, x)
	}

	// application 2
	e = "(\\x. x) y"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(syntax.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want syntax.LApp", e, x)
	}
	if _, ok := x.(syntax.LApp).First.(syntax.LAbs); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of (\\x. x), want syntax.LAbs", e, x)
	}

	// application 3
	e = "(\\x. \\y. x y) y z"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	var xx syntax.Lambda
	var ok bool
	if xx, ok = x.(syntax.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want syntax.LApp", e, x)
	}
	if _, ok = xx.(syntax.LApp).Second.(syntax.LVar); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"z\", want syntax.LVar", e, x)
	}
	if xx, ok = xx.(syntax.LApp).First.(syntax.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"%v\", want syntax.LApp", e, x, xx.String())
	}
	if xx, ok = xx.(syntax.LApp).First.(syntax.LAbs); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"%v\", want syntax.LAbs", e, x, xx.String())
	}
	if xx, ok = xx.(syntax.LAbs).Body.(syntax.LAbs); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"%v\", want syntax.LAbs", e, x, xx.String())
	}
	if _, ok = xx.(syntax.LAbs).Body.(syntax.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"%v\", want syntax.LApp", e, x, xx.String())
	}
}

var files = map[string]bool{
	"simple.lambda":     true,
	"simple2.lambda":    true,
	"simple3.lambda":    true,
	"simple4.lambda":    true,
	"illParen.lambda":   false,
	"nilParen.lambda":   false,
	"nobodyAbst.lambda": false,
}

// tests for some big examples
func TestParseFiles(t *testing.T) {
	for filename, success := range files {
		path := filepath.Join("testdata", filename)
		_, err := ParseFile(path)
		if (err == nil) != success {
			if success {
				t.Errorf("An error occured while parsing a file \"%v\". The error message is: %v", path, err)
			} else {
				t.Errorf("It should fail while parsing a file \"%v\", but it succeed.", path)
			}
		}
	}
}

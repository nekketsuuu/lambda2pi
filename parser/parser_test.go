package parser

import (
	"path/filepath"
	"testing"

	"github.com/nekketsuuu/lambda2pi"
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
	if _, ok := x.(lambda2pi.LVar); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want lambda2pi.LVar", e, x)
	}

	// a lambda abstraction
	e = "\\x. x"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(lambda2pi.LAbs); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want lambda2pi.LAbs", e, x)
	}

	// parenthesis
	e = "(x)"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(lambda2pi.LVar); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want lambda2pi.LVar", e, x)
	}

	// application 1
	e = "x y"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(lambda2pi.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want lambda2pi.LApp", e, x)
	}

	// application 2
	e = "(\\x. x) y"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	if _, ok := x.(lambda2pi.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want lambda2pi.LApp", e, x)
	}
	if _, ok := x.(lambda2pi.LApp).First.(lambda2pi.LAbs); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of (\\x. x), want lambda2pi.LAbs", e, x)
	}

	// application 3
	e = "(\\x. \\y. x y) y z"
	x, err = ParseExpr(e)
	if err != nil {
		t.Errorf("ParseExpr(%q): %v", e, err)
	}
	// sanity check
	var xx lambda2pi.Lambda
	var ok bool
	if xx, ok = x.(lambda2pi.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T, want lambda2pi.LApp", e, x)
	}
	if _, ok = xx.(lambda2pi.LApp).Second.(lambda2pi.LVar); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"z\", want lambda2pi.LVar", e, x)
	}
	if xx, ok = xx.(lambda2pi.LApp).First.(lambda2pi.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"%v\", want lambda2pi.LApp", e, x, xx.String())
	}
	if xx, ok = xx.(lambda2pi.LApp).First.(lambda2pi.LAbs); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"%v\", want lambda2pi.LAbs", e, x, xx.String())
	}
	if xx, ok = xx.(lambda2pi.LAbs).Body.(lambda2pi.LAbs); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"%v\", want lambda2pi.LAbs", e, x, xx.String())
	}
	if _, ok = xx.(lambda2pi.LAbs).Body.(lambda2pi.LApp); !ok {
		t.Errorf("ParseExpr(\"%v\"): got %T for the type of \"%v\", want lambda2pi.LApp", e, x, xx.String())
	}
}

var files = map[string]bool{
	"simple.lambda":   true,
	"simple2.lambda":  true,
	"simple3.lambda":  true,
	"simple4.lambda":  true,
	"illParen.lambda": false,
	"nilParen.lambda": false,
}

// tests for some big examples
func TestParseFiles(t *testing.T) {
	for filename, success := range files {
		path := filepath.Join("testdata", filename)
		_, err := ParseFile(path)
		if (err == nil) != success {
			if success {
				t.Errorf("An error occured while parsing a file \"%v\": %v", path, err)
			} else {
				t.Errorf("It should fail while parsing a file \"%v\", but it succeed.", path)
			}
		}
	}
}

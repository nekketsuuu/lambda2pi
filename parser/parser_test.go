package parser

import (
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
}

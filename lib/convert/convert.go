// The package convert provides the function Convert, which converts the lambda term into the pi term
//
// Each function Convert, ConvertFromFiles and ConvertFromString needs a mode for an argument. Only CallByValue mode is implemented currently.
//
// TODO(nekketsuuu): Avoid name collision
//
// TODO(nekketsuuu): Implement CallByName mode
//
package convert

import (
	"errors"
	"fmt"
	"strconv"

	. "github.com/nekketsuuu/lambda2pi/lib/syntax"
)

// Convert a lambda term into a pi term.
func Convert(l Lambda, mode EvalMode) (Pi, error) {
	// Forbid lambda terms using the name for conversion (super sanuki)
	if !isValidInput(l) {
		return PNull{}, errors.New("convert: Current implementation doesn't support a lambda term containing pp, qq, rr, y.+, zz, ww as the variable name")
	}

	// Convert
	switch mode {
	case CallByValue:
		i := 0
		p, err := convertAsCbV(l, "pp", &i)
		return p, err
	case CallByName:
		// not implemented
		return PNull{}, errors.New("convert: CallByName has not implemented yet")
		// return convertAsCbN(l)
	default:
		return PNull{}, errors.New("convert: Unknown mode")
	}
}

func isValidInput(l Lambda) bool {
	switch l.(type) {
	case LVar:
		if !isValidName(l.(LVar).Name) {
			return false
		}
		return true
	case LAbs:
		if !isValidName(l.(LAbs).Var) {
			return false
		}
		return isValidInput(l.(LAbs).Body)
	case LApp:
		return isValidInput(l.(LApp).First) && isValidInput(l.(LApp).Second)
	default:
		return false
	}
}

func isValidName(id LambdaIdent) bool {
	if id == "pp" || id == "qq" || id == "rr" || id == "zz" || id == "ww" {
		return false
	}
	// real tenuki
	if string(id)[0] == 'y' && len(string(id)) > 1 {
		return false
	}
	return true
}

func convertAsCbV(l Lambda, p PiIdent, index *int) (Pi, error) {
	switch l.(type) {
	case LambdaValue:
		// [[ V ]]p = p?y.[[ y := V ]] (y not free in V)
		y := ToPiIdent("yy" + strconv.Itoa(*index)) // new name
		(*index)++
		vp, err := convertSbst(y, l.(LambdaValue), index)
		return POut{
			Channel: p,
			Value:   y,
			Body:    vp,
		}, err
	case LApp:
		// [[ M N ]]p = new q in new r in (ap(p, q, r) | [[ M ]]q | [[ N ]]r)
		//   where ap(p, q, r) = q?y.new v in y!v.r?z.v!z.v!p
		mp, err := convertAsCbV(l.(LApp).First, "qq", index)
		if err != nil {
			return PNull{}, err
		}
		np, err := convertAsCbV(l.(LApp).Second, "rr", index)
		if err != nil {
			return PNull{}, err
		}
		return PNew{
			Name: "qq",
			Body: PNew{
				Name: "rr",
				Body: PPar{
					First: ap(p, "qq", "rr"),
					Second: PPar{
						First:  mp,
						Second: np,
					},
				},
			},
		}, nil
	default:
		return PNull{}, errors.New(fmt.Sprintf("convert: Unknown type %T. This error can't be occured", l))
	}
}

// ap(p, q, r) = q?y.new v in y!v.r?z.v!z.v!p
func ap(p PiIdent, q PiIdent, r PiIdent) Pi {
	return PIn{
		Channel: q,
		Value:   "yy",
		Body: PNew{
			Name: "vv",
			Body: POut{
				Channel: "yy",
				Value:   "vv",
				Body: PIn{
					Channel: r,
					Value:   "zz",
					Body: POut{
						Channel: "vv",
						Value:   "zz",
						Body: POut{
							Channel: "vv",
							Value:   "pp",
							Body:    PNull{},
						},
					},
				},
			},
		},
	}
}

func convertSbst(y PiIdent, v LambdaValue, index *int) (Pi, error) {
	switch v.(type) {
	case LVar:
		// [[ y := x ]] = *y?w.x!w
		return PRep{
			Body: PIn{
				Channel: y,
				Value:   "ww",
				Body: POut{
					Channel: ToPiIdent(v.(LVar).Name),
					Value:   "ww",
					Body:    PNull{},
				},
			},
		}, nil
	case LAbs:
		// [[ y := \x. M ]] = *y?w.w?x.new p in [[ M ]]p
		mp, err := convertAsCbV(v.(LAbs).Body, "pp", index)
		return PRep{
			Body: PIn{
				Channel: y,
				Value:   "ww",
				Body: PIn{
					Channel: "ww",
					Value:   ToPiIdent(v.(LAbs).Var),
					Body: PNew{
						Name: "pp",
						Body: mp,
					},
				},
			},
		}, err
	default:
		return PNull{}, errors.New(fmt.Sprintf("convert: Unknown type %T. This error can't be occured", v))
	}
}

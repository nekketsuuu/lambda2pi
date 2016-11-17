package lambda2pi

import (
	"errors"
	"fmt"
	"strconv"
)

// Convert a lambda term into a pi term.
func Convert(l Lambda, mode EvalMode) (Pi, error) {
	switch mode {
	case CallByValue:
		i := 0
		p, err := convertAsCbV(l, "p", &i)
		return p, err
	case CallByName:
		// not implemented
		return PNull{}, errors.New("Not implemented")
		// return convertAsCbN(l)
	default:
		return PNull{}, errors.New("Unknown mode")
	}
}

func convertAsCbV(l Lambda, p PiIdent, index *int) (Pi, error) {
	switch l.(type) {
	case lambdaValue:
		// [[ V ]]p = p?y.[[ y := V ]] (y not free in V)
		y := toPiIdent("y" + strconv.Itoa(*index)) // new name
		(*index)++
		vp, err := convertSbst(y, l.(lambdaValue), index)
		return POut{
			Channel: p,
			Value:   y,
			Body:    vp,
		}, err
	case LApp:
		// [[ M N ]]p = new q in new r in (ap(p, q, r) | [[ M ]]q | [[ N ]]r)
		//   where ap(p, q, r) = q?y.new v in y!v.r?z.v!z.v!p
		mp, err := convertAsCbV(l.(LApp).First, "q", index)
		if err != nil {
			return PNull{}, err
		}
		np, err := convertAsCbV(l.(LApp).Second, "r", index)
		if err != nil {
			return PNull{}, err
		}
		return PNew{
			Name: "q",
			Body: PNew{
				Name: "r",
				Body: PPar{
					First: ap(p, "q", "r"),
					Second: PPar{
						First:  mp,
						Second: np,
					},
				},
			},
		}, nil
	default:
		return PNull{}, errors.New(fmt.Sprintf("Unknown type %T. This error can't be occured", l))
	}
}

// ap(p, q, r) = q?y.new v in y!v.r?z.v!z.v!p
func ap(p PiIdent, q PiIdent, r PiIdent) Pi {
	return PIn{
		Channel: q,
		Value:   "y",
		Body: PNew{
			Name: "v",
			Body: POut{
				Channel: "y",
				Value:   "v",
				Body: PIn{
					Channel: r,
					Value:   "z",
					Body: POut{
						Channel: "v",
						Value:   "z",
						Body: POut{
							Channel: "v",
							Value:   "p",
							Body:    PNull{},
						},
					},
				},
			},
		},
	}
}

func convertSbst(y PiIdent, v lambdaValue, index *int) (Pi, error) {
	switch v.(type) {
	case LVar:
		// [[ y := x ]] = *y?w.x!w
		return PRep{
			Body: PIn{
				Channel: y,
				Value:   "w",
				Body: POut{
					Channel: convertIdent(v.(LVar).Name),
					Value:   "w",
					Body:    PNull{},
				},
			},
		}, nil
	case LAbs:
		// [[ y := \x. M ]] = *y?w.w?w.new p in [[ M ]]p
		mp, err := convertAsCbV(v.(LAbs).Body, "p", index)
		return PRep{
			Body: PIn{
				Channel: y,
				Value:   "w",
				Body: PIn{
					Channel: "w",
					Value:   convertIdent(v.(LAbs).Var),
					Body: PNew{
						Name: "p",
						Body: mp,
					},
				},
			},
		}, err
	default:
		return PNull{}, errors.New(fmt.Sprintf("Unknown type %T. This error can't be occured", v))
	}
}

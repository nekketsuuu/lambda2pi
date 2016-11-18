package convert

import "errors"

type EvalMode int

const (
	CallByValue = iota
	CallByName
)

func ToMode(s string) (EvalMode, error) {
	switch s {
	case "CallByValue":
		return CallByValue, nil
	case "CallByName":
		return CallByName, nil
	default:
		return CallByValue, errors.New("convert: Unknown EvalMode \"" + s + "\"")
	}
}

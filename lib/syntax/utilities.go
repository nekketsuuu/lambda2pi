// This package provides the types for the syntax
// of the lambda calculus and the pi calculus.
//
package syntax

// TODO(nekketsuuu): return error value
func ToLambdaIdent(id interface{}) LambdaIdent {
	switch id.(type) {
	case string:
		return LambdaIdent(id.(string))
	case PiIdent:
		return LambdaIdent(id.(PiIdent))
	default:
		// error
		return LambdaIdent("")
	}
}

// TODO(nekketsuuu): return error value
func ToPiIdent(id interface{}) PiIdent {
	switch id.(type) {
	case string:
		return PiIdent(id.(string))
	case LambdaIdent:
		return PiIdent(id.(LambdaIdent))
	default:
		// error
		return PiIdent("")
	}
}

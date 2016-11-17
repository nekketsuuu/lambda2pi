package lambda2pi

func convertIdent(id LambdaIdent) PiIdent {
	return PiIdent(id)
}

func toLambdaIdent(s string) LambdaIdent {
	return LambdaIdent(s)
}

func toPiIdent(s string) PiIdent {
	return PiIdent(s)
}

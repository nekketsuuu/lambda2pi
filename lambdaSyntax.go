package lambda2pi

// A type of identifiers
type LambdaIdent string

// A type of lambda terms
type Lambda interface {
	lambda()
}

type (
	// A type for variables x
	LVar struct {
		Name LambdaIdent
	}

	// A type for the abstraction (\x. M)
	LAbs struct {
		Var  LambdaIdent
		Body Lambda
	}

	// A type for the application (M N)
	LApp struct {
		First  Lambda
		Second Lambda
	}
)

func (t LVar) lambda() {}
func (t LAbs) lambda() {}
func (t LApp) lambda() {}

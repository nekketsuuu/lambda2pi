package lambda2pi

// A type of identifiers
type LambdaIdent string

func (id LambdaIdent) String() string {
	return string(id)
}

// A type of lambda terms
type Lambda interface {
	String() string
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

// For config: lambda letter for abstractions
const lambda string = "\\"

func (t LVar) String() string {
	return t.Name.String()
}
func (t LAbs) String() string {
	return "(" + lambda + t.Var.String() + ". " + t.Body.String() + ")"
}
func (t LApp) String() string {
	return "(" + t.First.String() + " " + t.Second.String() + ")"
}

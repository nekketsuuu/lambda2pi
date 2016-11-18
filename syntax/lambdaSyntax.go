package syntax

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
	// A variable: x
	LVar struct {
		Name LambdaIdent
	}

	// A lambda abstraction: (\x. M)
	LAbs struct {
		Var  LambdaIdent
		Body Lambda
	}

	// An application: (M N)
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

// A type of terms treated as values in CbV lambda calculus
type LambdaValue interface {
	String() string
	lambdaValue()
}

func (t LVar) lambdaValue() {}
func (t LAbs) lambdaValue() {}

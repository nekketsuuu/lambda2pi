package lambda2pi

// A type for identifiers
type PiIdent string

func (id PiIdent) String() string {
	return string(id)
}

// A type for pi terms
type Pi interface {
	String() string
	pi()
}

type (
	// The terminate process: O
	PNull struct{}

	// A variable: x
	PVar struct {
		Name PiIdent
	}

	// An input guard: x?y.P
	PIn struct {
		Channel PiIdent
		Value   PiIdent
		Body    Pi
	}

	// An output guard: x!y.P
	POut struct {
		Channel PiIdent
		Value   PiIdent
		Body    Pi
	}

	// A parallel composition: P|Q
	PPar struct {
		First  Pi
		Second Pi
	}

	// A replication: *P
	PRep struct {
		Body Pi
	}

	// A name restriction: new x in P
	PNew struct {
		Name PiIdent
		Body Pi
	}
)

func (PNull) pi() {}
func (PVar) pi()  {}
func (PIn) pi()   {}
func (POut) pi()  {}
func (PPar) pi()  {}
func (PRep) pi()  {}
func (PNew) pi()  {}

func (p PNull) String() string {
	return "O"
}
func (p PVar) String() string {
	return p.Name.String()
}
func (p PIn) String() string {
	return p.Channel.String() + "?" + p.Value.String() + "." + p.Body.String()
}
func (p POut) String() string {
	return p.Channel.String() + "!" + p.Value.String() + "." + p.Body.String()
}
func (p PPar) String() string {
	return "(" + p.First.String() + " | " + p.Second.String() + ")"
}
func (p PRep) String() string {
	return "(*" + p.Body.String() + ")"
}
func (p PNew) String() string {
	return "(new " + p.Name.String() + " in " + p.Body.String() + ")"
}

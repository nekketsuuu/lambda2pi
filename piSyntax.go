package lambda2pi

// A type for identifiers
type PiIdent string

// A type for pi terms
type Pi interface {
	pi()
}

type (
	// the terminate process: O
	PNull struct{}

	// a variable: x
	PVar struct {
		Name PiIdent
	}

	// an input guard: x?y.P
	PIn struct {
		Channel PiIdent
		Value   PiIdent
		Body    Pi
	}

	// an output guard: x!y.P
	POut struct {
		Channel PiIdent
		Value   PiIdent
		Body    Pi
	}

	// a parallel composition: P|Q
	PPar struct {
		First  Pi
		Second Pi
	}

	// a replication: *P
	PRep struct {
		Body Pi
	}

	// a name restriction: new x in P
	PNew struct {
		Name PiIdent
		Body Pi
	}
)

func (PNull) Pi() {}
func (PVar) Pi()  {}
func (PIn) Pi()   {}
func (POut) Pi()  {}
func (PPar) Pi()  {}
func (PRep) Pi()  {}
func (PNew) Pi()  {}

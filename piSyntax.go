package lambda2pi

// A type for identifiers
type PiIdent string

// A type for pi terms
type Pi interface {
	pi()
}

type (
	// O
	PNull struct{}

	// x
	PVar struct {
		Name PiIdent
	}

	// x?y.P
	PIn struct {
		Channel PiIdent
		Value   PiIdent
		Body    Pi
	}

	// x!y.P
	POut struct {
		Channel PiIdent
		Value   PiIdent
		Body    Pi
	}

	// P|Q
	PPar struct {
		First  Pi
		Second Pi
	}

	// *P
	PRep struct {
		Body Pi
	}

	// new x in P
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

package rt

import (
	"github.com/ionous/sashimi/util/ident"
)

// Reference provides an object pointer; the closest thing to an object literal.
type Reference ident.Id

// GetObject implements ObjEval allowing reference literals.
func (xr Reference) GetObject(run Runtime) (Object, error) {
	return run.GetObject(xr.Id())
}

func (xr Reference) Id() ident.Id {
	return ident.Id(xr)
}

func (xr Reference) String() string {
	return ident.Id(xr).String()
}

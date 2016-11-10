package rt

import (
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

// Reference provides an object pointer; the closest thing to an object literal.
type Reference ident.Id

// GetObject implements ObjEval allowing reference literals.
func (xr Reference) GetObject(run Runtime) (ret Object, err error) {
	if id := ident.Id(xr); id.Empty() {
		ret = Object{}
	} else if inst, ok := run.GetInstance(id); !ok {
		err = errutil.New("runtime.GetObject not found", id)
	} else {
		ret = Object{inst}
	}
	return
}

func (xr Reference) Id() ident.Id {
	return ident.Id(xr)
}

func (xr Reference) String() string {
	return ident.Id(xr).String()
}

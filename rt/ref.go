package rt

import (
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

// Reference provides an object pointer; the closest thing to an object literal.
type Reference struct {
	Value ident.Id
}

// GetObject implements ObjEval allowing reference literals.
func (xr Reference) GetObject(run Runtime) (ret Object, err error) {
	if id := xr.Value; id.Empty() {
		ret = Object{}
	} else if inst, ok := run.GetInstance(id); !ok {
		err = errutil.New("runtime.GetObject not found", id)
	} else {
		ret = Object{inst}
	}
	return
}

func (xr Reference) Id() ident.Id {
	return xr.Value
}

func (xr Reference) String() string {
	return string(xr.Value)
}

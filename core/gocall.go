package core

import (
	"github.com/ionous/mars/rt"
)

type GoCall struct {
	Action     Property
	Parameters []rt.ParameterSource
}

func (gc GoCall) Execute(r rt.Runtime) (err error) {
	if ref, e := gc.Action.Ref.GetReference(r); e != nil {
		err = e
	} else if obj, e := r.GetObject(ref); e != nil {
		err = e
	} else {
		// FIX: how much of looping, etc. do you want to leak in?
		// maybe none; except for a very special "partials"?
		if e := r.RunAction(string(gc.Action.Field), ObjectScope{obj}, gc.Parameters); e != nil {
			err = e
		}
	}
	return
}

package core

import (
	"github.com/ionous/mars/rt"
)

type GoCall struct {
	Action     Property
	Parameters []rt.ParameterSource
}

func (gc GoCall) Execute(run rt.Runtime) (err error) {
	panic("stuff")
	// if obj, e := gc.Action.Ref.GetObject(run); e != nil {
	// 	err = e
	// } else {
	// 	// FIX: how much of looping, etc. do you want to leak in?
	// 	// maybe none; except for a very special "partials"?
	// 	if e := run.RunAction(string(gc.Action.Field), ObjectScope{obj}, gc.Parameters); e != nil {
	// 		err = e
	// 	}
	// }
	// return
}

package inbuilt

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/rt"
)

// Inbuilt contains mars runtime interfaces.
func Inbuilt() *mars.Package {
	if inbuilt == nil {
		inbuilt = &mars.Package{
			Name:       "Inbuilt",
			Interfaces: (*InbuiltInterfaces)(nil),
		}
	}
	return inbuilt
}

var inbuilt *mars.Package

type InbuiltInterfaces struct {
	rt.Execute
	rt.BoolEval
	rt.NumberEval
	rt.TextEval
	rt.ObjEval
	rt.StateEval
	rt.NumListEval
	rt.TextListEval
	rt.ObjListEval
}

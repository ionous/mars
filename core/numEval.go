package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/errutil"
)

type AddNum struct {
	Augend, Addend rt.NumberEval
}

func (op AddNum) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if a, e := op.Augend.GetNumber(run); e != nil {
		err = errutil.New("add augend get", e)
	} else if b, e := op.Addend.GetNumber(run); e != nil {
		err = errutil.New("add addend get", e)
	} else {
		ret = rt.Number{a.Value + b.Value}
	}
	return
}

type Inc struct {
	Field types.NamedProperty
}

func (op Inc) Execute(run rt.Runtime) (err error) {
	ref := GetObj{"@"}
	if n, e := (PropertyNum{op.Field, ref}).GetNumber(run); e != nil {
		err = errutil.New("inc property get", e)
	} else {
		n := rt.Number{n.Value + 1}
		if e := (Property{op.Field, ref}).SetGeneric(run, n); e != nil {
			err = errutil.New("inc property set", e)
		}
	}
	return
}

package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

type AddNum struct {
	Augend, Addend rt.NumEval
}

func (op AddNum) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if a, e := op.Augend.GetNumber(run); e != nil {
		err = errutil.New("add augend get", e)
	} else if b, e := op.Addend.GetNumber(run); e != nil {
		err = errutil.New("add addend get", e)
	} else {
		ret = rt.Number(a.Float() + b.Float())
	}
	return
}

type Inc struct {
	Name string
}

func (op Inc) Execute(run rt.Runtime) (err error) {
	ref := GetObject{}
	if v, e := (PropertyNum{op.Name, ref}).GetNumber(run); e != nil {
		err = errutil.New("inc property get", e)
	} else {
		var n rt.Number = v + 1
		if e := (Property{op.Name, ref}).SetGeneric(run, n); e != nil {
			err = errutil.New("inc property set", e)
		}
	}
	return
}

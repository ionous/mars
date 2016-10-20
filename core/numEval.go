package core

import (
	"github.com/ionous/mars/rt"
)

type AddNum struct {
	Augend, Addend rt.NumEval
}

func (op AddNum) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if a, e := op.Augend.GetNumber(run); e != nil {
		err = e
	} else if b, e := op.Addend.GetNumber(run); e != nil {
		err = e
	} else {
		ret = rt.Number(a.Float() + b.Float())
	}
	return
}

package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

type CallWithNum struct {
	Num rt.NumEval
}

func (p CallWithNum) Resolve(run rt.Runtime) (ret meta.Generic, err error) {
	if v, e := p.Num.GetNumber(run); e != nil {
		err = e
	} else {
		ret = rt.NumEval(v)
	}
	return
}

type CallWithText struct {
	Text rt.TextEval
}

func (p CallWithText) Resolve(run rt.Runtime) (ret meta.Generic, err error) {
	if v, e := p.Text.GetText(run); e != nil {
		err = e
	} else {
		ret = rt.TextEval(v)
	}
	return
}

type CallWithRef struct {
	Ref rt.ObjEval
}

func (p CallWithRef) Resolve(run rt.Runtime) (ret meta.Generic, err error) {
	if v, e := p.Ref.GetObject(run); e != nil {
		err = e
	} else {
		ret = rt.ObjEval(v)
	}
	return
}

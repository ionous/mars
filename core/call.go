package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

type CallWithNum struct {
	Num rt.NumEval
}

func (p CallWithNum) Resolve(r rt.Runtime) (ret meta.Generic, err error) {
	if v, e := p.Num.GetNumber(r); e != nil {
		err = e
	} else {
		ret = rt.NumEval(v)
	}
	return
}

type CallWithText struct {
	Text rt.TextEval
}

func (p CallWithText) Resolve(r rt.Runtime) (ret meta.Generic, err error) {
	if v, e := p.Text.GetText(r); e != nil {
		err = e
	} else {
		ret = rt.TextEval(v)
	}
	return
}

type CallWithRef struct {
	Ref rt.RefEval
}

func (p CallWithRef) Resolve(r rt.Runtime) (ret meta.Generic, err error) {
	if v, e := p.Ref.GetReference(r); e != nil {
		err = e
	} else {
		ret = rt.RefEval(v)
	}
	return
}

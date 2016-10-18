package core

import (
	"github.com/ionous/mars/rt"
)

type CallWithNum struct {
	Num rt.NumEval
}

func (p CallWithNum) Resolve(r rt.Runtime) (ret rt.Value, err error) {
	return p.Num.GetNumber(r)
}

type CallWithText struct {
	Text rt.TextEval
}

func (p CallWithText) Resolve(r rt.Runtime) (ret rt.Value, err error) {
	return p.Text.GetText(r)
}

type CallWithRef struct {
	Ref rt.RefEval
}

func (p CallWithRef) Resolve(r rt.Runtime) (ret rt.Value, err error) {
	return p.Ref.GetReference(r)
}

package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

type SetNum struct {
	Tgt Property
	Num rt.NumEval
}

type SetTxt struct {
	Tgt Property
	Txt rt.TextEval
}

type SetRef struct {
	Tgt Property
	Ref rt.RefEval
}

type ChangeState struct {
	Ref    rt.RefEval
	States []ident.Id
}

func Change(tgt rt.RefEval) ChangeState {
	return ChangeState{tgt, nil}
}

func (x SetNum) Execute(r rt.Runtime) (err error) {
	if n, e := x.Num.GetNumber(r); e != nil {
		err = e
	} else if v, e := x.Tgt.value(r); e != nil {
		err = e
	} else {
		err = v.SetNum(n.Float())
	}
	return
}

func (x SetTxt) Execute(r rt.Runtime) (err error) {
	if t, e := x.Txt.GetText(r); e != nil {
		err = e
	} else if v, e := x.Tgt.value(r); e != nil {
		err = e
	} else {
		err = v.SetText(t.String())
	}
	return
}

func (x SetRef) Execute(r rt.Runtime) (err error) {
	if ref, e := x.Ref.GetReference(r); e != nil {
		err = e
	} else if v, e := x.Tgt.value(r); e != nil {
		err = e
	} else {
		err = v.SetObject(ref.Id())
	}
	return
}

func (p ChangeState) To(state string) ChangeState {
	return p.And(state)
}

func (p ChangeState) And(state string) ChangeState {
	p.States = append(p.States, MakeStringId(state))
	return p
}

// func (oa *GameObject) IsNow(state string) {
func (x ChangeState) Execute(r rt.Runtime) (err error) {
	if ref, e := x.Ref.GetReference(r); e != nil {
		err = e
	} else if o, e := MakeObject(r, ref); e != nil {
		err = e
	} else {
		for _, choice := range x.States {
			if prop, ok := o.GetPropertyByChoice(choice); !ok {
				err = errutil.New(o, "does not have choice", choice)
				break
			} else if e := prop.GetValue().SetState(choice); e != nil {
				err = e
				break
			}
		}
	}
	return
}

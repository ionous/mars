package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

type SetNum struct {
	Tgt NumProperty
	Num rt.NumEval
}

type SetTxt struct {
	Tgt TextProperty
	Txt rt.TextEval
}

type SetRef struct {
	Tgt RefProperty
	Ref rt.ObjEval
}

type ClearRef struct {
	Tgt RefProperty
}

type ChangeState struct {
	Ref    rt.ObjEval
	States []ident.Id
}

func Change(tgt rt.ObjEval) ChangeState {
	return ChangeState{tgt, nil}
}

func (x SetNum) Execute(run rt.Runtime) (err error) {
	if n, e := x.Num.GetNumber(run); e != nil {
		err = errutil.New("SetNum.Num", e)
	} else if e := Property(x.Tgt).SetGeneric(run, n); e != nil {
		err = errutil.New("SetNum.Tgt", e)
	}
	return
}

func (x SetTxt) Execute(run rt.Runtime) (err error) {
	if t, e := x.Txt.GetText(run); e != nil {
		err = errutil.New("SetTxt.Txt", e)
	} else if e := Property(x.Tgt).SetGeneric(run, t); e != nil {
		err = errutil.New("SetTxt.Tgt", e)
	}
	return
}

func (x SetRef) Execute(run rt.Runtime) (err error) {
	if obj, e := x.Ref.GetObject(run); e != nil {
		err = errutil.New("SetRef.Ref", e)
	} else {
		ref := rt.Reference(obj.GetId())
		if e := Property(x.Tgt).SetGeneric(run, ref); e != nil {
			err = errutil.New("SetRef.Tgt", e)
		}
	}
	return
}

func (x ClearRef) Execute(run rt.Runtime) (err error) {
	var empty rt.Reference
	if e := Property(x.Tgt).SetGeneric(run, empty); e != nil {
		err = errutil.New("ClearRef.Tgt", e)
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
func (x ChangeState) Execute(run rt.Runtime) (err error) {
	if obj, e := x.Ref.GetObject(run); e != nil {
		err = errutil.New("ChangeState.Ref", e)
	} else {
		for _, choice := range x.States {
			if prop, ok := obj.GetPropertyByChoice(choice); !ok {
				err = errutil.New("ChangeState", obj, "does not have choice", choice)
				break
			} else if e := prop.SetGeneric(choice); e != nil {
				err = errutil.New("ChangeState", e)
				break
			}
		}
	}
	return
}

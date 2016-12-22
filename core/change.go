package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/errutil"
)

type SetNum struct {
	Field  types.NamedProperty
	Object rt.ObjEval
	Num    rt.NumberEval
}

func (x SetNum) Execute(run rt.Runtime) (err error) {
	if n, e := x.Num.GetNumber(run); e != nil {
		err = errutil.New("SetNum.Num", e)
	} else {
		p := Property{x.Field, x.Object}
		if e := p.SetGeneric(run, n); e != nil {
			err = errutil.New("SetNum.Tgt", e)
		}
	}
	return
}

type SetTxt struct {
	Field  types.NamedProperty
	Object rt.ObjEval
	Txt    rt.TextEval
}

func (x SetTxt) Execute(run rt.Runtime) (err error) {
	if t, e := x.Txt.GetText(run); e != nil {
		err = errutil.New("SetTxt.Txt", e)
	} else {
		p := Property{x.Field, x.Object}
		if e := p.SetGeneric(run, t); e != nil {
			err = errutil.New("SetTxt.Tgt", e)
		}
	}
	return
}

type SetObj struct {
	Field  types.NamedProperty
	Object rt.ObjEval
	Ref    rt.ObjEval
}

func (x SetObj) Execute(run rt.Runtime) (err error) {
	if x.Ref == nil {
		err = errutil.New("SetObj Ref is nil")
	} else if obj, e := x.Ref.GetObject(run); e != nil {
		err = errutil.New("SetObj.Ref", e)
	} else {
		p := Property{x.Field, x.Object}
		if e := p.SetGeneric(run, obj); e != nil {
			err = errutil.New("SetObj.Tgt", e)
		}
	}
	return
}

type ChangeState struct {
	Ref    rt.ObjEval
	States types.NamedChoices
}

func Change(tgt rt.ObjEval) ChangeState {
	return ChangeState{tgt, nil}
}

func (p ChangeState) To(state string) ChangeState {
	return p.And(state)
}

func (p ChangeState) And(state string) ChangeState {
	p.States = append(p.States, state)
	return p
}

func (x ChangeState) Execute(run rt.Runtime) (err error) {
	if obj, e := x.Ref.GetObject(run); e != nil {
		err = errutil.New("ChangeState.Ref", e)
	} else {
		for _, choice := range x.States {
			state := MakeStringId(choice) // MARS: *is* there anyway of doing this at ... save? time?
			if prop, ok := obj.GetPropertyByChoice(state); !ok {
				err = errutil.New("ChangeState", obj, "does not have choice", choice)
				break
			} else if e := prop.SetGeneric(rt.State{choice}); e != nil {
				err = errutil.New("ChangeState", e)
				break

			}
		}
	}
	return
}

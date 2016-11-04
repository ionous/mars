package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type CompareType int

const (
	EqualTo CompareType = 1 << iota
	GreaterThan
	LesserThan
	NotEqual = GreaterThan | LesserThan
)

// maybe a regex or glob comparision
// type Match struct {
// }

//
type IsEmpty struct {
	Text rt.TextEval
}

func (empty IsEmpty) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if t, e := empty.Text.GetText(run); e != nil {
		err = errutil.New("IsEmpty.Text", e)
	} else {
		ret = !(len(t.String()) > 0)
	}
	return
}

// IsNot negates a rt.BoolEval (and is itself a rt.BoolEval)
type IsNot struct {
	Negate rt.BoolEval
}

func (neg IsNot) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if tgt, e := neg.Negate.GetBool(run); e != nil {
		err = errutil.New("IsNot.Negate", e)
	} else {
		ret = rt.Bool(!tgt)
	}
	return
}

// IsValid evals true when this refers to a valid object.
type IsValid struct {
	Ref rt.ObjEval
}

func (exists IsValid) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if obj, e := exists.Ref.GetObject(run); e != nil {
		ret = false // if the object doesnt exist, then it's invalid
	} else {
		ret = rt.Bool(!obj.Empty()) // if the object is empty, then it's invalid
	}
	return
}

// IsNumber two numbers (a rt.BoolEval)
type IsNumber struct {
	Src rt.NumEval
	Is  CompareType
	Tgt rt.NumEval
}

func (comp IsNumber) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if src, e := comp.Src.GetNumber(run); e != nil {
		err = errutil.New("IsNumber.Src", e)
	} else if tgt, e := comp.Tgt.GetNumber(run); e != nil {
		err = errutil.New("IsNumber.Tgt", e)
	} else {
		d := src.Float() - tgt.Float()
		switch {
		case d == 0:
			ret = (comp.Is & EqualTo) != 0
		case d < 0:
			ret = (comp.Is & LesserThan) != 0
		case d > 0:
			ret = (comp.Is & GreaterThan) != 0
		}
	}
	return
}

// IsText
type IsText struct {
	Src rt.TextEval
	Is  CompareType
	Tgt rt.TextEval
}

func (comp IsText) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if src, e := comp.Src.GetText(run); e != nil {
		err = errutil.New("IsText.Src", e)
	} else if tgt, e := comp.Tgt.GetText(run); e != nil {
		err = errutil.New("IsText.Tgt", e)
	} else {
		switch comp.Is {
		case EqualTo:
			ret = src == tgt
		case NotEqual:
			ret = src != tgt
		case LesserThan:
			ret = src < tgt
		case GreaterThan:
			ret = src > tgt
		case GreaterThan | EqualTo:
			ret = src >= tgt
		case LesserThan | EqualTo:
			ret = src <= tgt
		default:
			err = errutil.New("IsText.Is", comp.Is, "unknown operand")
		}
	}
	return
}

// IsObject evals true when both Src and Tgt match;
// ( regardless of whether the refs are valid )
type IsObject struct {
	Src, Tgt rt.ObjEval
}

func (op IsObject) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if a, e := op.Src.GetObject(run); e != nil {
		err = errutil.New("IsObject.Src", e)
	} else if tgt, e := op.Tgt.GetObject(run); e != nil {
		err = errutil.New("IsObject.Tgt", e)
	} else {
		ret = rt.Bool(a.GetId().Equals(tgt.GetId()))
	}
	return
}

// IsState (rt.BoolEval) determines if the object is in the named state.
type IsState struct {
	Ref   rt.ObjEval
	State string
}

func (op IsState) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if obj, e := op.Ref.GetObject(run); e != nil {
		err = errutil.New("IsState.Ref", e)
	} else {
		choice := MakeStringId(op.State)
		if prop, ok := obj.GetPropertyByChoice(choice); !ok {
			err = errutil.New("IsState", obj, "choice does not exist", op.State)
		} else if eval, ok := prop.GetGeneric().(rt.StateEval); !ok {
			err = errutil.New("IsState", obj, "property", prop, "unexpected type", sbuf.Type{eval})
		} else if curr, e := eval.GetState(run); e != nil {
			err = errutil.New("IsState", obj, "property", prop, "get state", e)
		} else {
			ret = rt.Bool(curr.Id() == choice)
		}
	}
	return
}

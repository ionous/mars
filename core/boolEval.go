package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type CompareType int

type CompareTo interface {
	Compare() CompareType
}

type EqualTo struct{}
type GreaterThan struct{}
type LesserThan struct{}
type NotEqualTo struct{}

func (EqualTo) Compare() CompareType     { return Compare_EqualTo }
func (GreaterThan) Compare() CompareType { return Compare_GreaterThan }
func (LesserThan) Compare() CompareType  { return Compare_LesserThan }
func (NotEqualTo) Compare() CompareType  { return Compare_NotEqualTo }

const (
	Compare_EqualTo CompareType = 1 << iota
	Compare_GreaterThan
	Compare_LesserThan
	Compare_NotEqualTo = Compare_GreaterThan | Compare_LesserThan
)

// maybe a regex or glob comparision
// type Match struct {
// }

//
type IsEmpty struct {
	Text rt.TextEval
}

func (empty IsEmpty) GetBool(run rt.Runtime) (ret bool, err error) {
	if t, e := empty.Text.GetText(run); e != nil {
		err = errutil.New("IsEmpty.Text", e)
	} else {
		ret = !(len(t) > 0)
	}
	return
}

// IsNot negates a rt.BoolEval (and is itself a rt.BoolEval)
type IsNot struct {
	Negate rt.BoolEval
}

func (neg IsNot) GetBool(run rt.Runtime) (ret bool, err error) {
	if tgt, e := neg.Negate.GetBool(run); e != nil {
		err = errutil.New("IsNot.Negate", e)
	} else {
		ret = !tgt
	}
	return
}

// IsValid evals true when this refers to a valid object.
type IsValid struct {
	Ref rt.ObjEval
}

func (exists IsValid) GetBool(run rt.Runtime) (ret bool, err error) {
	if obj, e := exists.Ref.GetObject(run); e != nil {
		ret = false // if the object doesnt exist, then it's invalid
	} else {
		ret = obj.Exists() // if the object is empty, then it's invalid
	}
	return
}

type IsFromClass struct {
	Ref   rt.ObjEval
	Class string
}

func (op IsFromClass) GetBool(run rt.Runtime) (ret bool, err error) {
	if obj, e := op.Ref.GetObject(run); e != nil {
		err = e
	} else {
		choice := MakeStringId(op.Class)
		ret = run.AreCompatible(obj.GetParentClass(), choice)
	}
	return
}

// IsNum two numbers (a rt.BoolEval)
type IsNum struct {
	Src rt.NumberEval
	Is  CompareTo
	Tgt rt.NumberEval
}

func (comp IsNum) GetBool(run rt.Runtime) (ret bool, err error) {
	if src, e := comp.Src.GetNumber(run); e != nil {
		err = errutil.New("IsNum.Src", e)
	} else if tgt, e := comp.Tgt.GetNumber(run); e != nil {
		err = errutil.New("IsNum.Tgt", e)
	} else {
		d := src - tgt
		switch cmp := comp.Is.Compare(); {
		case d == 0:
			ret = (cmp & Compare_EqualTo) != 0
		case d < 0:
			ret = (cmp & Compare_LesserThan) != 0
		case d > 0:
			ret = (cmp & Compare_GreaterThan) != 0
		}
	}
	return
}

// IsText
type IsText struct {
	Src rt.TextEval
	Is  CompareTo
	Tgt rt.TextEval
}

func (comp IsText) GetBool(run rt.Runtime) (ret bool, err error) {
	if src, e := comp.Src.GetText(run); e != nil {
		err = errutil.New("IsText.Src", e)
	} else if tgt, e := comp.Tgt.GetText(run); e != nil {
		err = errutil.New("IsText.Tgt", e)
	} else {
		switch cmp := comp.Is.Compare(); cmp {
		case Compare_EqualTo:
			ret = src == tgt
		case Compare_NotEqualTo:
			ret = src != tgt
		case Compare_LesserThan:
			ret = src < tgt
		case Compare_GreaterThan:
			ret = src > tgt
		case Compare_GreaterThan | Compare_EqualTo:
			ret = src >= tgt
		case Compare_LesserThan | Compare_EqualTo:
			ret = src <= tgt
		default:
			err = errutil.New("IsText.Is", cmp, "unknown operand")
		}
	}
	return
}

// IsObj evals true when both Src and Tgt match;
// ( regardless of whether the refs are valid )
type IsObj struct {
	Src, Tgt rt.ObjEval
}

func (op IsObj) GetBool(run rt.Runtime) (ret bool, err error) {
	if a, e := op.Src.GetObject(run); e != nil {
		err = errutil.New("IsObj.Src", e)
	} else if tgt, e := op.Tgt.GetObject(run); e != nil {
		err = errutil.New("IsObj.Tgt", e)
	} else {
		ret = a.GetId().Equals(tgt.GetId())
	}
	return
}

// IsState (rt.BoolEval) determines if the object is in the named state.
type IsState struct {
	Ref   rt.ObjEval
	State string
}

func (op IsState) GetBool(run rt.Runtime) (ret bool, err error) {
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
			ret = curr == choice
		}
	}
	return
}

package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/sbuf"
)

type CompareType int

const (
	EqualTo CompareType = 1 << iota
	GreaterThan
	LesserThan
)

// Compare two numbers (a rt.BoolEval)
type Compare struct {
	Src rt.NumEval
	Is  CompareType
	Tgt rt.NumEval
}

// Is the object in the named state (a rt.BoolEval)
type Is struct {
	Ref   rt.ObjEval
	State string
}

// Not negates a rt.BoolEval (and is itself a rt.BoolEval)
type Not struct {
	Negate rt.BoolEval
}

//
type IsEmpty struct {
	Text rt.TextEval
}

// Equals evals true when both Src and Tgt match;
// ( regardless of whether the refs are valid )
type Equals struct {
	Src, Tgt rt.ObjEval
}

// Exists evals true when this refers to a valid object.
type Exists struct {
	Ref rt.ObjEval
}

func (empty IsEmpty) GetBool(run rt.Runtime) (ret bool, err error) {
	if t, e := empty.Text.GetText(run); e != nil {
		err = errutil.New("IsEmpty.Text", e)
	} else {
		ret = !(len(t.String()) > 0)
	}
	return
}

func (neg Not) GetBool(run rt.Runtime) (ret bool, err error) {
	if b, e := neg.Negate.GetBool(run); e != nil {
		err = e
	} else {
		ret = !b
	}
	return
}

// FIX: what to do with exists?
func (exists Exists) GetBool(run rt.Runtime) (bool, error) {
	_, e := exists.Ref.GetObject(run)
	return e == nil, nil
}

func (comp Compare) GetBool(run rt.Runtime) (ret bool, err error) {
	if a, e := comp.Src.GetNumber(run); e != nil {
		err = errutil.New("Compare.Src", e)
	} else if b, e := comp.Tgt.GetNumber(run); e != nil {
		err = errutil.New("Compare.Tgt", e)
	} else {
		d := a.Float() - b.Float()
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

func (req Equals) GetBool(run rt.Runtime) (ret bool, err error) {
	if a, e := req.Src.GetObject(run); e != nil {
		err = errutil.New("Equals.Src", e)
	} else if b, e := req.Tgt.GetObject(run); e != nil {
		err = errutil.New("Equals.Tgt", e)
	} else {
		ret = a.GetId().Equals(b.GetId())
	}
	return
}

//func (oa *GameObject) Is(state string)
func (op Is) GetBool(run rt.Runtime) (ret bool, err error) {
	if obj, e := op.Ref.GetObject(run); e != nil {
		err = e
	} else {
		choice := MakeStringId(op.State)
		if prop, ok := obj.GetPropertyByChoice(choice); !ok {
			err = errutil.New("Is", obj, "choice does not exist", choice)
		} else if currChoice, ok := prop.GetGeneric().(ident.Id); !ok {
			err = errutil.New("Is op", obj, "property", prop, "unexpected type", sbuf.Type{currChoice})
		} else {
			ret = currChoice == choice
		}
	}
	return
}

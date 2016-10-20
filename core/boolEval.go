package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

type CompareType int

const (
	EqualTo CompareType = 1 << iota
	GreaterThan
	LesserThan
)

// Compare two numbers (a rt.BoolEval)
type Compare struct {
	A  rt.NumEval
	Is CompareType
	B  rt.NumEval
}

// Is the object in the named state (a rt.BoolEval)
type Is struct {
	Ref   rt.RefEval
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

// Equals evals true when both A and B match;
// ( regardless of whether the refs are valid )
type Equals struct {
	A, B rt.RefEval
}

// Exists evals true when this refers to a valid object.
type Exists struct {
	Ref rt.RefEval
}

func (empty IsEmpty) GetBool(r rt.Runtime) (ret bool, err error) {
	if t, e := empty.Text.GetText(r); e != nil {
		err = e
	} else {
		ret = !(len(t.String()) > 0)
	}
	return
}

func (neg Not) GetBool(r rt.Runtime) (ret bool, err error) {
	if b, e := neg.Negate.GetBool(r); e != nil {
		err = e
	} else {
		ret = !b
	}
	return
}

func (exists Exists) GetBool(r rt.Runtime) (ret bool, err error) {
	if ref, e := exists.Ref.GetReference(r); e != nil {
		err = e
	} else if _, e := r.GetObject(ref); e == nil {
		ret = true
	}
	return
}

func (comp Compare) GetBool(r rt.Runtime) (ret bool, err error) {
	if a, e := comp.A.GetNumber(r); e != nil {
		err = e
	} else if b, e := comp.B.GetNumber(r); e != nil {
		err = e
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

func (req Equals) GetBool(r rt.Runtime) (ret bool, err error) {
	if a, e := req.A.GetReference(r); e != nil {
		err = e
	} else if b, e := req.B.GetReference(r); e != nil {
		err = e
	} else {
		ret = a.Id().Equals(b.Id())
	}
	return
}

//func (oa *GameObject) Is(state string)
func (oi Is) GetBool(r rt.Runtime) (ret bool, err error) {
	if ref, e := oi.Ref.GetReference(r); e != nil {
		err = e
	} else if o, e := r.GetObject(ref); e != nil {
		err = e
	} else {
		choice := MakeStringId(oi.State)
		if prop, ok := o.GetPropertyByChoice(choice); !ok {
			err = errutil.New("object choice does not exist", o, choice)
		} else if currChoice, ok := prop.GetGeneric().(ident.Id); !ok {
			err = errutil.New("object property of unexpected type", o, choice)
		} else {
			ret = currChoice == choice
		}
	}
	return
}

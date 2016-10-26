package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type Context struct {
	Ref rt.ObjEval
	Run rt.Execute
}

func (c Context) Execute(run rt.Runtime) (err error) {
	if obj, e := c.Ref.GetObject(run); e != nil {
		err = e
	} else {
		newScope := scope.Make(run, scope.ObjectScope{obj})
		if e := c.Run.Execute(newScope); e != nil {
			err = e
		}
	}
	return
}

// GetNum returns a numer from the current context.
type GetNum struct {
	Name string
}

func (c GetNum) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if eval, e := run.FindValue(c.Name); e != nil {
		err = e
	} else if neval, ok := eval.(rt.NumEval); !ok {
		err = errutil.New("value", c.Name, "is not a number eval", sbuf.Type{eval})
	} else if v, e := neval.GetNumber(run); e != nil {
		err = e
	} else {
		ret = v
	}
	return
}

// GetText returns a text value from the current context.
type GetText struct {
	Name string
}

func (c GetText) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if eval, e := run.FindValue(c.Name); e != nil {
		err = e
	} else if teval, ok := eval.(rt.TextEval); !ok {
		err = errutil.New("value", c.Name, "is not text eval", sbuf.Type{eval})
	} else if v, e := teval.GetText(run); e != nil {
		err = e
	} else {
		ret = v
	}
	return
}

// GetText returns a text value from the current context.
type GetObject struct {
	Name string
}

func (c GetObject) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if eval, e := run.FindValue(c.Name); e != nil {
		err = e
	} else if objeval, ok := eval.(rt.ObjEval); !ok {
		err = errutil.New("value", c.Name, "is not text eval", sbuf.Type{eval})
	} else if v, e := objeval.GetObject(run); e != nil {
		err = e
	} else {
		ret = v
	}
	return
}

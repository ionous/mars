package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"reflect"
)

type Context struct {
	Ref rt.ObjEval
	Run rt.Execute
}

func (c Context) Execute(run rt.Runtime) (err error) {
	if obj, e := c.Ref.GetObject(run); e != nil {
		err = e
	} else {
		run.PushScope(ObjectScope{obj}, nil)
		defer run.PopScope()
		//
		if e := c.Run.Execute(run); e != nil {
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
	scope, _ := run.GetScope()
	if eval, e := scope.FindValue(c.Name); e != nil {
		err = e
	} else if neval, ok := eval.(rt.NumEval); !ok {
		err = errutil.New("value", c.Name, "is not a number eval", reflect.TypeOf(eval))
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
	scope, _ := run.GetScope()
	if eval, e := scope.FindValue(c.Name); e != nil {
		err = e
	} else if teval, ok := eval.(rt.TextEval); !ok {
		err = errutil.New("value", c.Name, "is not text eval", reflect.TypeOf(eval))
	} else if v, e := teval.GetText(run); e != nil {
		err = e
	} else {
		ret = v
	}
	return
}

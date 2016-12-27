package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type Using struct {
	Object    rt.ObjEval
	Run, Else rt.Statements
}

func (c Using) Execute(run rt.Runtime) (err error) {
	if obj, e := c.Object.GetObject(run); e != nil {
		err = e
	} else if !obj.Exists() {
		if e := c.Else.ExecuteList(run); e != nil {
			err = e
		}
	} else {
		run := scope.Make(run, scope.NewObjectScope(obj), run)
		if e := c.Run.ExecuteList(run); e != nil {
			err = e
		}
	}
	return
}

type GetBool struct {
	Field types.NamedProperty
}

func (c GetBool) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if eval, e := run.FindValue(c.Field.String()); e != nil {
		err = e
	} else if neval, ok := eval.(rt.BoolEval); !ok {
		err = errutil.New("value", c.Field, "is a", sbuf.Type{eval}, "not BoolEval")
	} else if v, e := neval.GetBool(run); e != nil {
		err = e
	} else {
		ret = v
	}
	return
}

// GetNum returns a number from the current context.
type GetNum struct {
	Field types.NamedProperty
}

func (c GetNum) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if eval, e := run.FindValue(c.Field.String()); e != nil {
		err = e
	} else if neval, ok := eval.(rt.NumberEval); !ok {
		err = errutil.New("value", c.Field, "is a", sbuf.Type{eval}, "not NumberEval")
	} else if v, e := neval.GetNumber(run); e != nil {
		err = e
	} else {
		ret = v
	}
	return
}

// GetText returns a text value from the current context.
type GetText struct {
	Field types.NamedProperty
}

func (c GetText) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if eval, e := run.FindValue(c.Field.String()); e != nil {
		err = e
	} else if teval, ok := eval.(rt.TextEval); !ok {
		err = errutil.New("value", c.Field, "is a", sbuf.Type{eval}, "not TextEval")
	} else if v, e := teval.GetText(run); e != nil {
		err = e
	} else {
		ret = v
	}
	return
}

// GetObject returns a object value from the current conobject.
type GetObj struct {
	Field types.NamedProperty
}

func (c GetObj) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if eval, e := run.FindValue(c.Field.String()); e != nil {
		err = e
	} else if objeval, ok := eval.(rt.ObjEval); !ok {
		err = errutil.New("value", c.Field, "is not object eval", sbuf.Type{eval})
	} else if v, e := objeval.GetObject(run); e != nil {
		err = e
	} else {
		ret = v
	}
	return
}

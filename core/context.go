package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
	"reflect"
)

type Context struct {
	Ref rt.RefEval
	Run rt.Execute
}

func (c Context) Execute(r rt.Runtime) (err error) {
	if ref, e := c.Ref.GetReference(r); e != nil {
		err = e
	} else if pushed, e := c.push(r, ref); e != nil {
		err = e
	} else {
		if pushed {
			defer r.PopScope()
		}
		if e := c.Run.Execute(r); e != nil {
			err = e
		}
	}
	return
}
func (c Context) push(r rt.Runtime, ref rt.Reference) (pushed bool, err error) {
	if !ref.Id().Empty() {
		if o, e := r.GetObject(ref); e != nil {
			err = e
		} else {
			r.PushScope(ObjectScope{o}, nil)
			pushed = true
		}
	}
	return
}

type GetNum struct {
	Name string
}

func (c GetNum) GetNumber(r rt.Runtime) (ret rt.Number, err error) {
	scope, _ := r.GetScope()
	if eval, e := scope.FindValue(c.Name); e != nil {
		err = e
	} else if neval, ok := eval.(rt.NumEval); ok {
		if v, e := neval.GetNumber(r); e != nil {
			ret = v
		} else {
			err = e
		}
	} else {
		err = errutil.New("value", c.Name, "is not a number eval", reflect.TypeOf(eval))
	}
	return
}

// GetText returns a text value from the current context.
type GetText struct {
	Name string
}

func (c GetText) GetText(r rt.Runtime) (ret rt.Text, err error) {
	scope, _ := r.GetScope()
	if eval, e := scope.FindValue(c.Name); e != nil {
		err = e
	} else if teval, ok := eval.(rt.TextEval); ok {
		if v, e := teval.GetText(r); e != nil {
			ret = v
		} else {
			err = e
		}
	} else {
		err = errutil.New("value", c.Name, "is not text eval", reflect.TypeOf(eval))
	}
	return
}

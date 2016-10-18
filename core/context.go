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
		if o, e := MakeObject(r, ref); e != nil {
			err = e
		} else {
			r.PushScope(o, nil)
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
	if v, e := scope.FindValue(c.Name); e != nil {
		err = e
	} else if n, ok := v.(rt.Number); ok {
		ret = n
	} else {
		err = errutil.New("value", c.Name, "is not a number", reflect.TypeOf(v))
	}
	return
}

// GetText returns a text value from the current context.
type GetText struct {
	Name string
}

func (c GetText) GetText(r rt.Runtime) (ret rt.Text, err error) {
	scope, _ := r.GetScope()
	if v, e := scope.FindValue(c.Name); e != nil {
		err = e
	} else if n, ok := v.(rt.Text); ok {
		ret = n
	} else {
		err = errutil.New("value", c.Name, "is not text", reflect.TypeOf(v))
	}
	return
}

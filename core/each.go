package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type EachNum struct {
	For      rt.NumListEval
	Go, Else rt.Execute
}

type EachText struct {
	For      rt.TextListEval
	Go, Else rt.Execute
}

type EachObj struct {
	For      rt.ObjListEval
	Go, Else rt.Execute
}

type IfEach struct {
	IsFirst, IsLast bool
}

func (t IfEach) GetBool(run rt.Runtime) (rt.Bool, error) {
	var err error
	b := (t.IsFirst && ifEach(run, "@first", &err)) ||
		(err == nil && t.IsLast && ifEach(run, "@last", &err))
	return rt.Bool(b), err
}

func ifEach(run rt.Runtime, name string, err *error) (ret bool) {
	if v, e := run.FindValue(name); e != nil {
		*err = errutil.New("ifEach, not in a loop", e)
	} else if eval, ok := v.(rt.BoolEval); !ok {
		*err = errutil.New("ifEach, expected bool", sbuf.Type{v})
	} else if b, e := eval.GetBool(run); e != nil {
		*err = e
	} else {
		ret = bool(b)
	}
	return
}

type EachIndex struct{}

func (t EachIndex) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if v, e := run.FindValue("@index"); e != nil {
		err = errutil.New("EachIndex", "not in a loop", e)
	} else if eval, ok := v.(rt.NumEval); !ok {
		err = errutil.New("ifEach, expected num", sbuf.Type{v})
	} else {
		ret, err = eval.GetNumber(run)
	}
	return
}

func (f EachNum) Execute(run rt.Runtime) error {
	return eachValue(run, f.For, f.Go, f.Else, func(i int) (ret meta.Generic, err error) {
		if v, e := f.For.GetNumberIdx(run, i); e != nil {
			err = e
		} else {
			ret = rt.NumEval(v)
		}
		return
	})
}

func (f EachText) Execute(run rt.Runtime) error {
	return eachValue(run, f.For, f.Go, f.Else, func(i int) (ret meta.Generic, err error) {
		if v, e := f.For.GetTextIdx(run, i); e != nil {
			err = e
		} else {
			ret = rt.TextEval(v)
		}
		return
	})
}

func (f EachObj) Execute(run rt.Runtime) error {
	return eachValue(run, f.For, f.Go, f.Else, func(i int) (ret meta.Generic, err error) {
		if v, e := f.For.GetReferenceIdx(run, i); e != nil {
			err = e
		} else {
			ret = rt.ObjEval(v)
		}
		return
	})
}

type makeValue func(i int) (meta.Generic, error)

func eachValue(run rt.Runtime, list rt.ListEval, loop, otherwise rt.Execute, value makeValue) (err error) {
	if c := list.GetCount(); c == 0 {
		otherwise.Execute(run)
	} else {
		it := scope.NewLoopMaker(run)
		for i := 0; i < c; i++ {
			if v, e := value(i); e != nil {
				err = e
				break
			} else {
				e := loop.Execute(it.Looper(i+1, i == 0, i+1 == c, v))
				if e != nil {
					err = e
					break
				}
			}
		}
	}
	return
}

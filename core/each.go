package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
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

func (t IfEach) GetBool(run rt.Runtime) (ret bool, err error) {
	if _, p := run.GetScope(); p == nil {
		err = errutil.New("IfEach", "not in a loop")
	} else {
		ret = (p.IsFirst && t.IsFirst) || (p.IsLast && t.IsLast)
	}
	return
}

type EachIndex struct{}

func (t EachIndex) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if _, p := run.GetScope(); p == nil {
		err = errutil.New("EachIndex", "not in a loop")
	} else {
		ret = rt.Number(float64(p.Index + 1))
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

type ValueScope struct {
	val meta.Generic
}

func (vs ValueScope) FindValue(name string) (ret meta.Generic, err error) {
	if name != "" {
		// FIX: what if it is?
		err = errutil.New("context is not an object")
	} else {
		ret = vs.val
	}
	return
}

func eachValue(run rt.Runtime, list rt.ListEval, loop, otherwise rt.Execute, value makeValue) (err error) {
	if c := list.GetCount(); c == 0 {
		otherwise.Execute(run)
	} else {
		for i := 0; i < c; i++ {
			if v, e := value(i); e != nil {
				err = e
				break
			} else {
				run.PushScope(ValueScope{v}, &rt.IndexInfo{
					Index:   i,
					IsFirst: i == 0,
					IsLast:  i+1 == c,
				})
				e := loop.Execute(run)
				run.PopScope()
				if e != nil {
					err = e
					break
				}
			}
		}
	}
	return
}

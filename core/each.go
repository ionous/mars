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
	For      rt.RefListEval
	Go, Else rt.Execute
}

type IfEach struct {
	IsFirst, IsLast bool
}

func (t IfEach) GetBool(r rt.Runtime) (ret bool, err error) {
	if _, p := r.GetScope(); p == nil {
		err = errutil.New("IfEach", "not in a loop")
	} else {
		ret = (p.IsFirst && t.IsFirst) || (p.IsLast && t.IsLast)
	}
	return
}

type EachIndex struct{}

func (t EachIndex) GetNumber(r rt.Runtime) (ret rt.Number, err error) {
	if _, p := r.GetScope(); p == nil {
		err = errutil.New("EachIndex", "not in a loop")
	} else {
		ret = rt.Number(float64(p.Index + 1))
	}
	return
}

func (f EachNum) Execute(r rt.Runtime) error {
	return eachValue(r, f.For, f.Go, f.Else, func(i int) (ret meta.Generic, err error) {
		if v, e := f.For.GetNumberIdx(r, i); e != nil {
			err = e
		} else {
			ret = rt.NumEval(v)
		}
		return
	})
}

func (f EachText) Execute(r rt.Runtime) error {
	return eachValue(r, f.For, f.Go, f.Else, func(i int) (ret meta.Generic, err error) {
		if v, e := f.For.GetTextIdx(r, i); e != nil {
			err = e
		} else {
			ret = rt.TextEval(v)
		}
		return
	})
}

func (f EachObj) Execute(r rt.Runtime) error {
	return eachValue(r, f.For, f.Go, f.Else, func(i int) (ret meta.Generic, err error) {
		if v, e := f.For.GetReferenceIdx(r, i); e != nil {
			err = e
		} else {
			ret = rt.RefEval(v)
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
		err = errutil.New("context is not an object")
	} else {
		ret = vs.val
	}
	return
}

func eachValue(r rt.Runtime, list rt.ListEval, loop, otherwise rt.Execute, value makeValue) (err error) {
	if c := list.GetCount(); c == 0 {
		otherwise.Execute(r)
	} else {
		for i := 0; i < c; i++ {
			if v, e := value(i); e != nil {
				err = e
				break
			} else {
				r.PushScope(ValueScope{v}, &rt.IndexInfo{
					Index:   i,
					IsFirst: i == 0,
					IsLast:  i+1 == c,
				})
				e := loop.Execute(r)
				r.PopScope()
				if e != nil {
					err = e
					break
				}
			}
		}
	}
	return
}

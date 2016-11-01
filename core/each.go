package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type EachNum struct {
	In       rt.NumListEval
	Go, Else rt.Execute
}

type EachText struct {
	In       rt.TextListEval
	Go, Else rt.Execute
}

type EachObj struct {
	In       rt.ObjListEval
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
		*err = errutil.New("ifEach, not in a l", e)
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
		err = errutil.New("EachIndex", "not in a l", e)
	} else if eval, ok := v.(rt.NumEval); !ok {
		err = errutil.New("ifEach, expected num", sbuf.Type{v})
	} else {
		ret, err = eval.GetNumber(run)
	}
	return
}

func (f EachNum) Execute(run rt.Runtime) (err error) {
	if it, e := f.In.GetNumStream(run); e != nil {
		err = e
	} else if !it.HasNext() {
		err = f.Else.Execute(run)
	} else {
		for l := scope.NewLooper(it); l.HasNext(); {
			if v, e := it.GetNext(); e != nil {
				err = e
				break
			} else {
				s := scope.MakeChain(run, l.NextScope(v), scope.NewNumScope(v))
				if e := f.Go.Execute(s); e != nil {
					err = e
					break
				}
			}
		}
	}
	return
}

func (f EachText) Execute(run rt.Runtime) (err error) {
	if it, e := f.In.GetTextStream(run); e != nil {
		err = e
	} else if !it.HasNext() {
		err = f.Else.Execute(run)
	} else {
		for l := scope.NewLooper(it); l.HasNext(); {
			if v, e := it.GetNext(); e != nil {
				err = e
				break
			} else {
				s := scope.MakeChain(run, l.NextScope(v), scope.NewTextScope(v))
				if e := f.Go.Execute(s); e != nil {
					err = e
					break
				}
			}
		}
	}
	return
}

func (f EachObj) Execute(run rt.Runtime) (err error) {
	if it, e := f.In.GetObjStream(run); e != nil {
		err = e
	} else if !it.HasNext() {
		err = f.Else.Execute(run)
	} else {
		for l := scope.NewLooper(it); l.HasNext(); {
			if obj, e := it.GetNext(); e != nil {
				err = e
				break
			} else {
				s := scope.MakeChain(run, l.NextScope(obj), scope.NewObjectScope(obj))
				if e := f.Go.Execute(s); e != nil {
					err = e
					break
				}
			}
		}
	}
	return
}

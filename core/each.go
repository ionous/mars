package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type ForEachNum struct {
	In       rt.NumListEval
	Go, Else rt.Execute
}

type ForEachText struct {
	In       rt.TextListEval
	Go, Else rt.Execute
}

type ForEachObj struct {
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
	} else if eval, ok := v.(rt.NumberEval); !ok {
		err = errutil.New("ifEach, expected num", sbuf.Type{v})
	} else {
		ret, err = eval.GetNumber(run)
	}
	return
}

func (f ForEachNum) Execute(run rt.Runtime) (err error) {
	if it, e := f.In.GetNumberStream(run); e != nil {
		err = e
	} else if !it.HasNext() {
		if f.Else != nil {
			if e := f.Else.Execute(run); e != nil {
				err = errutil.New("failed each num else", e)
			}
		}
	} else {
		for l := scope.NewLooper(it); l.HasNext(); {
			if v, e := it.GetNext(); e != nil {
				err = errutil.New("failed each num get", e)
				break
			} else {
				run := scope.Make(run, l.NextScope(v), run)
				if e := f.Go.Execute(run); e != nil {
					err = errutil.New("failed each num go", v, e)
					break
				}
			}
		}
	}
	return
}

func (f ForEachText) Execute(run rt.Runtime) (err error) {
	if it, e := f.In.GetTextStream(run); e != nil {
		err = e
	} else if !it.HasNext() {
		if f.Else != nil {
			if e := f.Else.Execute(run); e != nil {
				err = errutil.New("failed each text else", e)
			}
		}
	} else {
		for l := scope.NewLooper(it); l.HasNext(); {
			if v, e := it.GetNext(); e != nil {
				err = errutil.New("failed each text get", e)
				break
			} else {
				run := scope.Make(run, l.NextScope(v), run)
				if e := f.Go.Execute(run); e != nil {
					err = errutil.New("failed each text go", v, e)
					break
				}
			}
		}
	}
	return
}

func (f ForEachObj) Execute(run rt.Runtime) (err error) {
	if it, e := f.In.GetObjStream(run); e != nil {
		err = e
	} else if !it.HasNext() {
		if f.Else != nil {
			if e := f.Else.Execute(run); e != nil {
				err = errutil.New("failed each obj else", e)
			}
		}
	} else {
		for i, l := 0, scope.NewLooper(it); l.HasNext(); i++ {
			if obj, e := it.GetNext(); e != nil {
				err = errutil.New("failed each obj get", e)
				break
			} else {
				sub := scope.Make(run, l.NextScope(obj), scope.NewObjectScope(obj), run)
				if e := f.Go.Execute(sub); e != nil {
					err = errutil.New("failed each obj go", obj, e)
					break
				}
			}
		}
	}
	return
}

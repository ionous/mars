package stream

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
)

// First finds the first matching object.
// if none match it returns else.
// if else is nil, and none matched, it errors.
// https://api.dartlang.org/stable/1.20.1/dart-async/Stream/firstWhere.html
type First struct {
	In       rt.ObjListEval
	Matching rt.BoolEval
	Else     rt.ObjEval
}

func (f First) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if it, e := f.In.GetObjStream(run); e != nil {
		err = errutil.New("stream.First failed to get stream because", e)
	} else {
		completed := false
		for l := scope.NewLooper(it); l.HasNext(); {
			if obj, e := it.GetNext(); e != nil {
				err = errutil.New("stream.First failed get next because", e)
				break
			} else {
				run := scope.Make(run, l.NextScope(obj), scope.NewObjectScope(obj), run)
				if b, e := f.match(run); e != nil {
					err = errutil.New("stream.First matching error", e)
				} else if b {
					ret, completed = obj, true
					break
				}
			}
		}
		if err == nil && !completed {
			if f.Else == nil {
				err = rt.StreamEnd("stream.First match not found")
			} else {
				ret, err = f.Else.GetObject(run)
			}
		}
	}
	return
}

func (f *First) match(run rt.Runtime) (ret bool, err error) {
	if f.Matching == nil {
		ret = true
	} else {
		if b, e := f.Matching.GetBool(run); e != nil {
			err = e
		} else if b {
			ret = true
		}
	}
	return
}

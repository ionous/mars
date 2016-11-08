package stream

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
)

type StreamEnd struct {
	reason string
}

func (e StreamEnd) Error() string {
	return e.reason
}

type Generate struct {
	First, Next rt.ObjEval
}

// ObjListEval
func (gen Generate) GetObjStream(run rt.Runtime) (ret rt.ObjectStream, err error) {
	if p, e := gen.First.GetObject(run); e != nil {
		if _, end := e.(StreamEnd); end {
			ret = &genx{}
		} else {
			err = errutil.New("generator", e)
		}
	} else {
		ret = &genx{run, p, gen.Next, true}
	}
	return
}

type genx struct {
	run     rt.Runtime
	next    rt.Object
	src     rt.ObjEval
	hasNext bool
}

func (gen *genx) HasNext() bool {
	return gen.hasNext
}

func (gen *genx) GetNext() (next rt.Object, err error) {
	if !gen.hasNext {
		err = errutil.New("stream closed")
	} else {
		next = gen.next
		run := scope.Make(gen.run, scope.NewObjectScope(next), gen.run)
		if n, e := gen.src.GetObject(run); e != nil {
			if _, end := e.(StreamEnd); end {
				gen.hasNext = false
			} else {
				err = errutil.New("generator get next", e)
			}
		} else {
			gen.next = n
		}
	}
	return
}

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
		err = errutil.New("stream first matching failed to get stream", e)
	} else {
		completed := false
		for l := scope.NewLooper(it); l.HasNext(); {
			if obj, e := it.GetNext(); e != nil {
				err = errutil.New("stream first, get next", e)
				break
			} else {
				run := scope.Make(run, l.NextScope(obj), scope.NewObjectScope(obj), run)
				if b, e := f.Matching.GetBool(run); e != nil {
					err = errutil.New("stream first, matching", e)
				} else if b {
					ret, completed = obj, true
					break
				}
			}
		}
		if err == nil && !completed {
			if f.Else == nil {
				err = StreamEnd{"stream first match not found"}
			} else {
				ret, err = f.Else.GetObject(run)
			}
		}
	}
	return
}

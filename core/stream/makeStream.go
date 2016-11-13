package stream

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
)

// Generate creates a stream of objects.
type MakeStream struct {
	First, Next rt.ObjEval
}

// ObjListEval
func (g MakeStream) GetObjStream(run rt.Runtime) (ret rt.ObjectStream, err error) {
	if p, e := g.First.GetObject(run); e != nil {
		if _, streamEnd := e.(rt.StreamEnd); streamEnd {
			ret = EmptyStream{}
		} else {
			err = errutil.New("MakeStream", e)
		}
	} else {
		ret = &genx{run, p, g.Next, true}
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
		err = rt.StreamExceeded("MakeStream")
	} else {
		next = gen.next
		run := scope.Make(gen.run, scope.NewObjectScope(next), gen.run)
		if n, e := gen.src.GetObject(run); e != nil {
			if _, end := e.(rt.StreamEnd); end {
				gen.hasNext = false
			} else {
				err = errutil.New("MakeStream GetNext", e)
			}
		} else {
			gen.next = n
		}
	}
	return
}

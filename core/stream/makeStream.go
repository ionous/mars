package stream

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
)

// Generate creates a stream of objects.
type MakeStream struct {
	Using, Next rt.ObjEval
}

// ObjListEval
func (g MakeStream) GetObjStream(run rt.Runtime) (ret rt.ObjectStream, err error) {
	if prime, e := g.Using.GetObject(run); e != nil {
		if _, streamEnd := e.(rt.StreamEnd); streamEnd {
			ret = EmptyStream{}
		} else {
			err = errutil.New("MakeStream error:", e)
		}
	} else {
		g := &genx{run: run, gen: g.Next}
		g.next, g.hasNext, err = g.advance(prime)
		ret = g
	}
	return
}

type genx struct {
	run     rt.Runtime
	gen     rt.ObjEval
	next    rt.Object
	hasNext bool
}

func (g *genx) HasNext() bool {
	return g.hasNext
}

func (g *genx) GetNext() (ret rt.Object, err error) {
	if !g.hasNext {
		err = rt.StreamExceeded("MakeStream")
	} else {
		if n, ok, e := g.advance(g.next); e != nil {
			err = e
		} else {
			ret, g.next, g.hasNext = g.next, n, ok
		}
	}
	return
}

// Advance does not modify g.
func (g *genx) advance(obj rt.Object) (ret rt.Object, okay bool, err error) {
	run := scope.Make(g.run, scope.NewObjectScope(obj), g.run)
	if next, e := g.gen.GetObject(run); e != nil {
		if _, end := e.(rt.StreamEnd); !end {
			err = errutil.New("MakeStream failed to advance because", e)
		}
	} else {
		ret, okay = next, true
	}
	return
}

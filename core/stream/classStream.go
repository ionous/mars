package stream

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/sbuf"
)

type ClassStream struct {
	Name  string
	Exact bool
}

// ObjListEval
func (g ClassStream) GetObjStream(run rt.Runtime) (ret rt.ObjectStream, err error) {
	if cls, ok := run.GetClass(ident.MakeId(g.Name)); !ok {
		err = errutil.New("ClassStream: no such class", sbuf.Q(g.Name))
	} else {
		q := &clsx{run: run, cls: cls.GetId(), exact: g.Exact}
		q.idx, q.next = q.advance()
		ret = q
	}
	return
}

// clsx creates a stream of objects.
type clsx struct {
	run   rt.Runtime
	cls   ident.Id
	exact bool
	idx   int
	next  meta.Instance
}

func (q *clsx) HasNext() bool {
	return q.next != nil
}

func (q *clsx) GetNext() (ret rt.Object, err error) {
	if n := q.next; n != nil {
		ret = rt.Object{q.next}
		q.idx, q.next = q.advance()
	} else {
		err = rt.StreamExceeded("clsx")
	}
	return
}

// Advance does not modify q.
func (q *clsx) advance() (int, meta.Instance) {
	m, idx, clsid := q.run, q.idx, q.cls
	l := m.NumInstance()
	if q.exact {
		for ; idx < l; idx++ {
			inst := m.InstanceNum(idx)
			if id := inst.GetParentClass(); id == clsid {
				return idx + 1, inst // explicit return to handle idx renaming
			}
		}
	} else {
		for ; idx < l; idx++ {
			inst := m.InstanceNum(idx)
			if id := inst.GetParentClass(); m.AreCompatible(id, clsid) {
				return idx + 1, inst // explicit return to handle idx renaming
			}
		}
	}
	return l, nil
}

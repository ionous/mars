package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

// Named searches for objects by name, as opposed to core.Id which uses direct lookup.
type Named struct {
	Name string
}

// GetObject searches through the scope for an object matching Name
func (op Named) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if v, e := run.FindValue(op.Name); e != nil {
		err = errutil.New("Named.GetObject", e)
	} else if x, ok := v.(rt.ObjEval); !ok {
		err = errutil.New("Named.GetObject", op.Name, "is not an object")
	} else if r, e := x.GetObject(run); e != nil {
		err = errutil.New("Named.GetObject", e)
	} else {
		ret = r
	}
	return
}

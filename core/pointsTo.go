package core

import (
	"github.com/ionous/mars/rt"
)

func NullRef() (ret rt.Reference) {
	return ret
}

// PointsTo returns a reference if valid, error if the object doesnt exist.
type PointsTo struct {
	Ref rt.RefEval
}

func (xr PointsTo) GetReference(r rt.Runtime) (ret rt.Reference, err error) {
	if ref, e := xr.Ref.GetReference(r); e != nil {
		err = e
	} else if obj, e := r.GetObject(ref); e != nil {
		err = e
	} else {
		ret = rt.Reference(obj.GetId())
	}
	return
}

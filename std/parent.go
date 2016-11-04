package std

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

// AssignTo provides a shortcut for ChangeParent
func AssignTo(src rt.ObjEval, rel string, dst rt.ObjEval) ChangeParent {
	return ChangeParent{src, rel, dst}
}

// ChangeParent changes the "parent" of Src to Dst;
// after execution, Src will be a relative of Dst via the relation Rel.
type ChangeParent struct {
	Src rt.ObjEval
	Rel string
	Dst rt.ObjEval
}

// FIX: there's no error testing here ( ex. matching allowable properties, creating refrence loops, etc. ) its definitely possible to screw things up.
func (a ChangeParent) Execute(run rt.Runtime) (err error) {
	// FIX: would relation by value remove the need for transaction?
	if src, e := a.Src.GetObject(run); e != nil {
		err = e
	} else if dst, e := a.Dst.GetObject(run); e != nil {
		err = e
	} else {
		// note: objects which start out of world, dont have a parent to clear.
		if _, old, ok := run.LookupParent(src); ok {
			err = old.SetGeneric(rt.Object{})
		}
		if err == nil {
			if next, ok := src.FindProperty(a.Rel); !ok {
				err = errutil.New("ChangeParent:", src.GetId(), "does not have property", a.Rel)
			} else {
				next.SetGeneric(dst)
			}
		}
	}
	return
}

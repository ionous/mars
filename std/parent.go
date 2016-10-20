package std

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

// FIX - i think this will eventually be a machine that is passed in, and we will replace LookupParent entirely
type ChangeParent struct {
	Src rt.RefEval
	Rel string
	Dst rt.RefEval
}

// FIX: there's no error testing here ( ex. matching allowable properties, creating refrence loops, etc. ) its definitely possible to screw things up.
func (a ChangeParent) Execute(r rt.Runtime) (err error) {
	// would relation by value remove the need for transaction?
	if ref, e := a.Src.GetReference(r); e != nil {
		err = e
	} else if src, e := r.GetObject(ref); e != nil {
		err = e
	} else if dst, e := a.Dst.GetReference(r); e != nil {
		err = e
	} else {
		panic("!!!")
		// if _, old, ok := r.LookupParent(src); ok {
		// 	// note: objects which start out of world, dont have an owner to clear.
		// 	old.GetValue().SetObject(ident.Empty())
		// }
		if next, ok := src.FindProperty(a.Rel); !ok {
			err = errutil.New("ChangeParent:", src.GetId(), "does not have property", a.Rel)
		} else {
			next.GetValue().SetObject(dst.Id())
		}
	}
	return
}

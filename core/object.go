package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

// has to be a struct because instance is interface.
// as an internal object it cant be serialized accidentally.
type ObjectScope struct {
	meta.Instance
}

func (obj ObjectScope) String() string {
	return obj.GetId().String()
}

// FindValue implements Scope
func (obj ObjectScope) FindValue(name string) (ret meta.Generic, err error) {
	if prop, ok := obj.FindProperty(name); !ok {
		err = errutil.New("object property unknown", obj, name)
	} else {
		switch prop.GetType() {
		case meta.NumProperty:
			ret = prop.GetGeneric().(rt.NumEval)
		case meta.TextProperty:
			ret = prop.GetGeneric().(rt.TextEval)
		case meta.ObjectProperty:
			ret = prop.GetGeneric().(rt.ObjEval)
		default:
			err = errutil.New("object property type unknown", obj, name)
		}
	}
	return
}

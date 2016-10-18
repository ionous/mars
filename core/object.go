package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

// has to be a struct because instance is interface.
type Object struct {
	meta.Instance
}

// FIX: maybe runtime will return this? ie. rt.GetObject(ref) instead of exposing the model
func MakeObject(r rt.Runtime, xr rt.Reference) (ret Object, err error) {
	if inst, ok := r.Model().GetInstance(xr.Id()); !ok {
		err = errutil.New("instance not found", xr.Id)
	} else {
		ret = Object{inst}
	}
	return
}

func (obj Object) String() string {
	return obj.GetId().String()
}

// FindValue implements Scope
func (obj Object) FindValue(name string) (ret rt.Value, err error) {
	if prop, ok := obj.FindProperty(name); !ok {
		err = errutil.New(obj.GetId(), "has no property", name)
	} else {
		switch prop.GetType() {
		case meta.NumProperty:
			ret = rt.Number(prop.GetValue().GetNum())
		case meta.TextProperty:
			ret = rt.Text(prop.GetValue().GetText())
		case meta.ObjectProperty:
			ret = rt.Reference(prop.GetValue().GetObject())
		default:
			err = errutil.New(obj.GetId(), "unknown property type", name)
		}
	}
	return
}

package std

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

type Location struct {
	obj rt.ObjEval
}

func (l Location) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if obj, e := l.obj.GetObject(run); e != nil {
		err = e
	} else if p, ok := (Locator{run}._location(obj)); !ok {
		err = errutil.New("object is nowhere", obj)
	} else {
		ret = rt.Object{p}
	}
	return
}

type Locator struct {
	rt.Runtime
}

func (loc Locator) _location(obj meta.Instance) (parent meta.Instance, okay bool) {
	if room, ok := loc.get(obj, "whereabouts"); ok {
		parent, okay = room, true
	} else if wearer, ok := loc.get(obj, "wearer"); ok {
		parent, okay = loc._location(wearer)
	} else if owner, ok := loc.get(obj, "owner"); ok {
		parent, okay = loc._location(owner)
	} else if supporter, ok := loc.get(obj, "support"); ok {
		parent, okay = loc._location(supporter)
	} else if container, ok := loc.get(obj, "enclosure"); ok {
		parent, okay = loc._location(container)
	}
	return
}

func (loc Locator) get(obj meta.Instance, where string) (ret meta.Instance, okay bool) {
	// fix: use faster lookup?
	if w, ok := obj.FindProperty(where); ok {
		if id := w.GetValue().GetObject(); !id.Empty() {
			if inst, e := loc.GetObject(id); e == nil {
				ret, okay = inst, true
			}
		}
	}
	return
}

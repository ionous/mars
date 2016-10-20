package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

type Location struct {
	ref rt.RefEval
}

func (l Location) GetReference(run rt.Runtime) (ret rt.Reference, err error) {
	if ref, e := l.ref.GetReference(run); e != nil {
		err = e
	} else if obj, e := run.GetObject(ref); e != nil {
		err = e
	} else if p, ok := (Locator{run}._location(obj)); !ok {
		ret = core.NullRef()
	} else {
		ret = rt.Reference(p.GetId())
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
			if inst, e := loc.GetObject(rt.Reference(id)); e == nil {
				ret, okay = inst, true
			}
		}
	}
	return
}

package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
)

type Location struct {
	obj rt.ObjEval
}

func (l Location) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if obj, e := l.obj.GetObject(run); e != nil {
		err = e
	} else {
		run := scope.Make(run, scope.NewObjectScope(obj), run)
		if where, e := l.locate(run); e != nil {
			err = e
		} else if where.Empty() {
			err = errutil.New("object is nowhere", obj)
		} else {
			ret = where
		}
	}
	return
}

func (loc Location) locate(run rt.Runtime) (ret rt.Object, err error) {
	props := []string{"whereabouts", "wearer", "owner", "support", "enclosure"}
	for _, p := range props {
		if x, e := loc.get(run, p); e != nil {
			err = e
			break
		} else if !x.Empty() {
			ret = x
			break
		}
	}
	return
}

func (loc Location) get(run rt.Runtime, where string) (rt.Object, error) {
	return (GetObject{where}).GetObject(run)
}

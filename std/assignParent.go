package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

func clearRef(src rt.ObjEval, name string) core.SetObj {
	return core.SetObj{name, src, core.Nothing()}
}

type AssignParent struct {
	Src rt.ObjEval
	Rel OwnerRelation
	Dst rt.ObjEval
}

type OwnerRelation interface {
	GetRelation() string
}

type Wearer struct{}

func (_ Wearer) GetRelation() string { return "wearer" }

type Owner struct{}

func (_ Owner) GetRelation() string { return "owner" }

type Whereabouts struct{}

func (_ Whereabouts) GetRelation() string { return "whereabouts" }

type Support struct{}

func (_ Support) GetRelation() string { return "support" }

type Enclosure struct{}

func (_ Enclosure) GetRelation() string { return "enclosure" }

func (ap AssignParent) Execute(run rt.Runtime) (err error) {
	if ap.Rel == nil {
		err = errutil.New("relation not set")
	} else {
		rel := ap.Rel.GetRelation()
		if e := ap.clear(run, rel); e != nil {
			err = e
		} else {
			err = core.SetObj{string(rel), ap.Src, ap.Dst}.Execute(run)
		}
	}
	return
}

func (ap AssignParent) clear(run rt.Runtime, rel string) (err error) {
	for _, name := range []string{"wearer", "owner", "whereabouts", "support", "enclosure"} {
		if name != rel {
			if e := clearRef(ap.Src, name).Execute(run); e != nil {
				err = e
				break
			}
		}
	}
	return
}

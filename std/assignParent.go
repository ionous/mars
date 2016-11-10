package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func clearRef(src rt.ObjEval, name string) core.SetObj {
	return core.SetObj{core.PropertyRef{name, src}, core.NullRef()}
}

func AssignParent(src rt.ObjEval, rel string, dst rt.ObjEval) (ret core.ExecuteList) {
	for _, x := range []string{"wearer", "owner", "whereabouts", "support", "enclosure"} {
		if x != rel {
			ret = append(ret, clearRef(src, x))
		}
	}
	return append(ret, core.SetObj{core.PropertyRef{rel, src}, dst})
}

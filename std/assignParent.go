package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/source/types"
)

func clearRef(src rt.ObjEval, name string) core.SetObj {
	return core.SetObj{types.NamedProperty(name), src, core.Nothing()}
}

func AssignParent(src rt.ObjEval, rel string, dst rt.ObjEval) core.ExecuteList {
	var ret []rt.Execute
	for _, x := range []string{"wearer", "owner", "whereabouts", "support", "enclosure"} {
		if x != rel {
			ret = append(ret, clearRef(src, x))
		}
	}
	return core.ExecuteList{append(ret, core.SetObj{types.NamedProperty(rel), src, dst})}
}

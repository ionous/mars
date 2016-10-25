package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Carrier(obj string) (ret rt.ObjEval) {
	our := core.Id(obj)
	wearer := core.RefProperty{our, "wearer"}
	owner := core.RefProperty{our, "owner"}
	return core.ChooseRef{
		If:   core.IsValid{wearer},
		True: wearer,
		False: core.ChooseRef{
			If:   core.IsValid{owner},
			True: owner,
		},
	}
}

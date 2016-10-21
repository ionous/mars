package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Carrier(obj string) (ret rt.RefEval) {
	our := core.R(obj)
	wearer := core.RefProperty{our, "wearer"}
	owner := core.RefProperty{our, "owner"}
	return core.ChooseRef{
		If:   core.Exists{wearer},
		True: wearer,
		False: core.ChooseRef{
			If:   core.Exists{owner},
			True: owner,
		},
	}
}

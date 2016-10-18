package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Carrier(obj string) (ret rt.RefEval) {
	our := core.R(obj)
	wearer := core.Property{our, "wearer"}
	owner := core.Property{our, "owner"}
	return core.ChooseRef{
		If:   core.Exists{wearer},
		True: wearer,
		False: core.ChooseRef{
			If:   core.Exists{owner},
			True: owner,
		},
	}
}

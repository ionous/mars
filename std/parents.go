package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/g"
)

func Carrier(obj string) (ret rt.ObjEval) {
	wearer := g.The(obj).Object("wearer")
	owner := g.The(obj).Object("owner")
	return core.ChooseRef{
		If:   core.IsValid{wearer},
		True: wearer,
		False: core.ChooseRef{
			If:   core.IsValid{owner},
			True: owner,
		},
	}
}

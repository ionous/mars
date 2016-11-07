package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/std/compat"
)

// FIX: note: this wouldnt work for something in a container
func Carrier(r compat.ScriptRef) compat.ScriptRef {
	wearer := r.Object("wearer")
	owner := r.Object("owner")
	return compat.ScriptRef{core.ChooseRef{
		If:   core.IsValid{wearer},
		True: wearer,
		False: core.ChooseRef{
			If:   core.IsValid{owner},
			True: owner,
		},
	}}
}

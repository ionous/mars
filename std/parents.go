package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/core/stream"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/std/compat"
)

// FIX: note: this wouldnt work for something in a container
func Carrier(r compat.ScriptRef) compat.ScriptRef {
	wearer := r.Object("wearer")
	owner := r.Object("owner")
	return compat.ScriptRef{core.ChooseObj{
		If:   core.IsValid{wearer},
		True: wearer,
		False: core.ChooseObj{
			If:   core.IsValid{owner},
			True: owner,
		},
	}}
}

// the first room, closed container, or empty parent.
func Enclosure(obj rt.ObjEval) compat.ScriptRef {
	// a search through all releation properties:
	refs := rt.References{}
	for _, x := range []string{"wearer", "owner", "whereabouts", "support", "enclosure"} {
		refs = append(refs, g.TheObject().Object(x))
	}
	//
	return compat.ScriptRef{
		stream.First{
			In: stream.Generate{
				First: obj,
				Next: stream.First{
					In:       refs,
					Matching: g.TheObject().Exists(),
					Else:     core.NullRef(),
				},
			},
			Matching: core.Any(g.TheObject().FromClass("rooms"),
				core.All(g.TheObject().FromClass("container"), g.TheObject().Is("closed"))),
			Else: core.NullRef(),
		},
	}
}

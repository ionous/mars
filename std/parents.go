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

// Enclosure returns the first room, closed container, or empty parent.
func Enclosure(obj rt.ObjEval) compat.ScriptRef {
	return compat.ScriptRef{
		stream.First{
			In: Ancestors(obj),
			Matching: core.Any(g.TheObject().FromClass("rooms"),
				core.All(g.TheObject().FromClass("container"), g.TheObject().Is("closed"))),
			Else: core.Nothing(),
		},
	}
}

func Parent(obj rt.ObjEval) rt.ObjEval {
	return stream.First{In: Ancestors(obj), Else: core.Nothing()}
}

// Ancestors returns a stream of parent objects, starting from the passed object's parent, and moving upwards from there.
func Ancestors(obj rt.ObjEval) rt.ObjListEval {
	// a search through all relation properties:
	refs := rt.References{}
	for _, rel := range []string{"wearer", "owner", "whereabouts", "support", "enclosure"} {
		refs = append(refs, core.PropertySafeRef{rel, core.GetObj{"@"}})
	}
	return stream.MakeStream{
		Using: obj,
		Next: stream.First{
			In:       refs,
			Matching: g.TheObject().Exists(),
		},
	}
}

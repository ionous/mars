package facts

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

// these are the same:
// fact.Is("recollected")
// if facts.PlayerRecollects(g, "exeunt") {
// Is{Id("exeunt"), "recollected"}

// 		if facts.PlayerLearns(g, "everyone-is") {
// functions are "shortcuts" which use existing dl structs to "encapsulate" ( combine ) behavior
// if you want to extend behavior -- prefer making shortcuts over new dl elements
// ultimatey, both are okay.

func PlayerLearns(fact string) rt.BoolEval {
	return Choose{If: Not{Is{Id(fact), "recollected"}},
		True: Change(Id(fact)).To("recollected"),
	}
}

func PlayerRecollects(fact string) rt.BoolEval {
	return Is{Id(fact), "recollected"}
}

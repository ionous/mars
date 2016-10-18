package facts

import (
	. "github.com/ionous/mars/core"
)

// these are the same:
// fact.Is("recollected")
// if facts.PlayerRecollects(g, "exeunt") {
// Is{R("exenut"), "recollected"}

// 		if facts.PlayerLearns(g, "everyone-is") {
// functions are "shortcuts" which use existing dl structs to "encapsulate" ( combine ) behavior
// if you want to extend behavior -- prefer making shortcuts over new dl elements
// ultimatey, both are okay.

func PlayerLearns(fact string) rt.BoolEval {
	return Choose{If: Not{Is{R(fact), "recollected"}},
		True: Change(R(fact)).To("recollected"),
	}
}

func PlayerRecollects(fact string) rt.BoolEval {
	return Is{R(fact), "recollected"}
}

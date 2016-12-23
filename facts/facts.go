package facts

import (
	"github.com/ionous/mars"
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/inbuilt"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
)

// FIX: these should be script functions
// script functions should be
func PlayerLearns(fact string) rt.BoolEval {
	return Choose{If: IsNot{IsState{RawId(fact), "recollected"}},
		True: Change(RawId(fact)).To("recollected"),
	}
}

func PlayerRecollects(fact string) rt.BoolEval {
	return IsState{RawId(fact), "recollected"}
}

// Fact contains all of mar's built-in commands and primitives.
func Facts() *mars.Package {
	if facts == nil {
		facts = &mars.Package{
			Name: "Facts",
			// MARS, FIX: move "kinds" declaration to a custom backend script?
			Scripts:      pkg.Scripts,
			Tests:        pkg.Tests,
			Dependencies: mars.Dependencies(inbuilt.Inbuilt(), Package()),
			Commands:     (*FactCommands)(nil),
			Interfaces:   (*FactInterfaces)(nil),
		}
	}
	return facts
}

var facts *mars.Package
var pkg mars.PackageBuilder

type FactInterfaces struct {
}

type FactCommands struct {
}

func init() {
	pkg.AddScript("Facts", scriptFacts())
}

func scriptFacts() (s Script) {
	s.The("kinds", Called("facts"),
		// FIX: interestingly, kinds should have names
		// also: having the same property as a parent class probably shouldnt be an error
		Have("summary", "text"))
	// FIX: many-to-many doesnt exist; traversing a manually created table of all actors and facts would be fairly heavy; so just using a flag.
	s.The("facts", AreEither("recollected").Or("not recollected").Usually("not recollected"))
	return
}

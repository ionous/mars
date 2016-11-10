package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	addScript("PuttingItOn", // 1. source
		The("actors",
			// FIX? word-wise this is wrong ( see tickle-it-with, though it is "correct" )
			Can("put it onto").And("putting it onto").RequiresOne("supporter").AndOne("prop"),
			To("put it onto", g.ReflectWithContext("report put")),
			//  can't put clothes being worn
			Before("putting it onto").Always(
				Choose{
					If: g.The("action.Context").Object("wearer").Exists(),
					True: g.Go(g.Say("You can't put worn clothing."),
						g.StopHere(),
					),
				}),
			//  can't put what isn't held
			Before("putting it onto").Always(
				Choose{
					If: Carrier(g.The("action.Context")).Equals(g.The("action.Source")),
					False: g.Go(
						g.Say("You aren't holding", g.The("action.Context").Lower(), "."),
						g.StopHere(),
					),
				}),
			//  can't put something onto itself
			Before("putting it onto").Always(
				Choose{
					If: g.The("action.Target").Equals(g.The("action.Context")),
					True: g.Go(
						g.Say("You can't put something onto itself."),
						g.StopHere(),
					),
				}),
			//  can't put onto closed supporters
			Before("putting it onto").Always(
				Choose{
					If: g.The("action.Target").Is("closed"),
					True: g.Go(
						g.Say(g.The("action.Target").Upper(), "is closed."),
						g.StopHere(),
					),
				}),
		),
		// 2. supporters
		The("supporters",
			Can("report put").And("reporting put").RequiresOne("prop").AndOne("actor"),
			To("report put", g.ReflectWithContext("report placed"))),
		// 3. context
		The("props",
			Can("report placed").And("reporting placed").RequiresOne("actor").AndOne("supporter"),
			To("report placed",
				g.Say("You put", g.The("action.Source").Lower(), "onto", g.The("action.Context").Lower(), "."),
				Put(g.The("action.Source")).Onto(g.The("action.Context")),
			)),
		Understand("put {{something else}} on|onto {{something}}").
			And("drop|throw|discard {{something else}} on|onto {{something}}").
			As("put it onto"),
	)
}

func Put(prop rt.ObjEval) PutOntoPhrase {
	return PutOntoPhrase{prop}
}

func (p PutOntoPhrase) Onto(supporter rt.ObjEval) rt.Execute {
	// FIX: validate that the supporter is a supporter?
	return AssignParent(p.prop, "support", supporter)
}

type PutOntoPhrase struct {
	prop rt.ObjEval
}

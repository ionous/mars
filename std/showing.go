package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// all infom showing rules:
// 	"applies to one carried thing and one visible thing"
//  "you can't show what you haven't got"
//  	why does it need to do this, since it already applies to something carried?
//  "convert show to yourself to examine, tests if the actors are the same, and calls "convert examining""
//  	why it needs a special "convert" function?
//  "block showing - says: "the actor is unimpressed""
// 		why not an explicit report?
func init() {
	addScript("Showing",
		// 1. source
		The("actors",
			Can("show it to").And("showing it to").RequiresOne("actor").AndOne("prop"),
			To("show it to", g.ReflectWithContext("report show")),
			// "you can't show what you haven't got"
			Before("showing it to").Always(
				Choose{
					If: Carrier(g.The("prop")).Equals(g.The("action.Source")),
					False: g.Go(
						g.Say("You aren't holding", g.The("prop").Lower()),
						g.StopHere(),
					),
				}),
			// "convert show to yourself to examine"
			Before("showing it to").Always(
				Choose{
					If: g.The("action.Target").Equals(g.The("action.Source")),
					True: g.Go(
						g.The("action.Target").Go("examine it", g.The("prop")),
						g.StopHere(),
					),
				}),
		),
		// 2. receiver
		The("actors",
			Can("report show").And("reporting show").RequiresOne("prop").AndOne("actor"),
			To("report show", g.ReflectWithContext("report shown"))),

		// 3. context
		The("props",
			Can("report shown").And("reporting shown").RequiresTwo("actor"),
			To("report shown",
				g.The("action.Context").Go("impress"),
			)),
		// input
		Understand("show|present|display {{something}} {{something else}}").
			And("show|present|display {{something else}} to {{something}}").
			As("show it to"),
	)
}

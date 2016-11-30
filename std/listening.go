package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// From Inform7:
// Listening to something (past tense listened to) -- The Standard Rules define this action in only a minimal way, replying only that the player hears nothing unexpected.
//
// Typed commands leading to this action:
// 		"listen"
// 		"hear [something]"
// 		"listen to [something]"
//
// Rules controlling this action:
// 		report "an actor listening to" report listening rule
//    	 A   "[We] [hear] nothing unexpected."
//       B   "[The actor] [listen]."
//
// To override: report listening rule response (A) is "New text.".
//
func init() {
	pkg.AddScript("Listening",
		The("actors",
			Can("listen").And("listening").RequiresNothing(),
			To("listen", g.ReflectToLocation("report listen"))),
		The("actors",
			Can("listen to").And("listening to").RequiresOnly("kind"),
			To("listen to", g.ReflectToTarget("report listen"))),
		// kinds, to allow rooms and objects
		The("kinds",
			Can("report listen").And("reporting listen").RequiresOnly("actor"),
			To("report listen",
				Choose{
					If:    g.The("player").Equals(g.The("action.Target")),
					True:  g.Say("You hear nothing unexpected."),
					False: g.Say(g.The("actor").Upper(), "listens."),
				})),
		Understand("listen to {{something}}").And("listen {{something}}").As("listen to"),
		Understand("listen").As("listen"),
	)
}

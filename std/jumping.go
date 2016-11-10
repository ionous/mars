package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	addScript("Jumping",
		The("actors",
			Can("jump").And("jumping").RequiresNothing(),
			To("jump", g.ReflectToLocation("report jump")),
		),

		// kinds, to allow rooms and objects
		The("kinds",
			Can("report jump").And("reporting jump").RequiresOne("actor"),
			To("report jump",
				// FIX? inform often, but not always, tests for trying silently,
				// "if the action is not silent" ...
				// seems... strange. why report if if its silent?
				Choose{
					If:    g.The("player").Equals(g.The("action.Target")),
					True:  g.Say("You jump on the spot."),
					False: g.Say(g.The("action.Target").Upper(), "jumps on the spot."),
				})),

		Understand("jump|skip|hop").As("jump"),
	)
}

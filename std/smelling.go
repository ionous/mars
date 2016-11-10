package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	addScript("Smelling",
		The("actors",
			Can("smell").And("smelling").RequiresNothing(),
			To("smell", g.ReflectToLocation("report smell")),

			Can("smell it").And("smelling it").RequiresOne("kind"),
			To("smell it", g.ReflectToTarget("report smell")),
		),

		// kinds, to allow rooms and objects
		The("kinds",
			Can("report smell").And("reporting smell").RequiresOne("actor"),
			To("report smell", Choose{
				If:    g.The("player").Equals(g.The("action.Target")),
				True:  g.Say("You smell nothing unexpected."),
				False: g.Say(g.The("action.Target").Upper(), "sniffs."),
			}),
		),

		Understand("smell|sniff {{something}}").As("smell it"),
		Understand("smell|sniff").As("smell"),
	)
}

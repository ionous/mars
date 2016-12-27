package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// you are carrying:
// you are wearing:
// fwiw: Carry out taking inventory have the only good description of response text. (A)
func init() {
	pkg.AddScript("Inventory",
		// "taking inventory" in inform
		// again, as with some other actions: for players this happens in carry out, for npcs in report.
		// i'm sure that's useful... somehow....
		The("actors",
			Can("report inventory").And("reporting inventory").RequiresNothing(),
			To("report inventory", invList("clothing"), invList("inventory")),
		),

		// FIX: for some reason, the order must be biggest match to smallest, the other way doesnt work.
		Understand("inventory|inv|i").As("report inventory"),
	)
}
func invList(s string) (ret rt.Execute) {
	return ForEachObj{
		In:   g.The("actor").ObjectList(s),
		Else: g.Go(g.Say(s, ": none.")),
		Go: g.Go(
			Choose{
				If:   GetBool{"@first"},
				True: g.Go(g.Say(s, ":")),
			},
			g.TheObject().Go("print name")),
	}
}

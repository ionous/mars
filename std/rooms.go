package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	pkg.AddScript("Rooms",
		// vs. descriptions in "kind"
		// it seems to make sense for now to have two separate description fields.
		// rooms like to say their description, while objects like to say their brief initial appearance ( or name, if there's none. )
		The("rooms",
			Have("description", "text")),

		// inform's rooms: lighted, dark; unvisited, visited; description, region
		// the class hierarchy means rooms cant contain other rooms.
		The("rooms",
			HaveMany("contents", "objects").
				Implying("objects", HaveOne("whereabouts", "room"))),

		// inform's rooms: lighted, dark; unvisited, visited; description, region
		The("kinds",
			Called("rooms"),
			AreEither("visited").Or("unvisited").Usually("unvisited"),
		),

		The("rooms",
			Can("report the view").And("reporting the view").RequiresNothing(),
			After("reporting the view").Always(
				Change(g.The("room")).To("visited"),
			),
			To("report the view",
				g.Say(g.The("room").Text("name")),
				g.Say(g.The("room").Text("description")),
				ForEachObj{
					In: g.The("room").ObjectList("contents"),
					Go: g.Go(g.TheObject().Go("print description")),
				},
			)),
	)
}

package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/std/compat"
)

func init() {
	addScript("Rooms",
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
				g.Go(Change(g.The("room")).To("visited")),
			),
			To("report the view",
				g.Go(ViewRoom(g.The("room"))),
			)),
	)
}

// FUTURE? interestingly, while we wouldnt be able to encode them without special work, the contents of the phrases are fixed: we could have After("reporting").Execute(Phrase). maybe "standard" phrases could put themselves in some sort of wrapper? around the model? tho not quite sure howd that work.
func ViewRoom(obj compat.ScriptRef) rt.Execute {
	return g.Go(
		g.Say(obj.Text("name")),
		g.Say(obj.Text("description")),
		ForEachObject{
			In: obj.ObjectList("contents"),
			Go: g.TheObject().Go("print description"),
		},
	)
}

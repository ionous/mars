package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	pkg.AddScript("Containers",
		The("openers",
			Called("containers"),
			HaveMany("contents", "objects").
				Implying("objects", HaveOne("enclosure", "container")),
			AreEither("opaque").Or("transparent"),
			AreEither("lockable").Or("not lockable").Usually("not lockable"),
			AreEither("locked").Or("unlocked").Usually("unlocked"),
		),
		// FIX: conditional return instead of Always?
		// or some way ( dependency injection ) to get at the event object
		// of course, rules producing values and stacks might work too.
		// FIX: a container is an opener... where do we print the opener status name
		// put this on doors for now.
		The("containers",
			When("printing name text").
				Always(
					g.Say(
						g.The("container").Upper(),
						ChooseText{
							If: g.The("container").Is("closed"),
							True: ChooseText{
								If:   g.The("container").Is("hinged"),
								True: T(" (closed)"),
							},
							False: ChooseText{
								If: All(
									g.The("container").ObjectList("contents").Empty(),
									Any(
										g.The("container").Is("transparent"),
										g.The("container").Is("open")),
								),
								True: T(" (empty)"),
							},
						},
					),
					g.StopHere(),
				)),
	)
}

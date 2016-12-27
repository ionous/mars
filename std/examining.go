package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	pkg.AddScript("Examining", // one visible thing
		// examine studio: You can't see any such thing; sad face.
		The("actors",
			Can("examine it").And("examining it").RequiresOnly("object"),
			To("examine it", g.ReflectToTarget("be examined")),
		),
		// the default action prints the place holder text
		// the events system prints the specifics and prevents the defaults as needed
		// users can override for a particular object the entire thing
		The("objects",
			Can("be examined").And("being examined").RequiresOnly("actor"),
			To("be examined",
				g.The("object").Go("print details"),
				g.The("object").Go("print contents"),
			)),

		The("objects",
			Can("print details").And("printing details").RequiresNothing(),
			To("print details",
				Choose{
					If:    IsEmpty{g.The("object").Text("description")},
					False: g.Go(g.Say(g.The("object").Text("description"))),
					True:  g.Go(g.The("object").Go("print name")),
				},
			)),

		The("containers",
			When("printing contents").Always(
				Choose{
					If: Any(g.The("container").Is("open"),
						g.The("container").Is("transparent")),
					True: g.Go(ForEachObj{
						In: g.The("container").ObjectList("contents"),
						Go: g.Go(
							Choose{
								If:   GetBool{"@first"},
								True: g.Go(g.Say("In", g.The("container").Lower(), ":")),
							},
							g.TheObject().Go("print description"),
						),
					}),
				},
			),
		),
		// report an actor examining:
		// where best to do that switch?
		// carry out in inform seems to be limited to the player;....
		///	if the actor is not the player:
		//	say "[The actor] [look] closely at [the noun]." (A).
		The("supporters",
			When("printing contents").Always(
				ForEachObj{
					In: g.The("supporter").ObjectList("contents"),
					Go: g.Go(
						Choose{
							If:   GetBool{"@first"},
							True: g.Go(g.Say("On", g.The("supporter").Lower(), ":")),
						},
						g.TheObject().Go("print description"),
					),
				}),
		),
		Understand("examine|x|watch|describe|check {{something}}").
			And("look|l {{something}}").
			And("look|l at {{something}}").As("examine it"),
	)
}

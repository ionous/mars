package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	pkg.AddScript("Eating",
		The("actors",
			Can("eat it").And("eating it").RequiresOnly("prop"),
			To("eat it", g.ReflectToTarget("report eat")),
		),

		The("props", AreEither("inedible").Or("edible")),

		The("props",
			Can("report eat").And("reporting eat").RequiresOnly("actor"),
			To("report eat",
				Choose{
					If:    g.The("prop").Is("inedible"),
					True:  g.Say("That's not something you can eat."),
					False: g.The("actor").Go("impress"),
				}),
		),
		Understand("eat {{something}}").As("eat it"),
	)
}

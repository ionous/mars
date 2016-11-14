package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	addScript("Kissing",
		// kissing
		The("actors",
			Can("kiss it").And("kissing it").RequiresOnly("object"),
			To("kiss it", g.ReflectToTarget("report kiss")),
			//  kissing yourself rule
			Before("kissing it").Always(
				Choose{
					If: g.The("action.Source").Equals(g.The("action.Target")),
					True: g.Go(
						g.Say(g.The("actions.Source").Upper(), "didn't get much from that."),
						g.StopHere(),
					),
				})),
		//  block kissing rule
		The("objects",
			Can("report kiss").And("reporting kiss").RequiresOnly("actor"),
			To("report kiss",
				g.Say(g.The("action.Source").Upper(), "might not like that."),
			)),

		Understand("kiss|hug|embrace {{something}}").As("kiss it"),
	)
}

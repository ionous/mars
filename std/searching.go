package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// searching: requiring light; FIX: what does searching a room do?
func init() {
	pkg.AddScript("Searching",
		The("actors",
			Can("search it").And("searching it").RequiresOnly("prop"),
			To("search it", g.ReflectToTarget("report search"))),
		The("props",
			Can("report search").And("reporting search").RequiresOnly("actor"),
			To("report search",
				g.Say("You find nothing unexpected."))),

		// WARNING/FIX: multi-word statements must appear before their single word variants
		// ( or the parser will attempt to match the setcond word as a noun )
		Understand("search {{something}}").
			And("look inside|in|into|through {{something}}").
			As("search it"))
}

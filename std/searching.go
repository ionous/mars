package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// searching: requiring light; FIX: what does searching a room do?
var Searching = Script(
	The("actors",
		Can("search it").And("searching it").RequiresOne("prop"),
		To("search it", g.ReflectToTarget("report search"))),
	The("props",
		Can("report search").And("reporting search").RequiresOne("actor"),
		To("report search",
			g.Say("You find nothing unexpected."))),

	// WARNING/FIX: multi-word statements must appear before their single word variants
	// ( or the parser will attempt to match the setcond word as a noun )
	Understand("search {{something}}").
		And("look inside|in|into|through {{something}}").
		As("search it"))

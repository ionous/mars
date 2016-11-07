package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

var LookUnder = Script(
	// NOTE: inform has two entries for some actions (looking under as an example, jumping as a counter example):
	// 1. carry out an actor looking under: if the player
	// 2. report an actor looking under: if not the player
	// its not exactly clear to me why, the docs give guidelines for this, but not rationale.
	// it might be interesting to queue says, if they need to be cancelled or held back.
	// keep in mind, most of these really want to be animations, and only sometimes voice.

	// one visible thing and requiring light.
	The("actors",
		Can("look under it").And("looking under it").RequiresOne("object"),
		To("look under it", g.ReflectToTarget("report look under")),
	),
	Understand("look under {{something}}").As("look under it"),
)

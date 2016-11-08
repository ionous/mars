package std

import (
	. "github.com/ionous/mars/script"
)

// FIX: the player should really be a point of view object.
var Player = Script(
	The("actor",
		Called("the player"),
		Exists(),
		Is("scenery"),
	))

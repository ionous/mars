package std

import (
	. "github.com/ionous/mars/script"
)

//
var StatusBar = Script(The("globals",
	Called("status bar instances"),
	Have("left", "text"),
	Have("right", "text")),

	The("status bar instance",
		Called("status bar"),
		Exists()),

	// FIX:
	// The("rooms",
	// 		Can("report the view").And("reporting the view").RequiresNothing(),
	// 		When("reporting the view").Always(
	// 			room := g.The("room")
	// 			g.The("status bar").SetText("left", g.The("room").Upper())
	// 		}),

)

package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
	. "github.com/ionous/mars/std/script"
)

func init() {
	addScript("Looking",
		// one visible thing, and requiring light
		The("actors",
			Can("look").And("looking").RequiresNothing(),
			// note: reflect to location send the actor as a parameter,
			// but report the view doesn't expect parameters.
			To("look",
				g.Our("actor").Object("whereabouts").Go("report the view"),
			),
		),
		Understand("look|l").As("look"),
	)

	addTest("Looking",
		test.Setup(
			The("actor", Called("the player"), Exists(), In("the memories")),
			The("room", Called("memories"), Has("description", "You are trapped in your own unconsciousness.")),
		).Try("looking at the room",
			test.Parse("look").
				Match("memories", "You are trapped in your own unconsciousness.").
				Expect(
					g.The("player").Object("whereabouts").Equals(g.The("memories")),
					g.The("memories").ObjectList("contents").Contains(g.The("player")),
				)),
	)
}

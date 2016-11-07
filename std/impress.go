package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
)

var Impress = The("actors",
	Can("impress").And("impressing").RequiresNothing(),
	To("impress", g.Say(g.The("actor").Upper(), "is unimpressed.")))

var ImpressTest = test.Suite{"Impress",
	test.Setup(
		The("actor", Called("the player"), Exists()),
	),
	test.Trials(
		test.Run("impress", g.The("player")).
			Match("The player is unimpressed.").
			Expect(),
	),
}

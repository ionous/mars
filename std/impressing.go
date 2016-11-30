package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
)

func init() {
	pkg.AddScript("Impressing",
		The("actors",
			Can("impress").And("impressing").RequiresNothing(),
			To("impress", g.Say(g.The("actor").Upper(), "is unimpressed."))))

	pkg.AddTest("Impressing",
		test.Setup(
			The("actor", Called("the player"), Exists()),
		).Try("to impress",
			test.Run("impress", g.The("player")).
				Match("The player is unimpressed."),
		),
	)
}

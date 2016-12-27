package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/lang"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	pkg.AddScript("Attacking",
		The("actors",
			Can("attack it").And("attacking it").RequiresOnly("object"),
			To("attack it", g.ReflectToTarget("report attack"))),

		The("objects",
			Can("report attack").And("reporting attack").RequiresOnly("actor"),
			To("report attack", Choose{
				If:   g.Our("player").Equals(g.Our("actor")),
				True: g.Go(g.Say("Violence isn't the answer.")),
			})),

		Understand("attack|break|smash|hit|fight|torture {{something}}").
			And("wreck|crack|destroy|murder|kill|punch|thump {{something}}").
			As("attack it"),
	)

	// pkg.AddTest("AttackingTest",
	// 	test.Setup(
	// 		The("object", Called("the wardrobe"), Exists()),
	// 		The("actor", Called("the player"), Exists()),
	// 		The("actor", Called("the lion"), Exists()),
	// 	).Try(
	// 		test.Parse("look under the wardrobe").
	// 			Match("You find nothing of interest."),
	// 		test.Execute(
	// 			g.The("lion").Go("look under it", g.The("wardrobe"))).
	// 			Match("The lion looks under the wardrobe."),
	// 	),
	// )
}

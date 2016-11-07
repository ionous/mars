package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
)

var Wearing = Script(
	The("actors",
		Can("wear it").And("wearing it").RequiresOne("prop"),
		To("wear it", g.ReflectToTarget("report wear")),
	),

	The("props",
		AreEither("wearable").Or("not wearable").Usually("not wearable")),

	The("props",
		Can("report wear").And("reporting wear").RequiresOne("actor"),
		To("report wear",
			Choose{
				If:    g.The("prop").Is("wearable"),
				False: g.Say("That's not something you can wear."),
				True: g.Go(
					Clothe("actor").With("prop"),
					g.Say("Now", g.The("actor").Lower(), "is wearing", g.The("prop").Lower(), "."),
				),
			}),
	),
	Understand("wear|don {{something}}").
		And("put on {{something}}").
		And("put {{something}} on").As("wear it"),
)

// Clothe provides a shortcut for the passed actor to wear some prop.
func Clothe(actor string) ClothePhrase {
	return ClothePhrase{g.The(actor)}
}

func (p ClothePhrase) With(clothing string) rt.Execute {
	return AssignTo(g.The(clothing), "wearer", p.actor)
}

type ClothePhrase struct {
	actor rt.ObjEval
}

var WearingTest = test.Suite{"Wearing",
	test.Setup(
		The("actor", Called("the player"), Exists()),
		The("prop", Called("the hat"), Is("wearable")),
		The("prop", Called("the cat"), Exists()),
	),
	test.Trials(
		test.Parse("don the hat").
			Match("Now the player is wearing the hat.").
			Expect(
				g.The("hat").Object("wearer").Equals(g.The("player")),
				g.The("player").ObjectList("clothing").Contains(g.The("hat"))),
		test.Parse("put the cat on").
			Match("That's not something you can wear.").
			Expect(
				IsNot{
					g.The("player").ObjectList("clothing").Contains(g.The("cat")),
				}),
	),
}

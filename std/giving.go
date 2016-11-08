package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
	. "github.com/ionous/mars/std/script"
)

// all infom giving rules:
// 	"applies to one carried thing and one visible thing."
//  "can't give what you haven't got"
//  "can't give to yourself"
//  "can't give to a non-person"
//  "can't give clothes being worn"
//  "block giving rule"
//  "the actor doesnt seem interested"
//  "can't exceed carrying capacity when giving"
//  "carry out giving something to"
//  "report an an actor giving something to"
var Giving = Script(
	// for summarily ( client side ) rejecting items
	The("actors", AreEither("items receiver").Or("items rejector")),

	The("actors",
		Can("acquire it").And("acquiring it").RequiresOne("prop"),
		To("acquire it", g.ReflectToTarget("be acquired"))),

	The("props",
		Can("be acquired").And("being acquired").RequiresOne("actor"),
		To("be acquired",
			AssignParent(g.The("prop"), "owner", g.The("actor")),
		)),

	// 1. source
	The("actors",
		Can("give it to").And("giving it to").RequiresOne("actor").AndOne("prop"),
		To("give it to", g.ReflectWithContext("report give")),
		// "convert give to yourself to examine"
		Before("giving it to").Always(
			Choose{
				If: g.The("action.Source").Equals(g.The("action.Target")),
				True: g.Go(
					g.Say("You can't give to yourself."),
					g.StopHere(),
				),
			}),
		// "can't give clothes being worn"
		Before("giving it to").Always(
			Choose{
				If: g.The("action.Context").Object("wearer").Exists(),
				True: g.Go(
					g.Say("You can't give worn clothing."),
					// FIX: try taking off the noun
					g.StopHere(),
				),
			}),
		// "you can't give what you haven't got"
		Before("giving it to").Always(
			Choose{
				If: Carrier(g.The("prop")).Equals(g.The("action.Source")),
				False: g.Go(
					g.Say("You aren't holding", g.The("prop").Lower(), "."),
					g.StopHere(),
				),
			}),
	),
	// 2. receiver
	The("actors",
		Can("report give").And("reporting give").RequiresOne("prop").AndOne("actor"),
		To("report give",
			g.ReflectWithContext("report gave"))),
	// 3. context
	The("props",
		Can("report gave").And("reporting gave").RequiresTwo("actor"),
		To("report gave",
			g.The("action.Context").Go("impress"))),
	// input
	Understand("give|pay|offer|feed {{something}} {{something else}}").
		And("give|pay|offer|feed {{something else}} to {{something}}").
		As("give it to"),
)

// Give, a shortcut for giving.
func Give(prop string) GivePropPhrase {
	return GivePropPhrase(prop)
}

func (prop GivePropPhrase) To(actor string) rt.Execute {
	// added indirection so we can transform props after the rules of taking/giving have run
	return g.The(actor).Go("acquire it", g.The(string(prop)))
}

type GivePropPhrase string

// MARS: move all tests to a sub-directory.
var GivingTest = test.NewSuite("Giving",
	test.Setup(
		The("actor", Called("the player"), Exists()),
		The("actor", Called("the firefighter"), Exists()),
		The("prop", Called("the cat"), Exists()),
	).Try(
		test.Parse("give the cat to the firefighter").
			Match("You aren't holding the cat."),
	),
	test.Setup(
		The("actor", Called("the player"), Exists()),
		The("actor", Called("the firefighter"), Exists()),
		The("prop", Called("the cat"), Exists()),
		The("prop", Called("the hat"), Exists()),
		The("player", Possesses("the cat")),
		The("player", Wears("the hat")),
	).Try(
		test.Parse("give the cat to the player").
			Match("You can't give to yourself."),
		test.Parse("give the hat to the firefighter").
			Match("You can't give worn clothing."),
		test.Parse("give the cat to the firefighter").
			Match("The firefighter is unimpressed."),
	),
)

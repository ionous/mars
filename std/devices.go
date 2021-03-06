package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// "Represents a machine or contrivance of some kind which can be switched on or off."
func init() {
	pkg.AddScript("Devices",
		// FIX: having problems with a lack of parts
		The("objects",
			AreEither("operable").Or("inoperable").Usually("inoperable")),

		The("props",
			Called("devices"),
			AreEither("switched off").Or("switched on")),

		The("devices",
			When("printing name text").Always(
				g.Say(g.The("device").Upper(),
					ChooseText{
						If: g.The("device").Is("operable"),
						True: ChooseText{
							If:    g.The("device").Is("switched on"),
							True:  T(" (switched on)"),
							False: T(" (switched off)"),
						},
					},
				),
				g.StopHere(),
			)),
		The("props", Can("report inoperable").And("reporting inoperable").RequiresNothing(),
			To("report inoperable",
				g.Say("It's inoperable."),
			)),

		//
		// Turn on, fix: was "prop", now "object" to handle outlet actors :(
		//
		The("actors",
			Can("switch it on").And("switching it on").RequiresOnly("object"),
			To("switch it on", g.ReflectToTarget("report switched on"))),

		The("objects",
			Can("report switched on").And("reporting switched on").RequiresOnly("actor"),
			To("report switched on",
				Choose{
					If:    g.The("action.source").Is("operable"),
					False: g.Go(g.The("action.source").Go("report inoperable")),
					True: g.Go(
						Choose{
							If: g.The("action.source").Is("switched on"),
							True: g.Go(
								g.The("action.source").Go("report already on", g.The("actor")),
							),
							False: g.Go(
								g.The("action.source").IsNow("switched on"),
								g.The("action.source").Go("report now on", g.The("actor")),
							),
						}),
				}),
			Can("report already on").And("reporting already on").RequiresOnly("actor"),
			To("report already on",
				g.Say("It's already switched on."),
			),
			Can("report now on").And("reporting now on").RequiresOnly("actor"),
			To("report now on",
				g.Say("Now", g.The("device").Lower(), "is on."),
			)),

		//
		// Turn off
		//
		The("actors",
			Can("switch it off").And("switching it off").RequiresOnly("prop"),
			To("switch it off", g.ReflectToTarget("report switch off"))),

		The("devices",
			Can("report switch off").And("reporting switch off").RequiresNothing(),
			To("report switch off",
				Choose{
					If: g.The("device").Is("switched off"),
					True: g.Go(
						g.The("device").Go("report already off", g.The("actor")),
					),
					False: g.Go(
						g.The("device").IsNow("switched off"),
						g.The("device").Go("report now off", g.The("actor")),
					),
				},
			),
			Can("report already off").And("reporting already off").RequiresOnly("actor"),
			To("report already off",
				g.Say("It's already off."), //[regarding the noun]?
			),
			Can("report now off").And("reporting now off").RequiresOnly("actor"),
			To("report now off",
				g.Say("Now", g.The("device").Lower(), "is off."),
			)),

		// understandings:
		// note: inform has "template Understand" here --
		// "switch [something switched on]" as switching off.
		// FIX:  inform's  "understand" has many meanings, but i think itd be better here
		// maybe: s.Understand.Or.As; Understand().WhenUnderstand("").Or()

		Understand("switch|turn on {{something}}").
			And("switch {{something}} on").As("switch it on"),

		Understand("turn|switch off {{something}}").
			And("turn|switch {{something}} off").As("switch it off"),
	)
}

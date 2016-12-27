package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// (note: the action uses props, so that stories can make any prop behave similar to an g.The("opener"). )
func init() {
	pkg.AddScript("OpeningClosing",
		The("props",
			Called("openers"),
			AreEither("open").Or("closed"),
			// note: not openable sounds like it cant become open, rather than it cannot be opened (by someone).
			AreEither("hinged").Or("hingeless"),
			AreEither("locakable").Or("not lockable").Usually("not lockable"),
			AreEither("unlocked").Or("locked"),
		),

		// Open:
		The("actors",
			Can("open it").And("opening it").RequiresOnly("prop"),
			To("open it", g.ReflectToTarget("be opened by")),
		),

		// "[regarding the noun][They] [aren't] something [we] [can] open."
		The("props",
			Can("be opened by").And("being opened by").RequiresOnly("actor"),
			To("be opened by",
				Choose{
					If: g.The("prop").Is("hinged"),
					False: g.Go(
						g.The("prop").Go("report unopenable", g.The("actor"))),
					True: g.Go(
						Choose{
							If: g.The("prop").Is("locked"),
							True: g.Go(
								g.The("prop").Go("report locked", g.The("actor"))),
							False: g.Go(
								Choose{
									If: g.The("prop").Is("open"),
									True: g.Go(
										g.The("prop").Go("report already open", g.The("actor"))),
									False: g.Go(
										Change(g.The("prop")).To("open"),
										g.The("prop").Go("report now open", g.The("actor"))),
								}),
						}),
				},
			),
			Can("report locked").And("reporting locked").RequiresOnly("actor"),
			To("report locked",
				// FIX? g.The("actor").Says("It's locked!"),
				g.Say("It's locked!"),
			),
			Can("report unopenable").And("reporting unopenable").RequiresOnly("actor"),
			To("report unopenable",
				g.Say("That's not something you can open."),
			),
			Can("report already open").And("reporting already opened").RequiresOnly("actor"),
			To("report already open",
				g.Say("It's already opened."),
			),
			Can("report now open").And("reporting now open").RequiresOnly("actor"),
			To("report now open",
				g.Say(g.The("opener").Upper(), "is now open."),
				// if the noun doesnt not enclose the actor
				// list the contents of the noun, as a sentence, tersely, not listing concealed items;
				// FIX? not all openers are opaque/transparent, and not all openers have contents.
				Choose{
					If: g.The("opener").Is("opaque"),
					True: g.Go(
						ForEachObj{
							In: g.The("opener").ObjectList("contents"),
							Go: g.Go(
								Choose{
									If: GetBool{"@first"},
									True: g.Go(
										g.Say("In", g.The("opener").Lower(), ":")),
								},
								g.Call("print description", GetObj{"@"}),
							),
						}),
				},
			),
		),

		// Close:
		// one visible thing, and requiring light
		The("actors",
			Can("close it").And("closing it").RequiresOnly("prop"),
			To("close it", g.ReflectToTarget("be closed by")),
		),
		The("props",
			Can("be closed by").And("being closed by").RequiresOnly("actor"),
			To("be closed by",
				Choose{
					If: g.The("prop").Is("hinged"),
					False: g.Go(
						g.The("prop").Go("report not closeable", g.The("actor"))),
					True: g.Go(
						Choose{ // FIX: locked?
							If: g.The("prop").Is("closed"),
							True: g.Go(
								g.The("prop").Go("report already closed", g.The("actor"))),
							False: g.Go(
								g.The("prop").IsNow("closed"),
								g.The("prop").Go("report now closed", g.The("actor"))),
						}),
				},
			),
			Can("report not closeable").And("reporting not closeable").RequiresOnly("actor"),
			To("report not closeable",
				g.Say("That's not something you can close."),
			),
			Can("report already closed").And("reporting already closed").RequiresOnly("actor"),
			To("report already closed",
				g.Say("It's already closed."), //[regarding the noun]?
			),
			Can("report now closed").And("reporting now closed").RequiresOnly("actor"),
			To("report now closed",
				g.Say("Now", g.The("prop").Lower(), "is closed."),
			),
		),

		// understandings:
		Understand("open {{something}}").As("open it"),
		Understand("close {{something}}").As("close it"),
	)
}

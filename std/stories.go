package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
	. "github.com/ionous/mars/std/script"
)

// FIX: we should have a "starting the turn" instead
// currently, to get the "frame" counter to see a different number than the last frame
// due to this frame's player input -- we have to increment both on end turn, and ending the story
var endTurn = Using{
	Object: g.The("story"),
	Run: g.Go(
		Inc{"turn count"},
		updateStatusBar,
	),
}

var updateStatusBar = Choose{
	If: g.The("story").Is("scored"),
	True: g.The("status bar").SetText("right",
		MakeText(GetNum{"score"}, "/", GetNum{"turn count"})),
}

// System actions
func init() {
	s := NewScript()
	// globals are transmitted to the client in the default view.
	s.The("kinds", Called("globals"), Exist())
	s.The("globals",
		Called("stories"), HasSingularName("story"),
		Have("author", "text"),
		Have("headline", "text"),
		AreEither("scored").Or("unscored").Usually("unscored"),
		// Inform uses global variables, which would be much nicer.
		// ex. The maximum score is 1.
		Have("score", "num"),
		Have("maximum score", "num"),
		Have("turn count", "num"),
		AreOneOf("playing", "completed", "starting").Usually("starting"),
	)

	s.The("stories",
		Can("commence").And("commencing").RequiresNothing(),
		Can("end the story").And("ending the story").RequiresNothing(),
		Can("end turn").And("ending the turn").RequiresNothing(),
		Before("commencing").Always(
			g.The("status bar").SetText("left", g.The("story").Upper()),
			g.The("status bar").SetText("right", MakeText(
				g.The("story").Upper(), "by", g.The("story").Text("author"))),
		),
		Before("ending the turn").Always(
			Choose{
				If:   g.The("story").Is("completed"),
				True: g.StopHere(),
			}),
		To("end turn", endTurn),
		After("ending the story").Always(endTurn),
	)

	s.The("stories",
		To("commence",
			updateStatusBar,
			Using{
				Object: g.The("player").Object("whereabouts"),
				Run: g.Go(
					g.The("story").Go("set initial position", g.The("player"), g.TheObject()),
					g.The("story").Go("print the banner"),
					g.The("story").Go("describe the first room", g.TheObject()),
					g.The("story").IsNow("playing"),
				),
				Else: Error{T("The player is nowhere.")},
			}),
	)

	s.The("stories",
		Have("player input", "text"),
		Can("parse player input").And("parsing player input").RequiresNothing())

	s.The("stories", To("end the story",
		g.Say("*** The End ***"),
		Using{
			Object: g.The("story"),
			Run: g.Go(
				g.TheObject().IsNow("completed"),
				Choose{
					If: g.TheObject().Is("scored"),
					True: g.Say(
						"In that game you scored", GetNum{"score"}, "out of a possible", GetNum{"maximum score"}, "in", AddNum{GetNum{"turn count"}, I(1)}, "turns"),
				}),
		}))

	s.The("stories",
		Can("set initial position").
			And("setting initial position").
			RequiresOne("actor").
			AndOne("room"),
		To("set initial position",
			g.The("action.Target").SetObject("whereabouts", g.The("action.Context")),
		))

	s.The("stories",
		Can("describe the first room").
			And("describing the first room").RequiresOnly("room"),
		To("describe the first room",
			g.The("action.Target").Go("report the view"),
		))

	s.The("stories",
		Can("print the banner").
			And("printing the banner").RequiresNothing(),

		To("print the banner",
			Using{
				Object: g.Our("story"),
				Run: g.Go(
					g.Say(GetText{"name"}, "."),
					g.Say(ChooseText{
						If:    IsEmpty{GetText{"headline"}},
						False: GetText{"headline"},
						// FIX: default for headline in class.
						True: MakeText("An interactive fiction"),
					}, "by", GetText{"author"}, "."),
					g.Say(VersionString),
				),
			}))

	pkg.AddScript("Stories", s)

	t := NewScript()
	t.The("story",
		Called("testing"),
		HasText("author", T("me")),
		HasText("headline", T("extra extra")))
	t.The("room",
		Called("the nothing"),
		HasText("printed name", T("the nothing")),
		HasText("description", T("an empty room")),
		When("describing").Always(g.StopHere()),
	)
	t.The("player", Exists(), In("the nothing"))

	pkg.AddTest("Stories",
		test.Setup(t).Try("player location",
			test.Expect(IsObj{Parent(g.The("player")), g.The("nothing")}),
		),
		test.Setup(t).Try("commencing the story",
			test.Expect(IsText{g.The("testing").Text("name"), EqualTo{}, T("testing")}),
			test.Run("commence", g.The("testing")).Match(
				"testing.",
				"extra extra by me.",
				VersionString,
				"the nothing",
				"an empty room",
			),
		))
}

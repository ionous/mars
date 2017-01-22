package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/core/stream"
	"github.com/ionous/mars/lang"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
	. "github.com/ionous/mars/std/script"
)

func init() {
	// FIX: a special AddDebugScript? that only gets activated with special command line parameters?
	pkg.Add("Debugging", debugScript())

	t := NewScript()
	t.The("room",
		Called("somewhere"),
		Is("proper named"),
		Exists(),
	)
	t.The("player", Exists(), In("somewhere"))
	t.The("prop", Called("the key"), Exists())
	t.The("player", Possesses("the key"))

	pkg.AddTest("Debugging",
		test.SetupScript(t).
			Try("parent of",
				test.Parse("parent of player").Match("The player => Somewhere"),
			).
			Try("contents of",
				test.Parse("contents of room").Match("Somewhere => player"),
			).
			Try("ancestors",
				test.Parse("parents of key").Match("The key => The player => Somewhere"),
			),
	)
}

func debugScript() (s Script) {
	s.The("actors",
		Can("debug direct parent").And("debugging direct parent").RequiresOnly("kind"),
		To("debug direct parent",
			Using{
				Object: g.The("action.target"),
				Run:    g.Go(g.Say(g.TheObject().Upper(), "=>", lang.TheUpper{Parent(g.TheObject())})),
			},
		))
	s.The("actors",
		Can("debug ancestors").And("debugging ancestors").RequiresOnly("kind"),
		To("debug ancestors",
			Using{
				Object: g.The("action.target"),
				Run: g.Go(g.Say(g.TheObject().Upper(),
					ForEachObj{
						In:   Ancestors(g.TheObject()),
						Go:   Print("=>", g.TheObject().Upper()),
						Else: Print("has no parents!"),
					})),
			},
		))
	// FIX: sometimes parent of -- matches unexpected objects
	// >parent of automat
	//	hall-automat-door => whereabouts main hallway
	s.The("actors",
		Can("debug contents").And("debugging contents").RequiresOnly("kind"),
		To("debug contents",
			g.Say(g.The("action.Target").Lower(), "=>",
				stream.KeySort{"name", g.The("action.target").ObjectList("contents")}),
		))

	s.The("actors",
		Can("debug room contents").And("debugging room contents").RequiresNothing(),
		To("debug room contents",
			g.Call("debug contents", g.The("action.Source"), g.The("action.source").Object("whereabouts"))))

	s.Understand("parent of {{something}}").As("debug direct parent")
	s.Understand("parents of {{something}}").As("debug ancestors")
	s.Understand("contents of {{something}}").As("debug contents")
	s.Understand("contents of room").As("debug room contents")

	return
}

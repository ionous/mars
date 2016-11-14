package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/core/stream"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	// FIX: a special AddDebugScript? that only gets activated with special command line parameters?
	addScript("Debugging", debugScript())
}

func debugScript() (s Script) {
	s.The("actors",
		Can("debug direct parent").And("debugging direct parent").RequiresOne("object"),
		To("debug direct parent",
			Using{Object: g.The("action.target"),
				Run: g.Say(g.TheObject().Upper(), "=>", Ancestors(g.TheObject())),
			},
		))
	// FIX: sometimes parent of -- matches unexpected objects
	// >parent of automat
	//	hall-automat-door => whereabouts main hallway
	s.The("actors",
		Can("debug contents").And("debugging contents").RequiresOne("object"),
		To("debug contents",
			g.Say("debugging contents of", g.The("action.Target").Lower(),
				stream.KeySort{"name", g.The("action.target").ObjectList("contents")}),
		))

	s.The("actors",
		Can("debug room contents").And("debugging room contents").RequiresNothing(),
		To("debug room contents",
			g.Call("debug contents", g.The("action.Source"), g.The("action.source").Object("whereabouts"))))

	s.Understand("parent of {{something}}").As("debug direct parent")
	s.Understand("contents of {{something}}").As("debug contents")
	s.Understand("contents of room").As("debug room contents")

	return
}

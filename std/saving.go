package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// FIX: future runtime load, incl
// listing and selection of save games.
func init() {
	addScript("Saving",
		// future: maybe a named or enum for require? ( tho an arbitrary string would be  unsafe )
		// maybe an optional require?
		The("kinds", Called("globals"), Exist()),
		The("globals", Called("save-settings"), Exist()),
		The("save-setting", Called("auto-save"), Exist()),
		The("save-setting", Called("normal-save"), Exist()),

		The("actors",
			Can("save via input").And("saving via input").RequiresNothing(),
			To("save via input",
				g.The("actor").Go("save it", g.The("normal-save")),
			),
			Can("autosave via input").And("autosaving via input").RequiresNothing(),
			To("autosave via input",
				g.The("actor").Go("save it", g.The("auto-save")),
			),
			Can("save it").And("saving it").RequiresOnly("save-setting"),
			To("save it",
				SaveGame{
					AutoSave: g.The("save-setting").Equals(g.The("auto-save")),
					Saved:    g.Say("saved", GetText{"@"}),
					Failed:   g.Say("couldnt save", GetText{"@"}),
				}),
		),
		Understand("save").As("save via input"),
		Understand("autosave").As("autosave via input"),
	)
}

package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

//
func init() {
	s := NewScript()

	s.The("globals",
		Called("status bar instances"),
		Have("left", "text"),
		Have("right", "text"))

	s.The("status bar instance",
		Called("status bar"),
		Exists())

	s.The("rooms",
		When("reporting the view").Always(
			g.The("status bar").SetText("left", g.The("room").Upper()),
		))

	pkg.Add("StatusBar", s)
}

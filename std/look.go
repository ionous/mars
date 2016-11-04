package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/g"
)

var Look = []backend.Spec{
	// one visible thing, and requring light
	The("actors",
		Can("look").And("looking").RequiresNothing(),
		// note: reflect to location send the actor as a parameter,
		// but report the view doesnt expect parameters.
		To("look",
			g.Our("actor").Object("whereabouts").Go("report the view"),
		),
	),
	Understand("look|l").As("look"),
}

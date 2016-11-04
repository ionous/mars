package std

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/g"
)

var Impress = []backend.Spec{
	The("actors",
		Can("impress").And("impressing").RequiresNothing(),
		To("impress", g.Say(g.The("actor").Upper(), "is unimpressed."))),
}

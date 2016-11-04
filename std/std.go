package std

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/lang"
	"github.com/ionous/mars/std/compat"
)

// Package std contains the basic objects and actions used for for sashimi-style games. It fulfills a role similar to the Inform7 standard library.
var Package = mars.Package{
	Name: "Std",
	Scripts: mars.Scripts(
		//	Inventory,
		Impress,
		Look,
		LookUnder,
		Reports,
		Types,
		Wearing),
	Tests:    mars.Tests(WearingTest),
	Imports:  mars.Imports(&core.Package, &lang.Package),
	Commands: (*Std)(nil),
}

type Std struct {
	*compat.ScriptRef
	*compat.ScriptRefList
	// give
	*GivePropTo
	// locate
	*Location
	// move: shortcuts
	// parent
	*ChangeParent
	// parents: shortcuts
	// put: shortcuts
	// speaker: shortcuts
	// version: constants
}

package std

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/lang"
	"github.com/ionous/mars/std/compat"
)

// Package std contains the basic objects and actions used for for sashimi-style games. It fulfills a role similar to the Inform7 standard library.
var Std = mars.Package{
	Name: "Std",
	Scripts: mars.Scripts(
		//	Inventory,
		Giving,
		Impress,
		Look,
		LookUnder,
		Reports,
		Types,
		Wearing),
	Tests:    mars.Tests(WearingTest, GivingTest),
	Imports:  mars.Imports(&core.Core, &lang.Lang),
	Commands: (*StdDl)(nil),
}

type StdDl struct {
	*compat.ScriptRef
	*compat.ScriptRefList
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

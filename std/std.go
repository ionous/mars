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
		Actors,
		Desc,
		Giving,
		Impressing,
		Looking,
		LookingUnder,
		Rooms,
		Objects,
		//Supporters,
		Wearing),
	Tests: mars.Tests(
		DescTest,
		// GivingTest,
		// ImpressTest,
		//		LookingTest,
		LookingUnderTest,
	// WearingTest,
	),
	Imports:  mars.Imports(&core.Core, &lang.Lang),
	Commands: (*StdDl)(nil),
}

type StdDl struct {
	*compat.ScriptRef
	*compat.ScriptRefList
	*SaveGame
}

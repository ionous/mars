package std

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/lang"
	"github.com/ionous/mars/std/compat"
	"github.com/ionous/mars/std/script"
)

// Package std contains the basic objects and actions used for for sashimi-style games. It fulfills a role similar to the Inform7 standard library.
func Std() *mars.Package {
	if std == nil {
		std = &mars.Package{
			Name:         "Std",
			Scripts:      pkg.Scripts,
			Tests:        pkg.Tests,
			Dependencies: mars.Dependencies(core.Core(), lang.Lang()),
			Commands:     (*StdCommands)(nil),
		}
	}
	return std
}

var std *mars.Package
var pkg mars.PackageBuilder

type StdCommands struct {
	*compat.ScriptRef
	*compat.ScriptRefList
	*compat.ObjListIsEmpty
	*compat.ObjListContains
	*SaveGame
	*script.GoesToFragment
	*script.InLocation
	*script.SupportsContents
	*script.ContainsContents
	*script.PossessesInventory
	*script.WearsClothing
	*DoorHack
}

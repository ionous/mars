package lang

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
)

// Lang provides common language based string manipulations.
func Lang() *mars.Package {
	if lp == nil {
		lp = &mars.Package{
			Name:         "Lang",
			Scripts:      pkg.Scripts,
			Tests:        pkg.Tests,
			Dependencies: mars.Dependencies(core.Core()),
			Commands:     (*LangCommands)(nil),
		}
	}
	return lp
}

var lp *mars.Package
var pkg mars.PackageBuilder

type LangCommands struct {
	*TheUpper
	*TheLower
	*AnUpper
	*ALower
}

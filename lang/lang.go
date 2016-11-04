package lang

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
)

// Lang provides common language based string manipulations.
var Lang = mars.Package{
	Name:     "Lang",
	Scripts:  mars.Scripts(articles),
	Imports:  mars.Imports(&core.Core),
	Commands: (*LangDL)(nil),
	Tests:    mars.Tests(ArticleTest),
}

type LangDL struct {
	*TheUpper
	*TheLower
	*AnUpper
	*ALower
}

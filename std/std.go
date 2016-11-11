package std

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/lang"
	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/test"
	"github.com/ionous/mars/std/compat"
)

// Package std contains the basic objects and actions used for for sashimi-style games. It fulfills a role similar to the Inform7 standard library.
func Std() *mars.Package {
	if std == nil {
		std = &mars.Package{
			Name:     "Std",
			Scripts:  scripts,
			Tests:    tests,
			Imports:  mars.Imports(core.Core(), lang.Lang()),
			Commands: (*StdDL)(nil),
		}
	}
	return std
}

var std *mars.Package

var scripts mars.SpecList

func addScript(_ string, specs ...backend.Spec) {
	scripts = append(scripts, script.NewScript(specs...))
}

var tests []test.Suite

func addTest(name string, units ...test.Unit) {
	tests = append(tests, test.NewSuite(name, units...))
}

type StdDL struct {
	*compat.ScriptRef
	*compat.ScriptRefList
	*SaveGame
}

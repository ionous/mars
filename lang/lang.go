package lang

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/test"
)

// Lang provides common language based string manipulations.
func Lang() *mars.Package {
	if pkg == nil {
		pkg = &mars.Package{
			Name:     "Lang",
			Scripts:  scripts,
			Tests:    tests,
			Imports:  mars.Imports(core.Core()),
			Commands: (*LangDL)(nil),
		}
	}
	return pkg
}

var pkg *mars.Package

var scripts mars.SpecList

func addScript(_ string, specs ...backend.Spec) {
	scripts = append(scripts, backend.SpecList(specs))
}

var tests []test.Suite

func addTest(name string, units ...test.Unit) {
	tests = append(tests, test.NewSuite(name, units...))
}

type LangDL struct {
	*TheUpper
	*TheLower
	*AnUpper
	*ALower
}

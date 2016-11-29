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
			Name:         "Lang",
			Scripts:      scripts,
			Tests:        tests,
			Dependencies: mars.Dependencies(core.Core()),
			Commands:     (*LangCommands)(nil),
		}
	}
	return pkg
}

var pkg *mars.Package

var scripts backend.SpecList

func addScript(_ string, specs ...backend.Spec) {
	scripts.Specs = append(scripts.Specs, specs...)
}

var tests []test.Suite

func addTest(name string, units ...test.Unit) {
	tests = append(tests, test.NewSuite(name, units...))
}

type LangCommands struct {
	*TheUpper
	*TheLower
	*AnUpper
	*ALower
}

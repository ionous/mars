package mars

import (
	//	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/test"
	S "github.com/ionous/sashimi/source"
)

type Package struct {
	Name string
	// Commands enumerates all commands in the package
	// represented by a nil pointer to a structure of command pointers.
	Commands interface{}
	// Interfaces enumerates all interfaces in the package
	// represented by a nil pointer to a structure of interface objects.
	Interfaces interface{}
	// Scripts contains all declarations for the package.
	Scripts backend.SpecList
	// Test contains all test suites for the package.
	Tests TestList
	// Dependencies contains all package dependencies.
	Dependencies DependencyList
}

type DependencyList []Dependency
type TestList []test.Suite

type Dependency *Package

func Scripts(scripts ...backend.Spec) backend.SpecList {
	return backend.SpecList{scripts}
}

func Tests(tests ...test.Suite) TestList {
	return tests
}

func Dependencies(imports ...Dependency) DependencyList {
	return imports
}

type pkgGen struct {
	rem map[string]bool
	src *S.Statements
}

func (p *Package) Generate(src *S.Statements) (err error) {
	g := pkgGen{make(map[string]bool), src}
	return p.genPackage(g)
}

func (i DependencyList) genDependencies(g pkgGen) (err error) {
	for _, p := range i {
		if !g.rem[p.Name] {
			g.rem[p.Name] = true
			if e := (*Package)(p).genPackage(g); e != nil {
				err = e
				break
			}
		}
	}
	return
}

func (p *Package) genPackage(g pkgGen) (err error) {
	if e := p.Dependencies.genDependencies(g); e != nil {
		err = e
	} else {
		err = p.Scripts.Generate(g.src)
	}
	return err
}

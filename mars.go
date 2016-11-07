package mars

import (
	//	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/test"
	S "github.com/ionous/sashimi/source"
)

type Package struct {
	Name string
	// Commands should be a nil pointer to a structure containing pointers to all of the commands in the package.
	Commands interface{}
	// Scripts contains all declarations for the package.
	Scripts SpecList
	// Test contains all test suites for the package.
	Tests TestList
	// Imports contains all package dependencies.
	Imports ImportList
}

type ImportList []Import
type SpecList []backend.Spec
type TestList []test.Suite

type Import *Package

func Scripts(scripts ...backend.Spec) SpecList {
	return scripts
}

func Tests(tests ...test.Suite) TestList {
	return tests
}

func Imports(imports ...Import) ImportList {
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

func (sl SpecList) Generate(src *S.Statements) (err error) {
	for _, s := range sl {
		if e := s.Generate(src); e != nil {
			err = e
			break
		}
	}
	return
}

func (i ImportList) genImports(g pkgGen) (err error) {
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
	if e := p.Imports.genImports(g); e != nil {
		err = e
	} else {
		err = p.Scripts.Generate(g.src)
	}
	return err
}

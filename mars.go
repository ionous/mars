package mars

import (
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/test"
	S "github.com/ionous/sashimi/source"
)

type ScriptLike interface {
	Declarations() []backend.Declaration
}

type PackageBuilder struct {
	Scripts []backend.Declaration
	Tests   []test.Suite
}

func (pb *PackageBuilder) Add(_ string, s ScriptLike) {
	pb.Scripts = append(pb.Scripts, s.Declarations()...)
}

func (pb *PackageBuilder) AddScript(_ string, decls ...backend.Declaration) {
	pb.Scripts = append(pb.Scripts, decls...)
}

func (pb *PackageBuilder) AddTest(name string, units ...test.Unit) {
	pb.Tests = append(pb.Tests, test.NewSuite(name, units...))
}

type Package struct {
	Name string
	// Commands enumerates all commands in the package
	// represented by a nil pointer to a structure of command pointers.
	Commands interface{}
	// Interfaces enumerates all interfaces in the package
	// represented by a nil pointer to a structure of interface objects.
	Interfaces interface{}
	// Scripts contains all declarations for the package.
	Scripts []backend.Declaration
	// Test contains all test suites for the package.
	Tests []test.Suite
	// Dependencies contains all package dependencies.
	Dependencies []Dependency
}

func Dependencies(imports ...Dependency) []Dependency {
	return imports
}

type Dependency *Package

type pkgGen struct {
	rem map[string]bool
	src *S.Statements
}

func (p *Package) Generate(src *S.Statements) (err error) {
	g := pkgGen{make(map[string]bool), src}
	return p.genPackage(g)
}

func genDependencies(g pkgGen, i []Dependency) (err error) {
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
	if e := genDependencies(g, p.Dependencies); e != nil {
		err = e
	} else {
		for _, b := range p.Scripts {
			if e := b.Generate(g.src); e != nil {
				err = e
				break
			}
		}
	}
	return err
}

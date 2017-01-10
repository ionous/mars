package encode

type Libraries []Library

type Library struct {
	Name string
	// Types contains all package interfaces and commands
	Types TypeBlocks
	//
	Scripts []DataBlock
	Tests   []Suite
	//
	Dependencies []string
}

type LibEncoder struct {
	Packages  PackageMap
	Libraries Libraries
}

func (b *LibEncoder) Build() TypeBlocks {
	return b.types
}

func NewLibEncoder() *LibEncoder {
	return &LibEncoder{gen: make(TypeExists), Packages: make(PackageMap)}
}
func (enc *LibEncoder) AddPackage(p *mars.Package) error {
	if deps, e := enc.Packages.AddPackage(p); e != nil {
		err = e
	} else {
		for _, p := range deps.packages {
			if lib, e := enc.addLibrary(p); e != nil {
				err = e
				break
			} else {
				enc.Libraries = append(enc.Libraries, lib)
			}
		}

	}
}

// x Name string
// 	Commands interface{}
// 	Interfaces interface{}
// 	Scripts SpecList
// x 	Tests TestList
// x 	Dependencies DependencyList
func (enc *LibEncoder) addLibrary(p *mars.Package) (lib Library, err error) {
	lib.Name = p.Name
	if tests, e := addSuites(p); e != nil {
		err = e
	} else {
		lib.Tests = tests
		for _, dep := range p.Dependencies {
			lib.Dependencies = append(lib.Dependencies, dep.Name)
		}
	}
}

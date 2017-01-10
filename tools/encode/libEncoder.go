package encode

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/script/backend"
	"log"
)

type Library struct {
	Name         string      `json:"name"`
	Types        []TypeBlock `json:"types,omitempty"`
	Scripts      []DataBlock `json:"scripts,omitempty"`
	Tests        []Suite     `json:"tests,omitempty"`
	Dependencies []string    `json:"deps,omitempty"`
}

func AddLibraries(ps ...*mars.Package) (ret []Library, err error) {
	pkgs := make(PackageMap)
	if deps, e := pkgs.AddPackages(ps...); e != nil {
		err = e
	} else {
		b := &TypeRecorder{}
		for _, p := range deps {
			log.Println("adding library", p.Name)
			if types, e := b.addTypes(p); e != nil {
				err = e
				break
			} else if tests, e := addSuites(p.Tests); e != nil {
				err = e
				break
			} else if deps, e := addDependencies(p.Dependencies); e != nil {
				err = e
				break
			} else if scripts, e := addScripts(p.Scripts); e != nil {
				err = e
				break
			} else {
				lib := Library{p.Name, types, scripts, tests, deps}
				ret = append(ret, lib)
			}
		}
	}
	return
}

func addDependencies(deps []mars.Dependency) (ret []string, err error) {
	for _, dep := range deps {
		ret = append(ret, dep.Name)
	}
	return
}

func addScripts(scripts []backend.Declaration) (ret []DataBlock, err error) {
	for _, script := range scripts {
		if data, e := Compute(script); e != nil {
			err = e
			break
		} else {
			ret = append(ret, data)
		}
	}
	return
}

package encode

import (
	"github.com/ionous/mars"
	"github.com/ionous/sashimi/util/errutil"
)

type PackageMap map[*mars.Package]bool
type PackageList []*mars.Package

func (m PackageMap) AddPackages(ps ...*mars.Package) (ret PackageList, err error) {
	for _, p := range ps {
		if l, e := m.AddPackage(p); e != nil {
			err = e
			break
		} else {
			ret = append(ret, l...)
		}
	}
	return
}

// AddPackage adds p and its dependencies to the PackageMap, returning a list of all new packages it sees. The list can be empty if the passed package has been seen before.
func (m PackageMap) AddPackage(p *mars.Package) (ret PackageList, err error) {
	if !m[p] {
		m[p] = true
		if e := m.addDependencies(p, &ret); e != nil {
			err = e
		} else {
			ret = append(ret, p)
		}
	}
	return
}

// AddDependencies adds the dependencies of the passed package, without adding the package itself.
func (m PackageMap) AddDependencies(p *mars.Package) (ret PackageList, err error) {
	err = m.addDependencies(p, &ret)
	return
}

func (m PackageMap) addDependencies(p *mars.Package, plist *PackageList) (err error) {
	for _, dep := range p.Dependencies {
		if !m[dep] {
			m[dep] = true
			if e := m.addDependencies(dep, plist); e != nil {
				err = errutil.New("error adding dependency", dep.Name, "because", e)
				break
			} else {
				*plist = append(*plist, dep)
			}
		}
	}
	return
}

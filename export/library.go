package export

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/export/encode"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/sashimi/util/errutil"
)

// note: the order of the library fields matters
// "types", for instance, need to be read before "declarations".
type Library struct {
	Name         string
	Dependencies []string
	Types        []encode.TypeBlock
	Declarations []backend.Declaration
	Tests        []encode.SuiteData
}

func NewLibrary(ctx *encode.Context, p *mars.Package) (ret *Library, err error) {
	if types, e := ctx.AddTypes(p); e != nil {
		err = errutil.New("NewLibrary", e)
	} else if tests, e := encode.RecodeSuites(p.Tests); e != nil {
		err = e
	} else {
		ret = &Library{
			Name:         p.Name,
			Dependencies: dependencyNames(p),
			Types:        types,
			Declarations: p.Scripts,
			Tests:        tests,
		}
	}
	return
}

func (*Library) GetSectionType() SectionType {
	return LibrarySectionType
}

// Encode implements encode.Encoder so that we can store blobs for types and tests.
func (u *Library) Encode() (ret encode.DataBlock, err error) {
	if v, e := encode.RecodeScript(u.Declarations); e != nil {
		err = e
	} else {
		ret.Name = "Library"
		a := encode.ArgMap{
			"Name": u.Name,
		}
		if x := u.Dependencies; len(x) > 0 {
			a["Dependencies"] = x
		}
		if x := u.Types; len(x) > 0 {
			a["Types"] = x
		}
		if x := u.Tests; len(x) > 0 {
			a["Tests"] = x
		}
		if x := v; len(v) > 0 {
			a["Declarations"] = x
		}
		ret.Args = a
	}
	return
}

func dependencyNames(p *mars.Package) (ret []string) {
	for _, dep := range p.Dependencies {
		ret = append(ret, dep.Name)
	}
	return
}

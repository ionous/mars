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
	Dependencies []string              `json:",omitempty"`
	Types        []encode.CommandType  `json:",omitempty"`
	Declarations []backend.Declaration `json:",omitempty"`
	Tests        []encode.SuiteData    `json:",omitempty"`
}

func (*Library) GetSectionType() SectionType {
	return LibrarySectionType
}

func NewLibrary(ctx *encode.Context, p *mars.Package) (ret *Library, err error) {
	if types, e := ctx.AddTypes(p); e != nil {
		err = errutil.New("NewLibrary", e)
	} else {
		ue := encode.NewUniformEncoder(ctx.Types)
		if tests, e := encode.RecodeSuites(ue, p.Tests); e != nil {
			err = e
			// } else if decls, e := encode.RecodeScript(ue, p.Scripts); e != nil {
			// 	err = e
		} else {
			ret = &Library{
				Name:         p.Name,
				Dependencies: dependencyNames(p),
				Types:        types,
				Declarations: p.Scripts,
				Tests:        tests,
			}
		}
	}
	return
}

func dependencyNames(p *mars.Package) (ret []string) {
	for _, dep := range p.Dependencies {
		ret = append(ret, dep.Name)
	}
	return
}

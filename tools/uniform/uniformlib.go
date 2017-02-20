package uniform

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/export"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/sashimi/util/errutil"
)

func NewChapter(name string, decl []backend.Declaration) (ret *export.Chapter, err error) {
	ret = &export.Chapter{
		name,
		decl,
	}
	return
}

func NewLibrary(ctx *Context, p *mars.Package) (ret *export.Library, err error) {
	if types, e := ctx.AddTypes(p); e != nil {
		err = errutil.New("NewLibrary", e)
	} else {
		ue := NewUniformEncoder(ctx.Types)
		if tests, e := MakeUniformSuites(ue, p.Tests); e != nil {
			err = e
		} else {
			ret = &export.Library{
				Name:         p.Name,
				Dependencies: dependencyNames(p),
				Types:        []interface{}{types},
				Declarations: p.Scripts,
				Tests:        []interface{}{tests},
			}
		}
	}
	return
}

// create an array of library sections.
func NewLibraries(ctx *Context, ps ...*mars.Package) (sections []export.StorySection, err error) {
	if packs, e := ctx.AddPackages(ps...); e != nil {
		err = e
	} else {
		for _, p := range packs {
			if l, e := NewLibrary(ctx, p); e != nil {
				err = e
			} else {
				sections = append(sections, l)
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

package export

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/export/encode"
)

// create an array of library sections.
func NewLibraries(ctx *encode.Context, ps ...*mars.Package) (sections []StorySection, err error) {
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

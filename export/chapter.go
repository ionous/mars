package export

import (
	"github.com/ionous/mars/export/encode"
	"github.com/ionous/mars/script/backend"
)

// a single chapter; chapters have no dependencies.
type Chapter struct {
	Name         string
	Declarations []backend.Declaration
}

func NewChapter(ctx *encode.Context, name string, decl []backend.Declaration) (ret *Chapter, err error) {
	ret = &Chapter{
		name,
		decl,
	}
	return
}

func (*Chapter) GetSectionType() SectionType {
	return ChapterSectionType
}

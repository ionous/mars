package export

import (
	"github.com/ionous/mars/script/backend"
)

// a single chapter; chapters have no dependencies.
type Chapter struct {
	Name       string
	Directives []backend.Directive
}

func (*Chapter) GetSectionType() SectionType {
	return ChapterSectionType
}

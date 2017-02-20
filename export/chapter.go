package export

import (
	"github.com/ionous/mars/script/backend"
)

// a single chapter; chapters have no dependencies.
type Chapter struct {
	Name         string
	Declarations []backend.Declaration
}

func (*Chapter) GetSectionType() SectionType {
	return ChapterSectionType
}

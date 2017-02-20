package export

import (
	"github.com/ionous/mars/script/backend"
)

// note: the order of the library fields matters
// "types" needs to be read before "declarations".
type Library struct {
	Name         string
	Dependencies []string              `json:",omitempty"`
	Types        []interface{}         `mars:";TypeBlock" json:",omitempty"`
	Declarations []backend.Declaration `json:",omitempty"`
	Tests        []interface{}         `mars:";SuiteData" json:",omitempty"`
}

func (*Library) GetSectionType() SectionType {
	return LibrarySectionType
}

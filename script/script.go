package script

import (
	"github.com/ionous/mars/script/internal"
	S "github.com/ionous/sashimi/source"
)

type Script []internal.BackendPhrase

func (s Script) BuildStatements() (ret S.Statements, err error) {
	b := S.BuildingBlocks{}
	if e := s.Build(internal.Source{&b}); e != nil {
		err = e
	} else {
		ret = b.Statements()
	}
	return
}

// Build implements BackendPhrase
func (s Script) Build(src internal.Source) (err error) {
	for _, b := range s {
		if e := b.Build(src); e != nil {
			err = e
			break
		}
	}
	return err
}

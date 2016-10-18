package script

import (
	"github.com/ionous/sashimi/source"
)

type Script []BackendPhrase

type Source struct {
	*source.BuildingBlocks
}

type BackendPhrase interface {
	Build(Source) error
}

const Unknown = source.Code("unknown")

func (s Script) Build(src Source) (err error) {
	for _, b := range s {
		if e := b.Build(src); e != nil {
			err = e
			break
		}
	}
	return err
}

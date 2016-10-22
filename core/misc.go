package core

import (
	"github.com/ionous/sashimi/util/ident"
)

func MakeStringId(name string) ident.Id {
	return ident.MakeId(name)
}

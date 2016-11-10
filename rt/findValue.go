package rt

import (
	"github.com/ionous/sashimi/meta"
	"strings"
)

type Scope interface {
	FindValue(string) (meta.Generic, error)
	ScopePath() ScopePath
}

type ScopePath []string

func (sp ScopePath) String() string {
	return strings.Join(sp, "/")
}

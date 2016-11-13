package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

type ValueScope struct {
	val meta.Generic
}

func NewValue(val meta.Generic) rt.Scope {
	return &ValueScope{val}
}

func (s *ValueScope) FindValue(name string) (ret meta.Generic, err error) {
	if name != "@" {
		err = NotFound(s, "value is not an object")
	} else {
		ret = s.val
	}
	return
}

func (s *ValueScope) ScopePath() rt.ScopePath {
	return []string{"value"}
}

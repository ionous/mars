package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

// has to be a struct because instance is interface.
// as an internal object it cant be serialized accidentally.
type ObjectScope struct {
	obj rt.Object
}

func NewObjectScope(obj rt.Object) rt.Scope {
	return &ObjectScope{obj}
}

func (s *ObjectScope) String() string {
	return s.obj.GetId().String()
}

// FindValue implements Scope
func (s *ObjectScope) FindValue(name string) (ret meta.Generic, err error) {
	if p, ok := s.obj.FindProperty(name); !ok {
		err = NotFound(s, name)
	} else {
		ret = p.GetGeneric()
	}
	return
}

func (s *ObjectScope) ScopePath() rt.ScopePath {
	return []string{"object", s.obj.GetId().String()}
}

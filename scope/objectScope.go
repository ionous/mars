package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

// has to be a struct because instance is interface.
// as an internal object it cant be serialized accidentally.
type ObjectScope struct {
	inst meta.Instance // rt.Object?
}

func NewObjectScope(inst meta.Instance) rt.Scope {
	return &ObjectScope{inst}
}

func (s *ObjectScope) String() string {
	return s.inst.GetId().String()
}

// FindValue implements Scope
func (s *ObjectScope) FindValue(name string) (ret meta.Generic, err error) {
	if p, ok := s.inst.FindProperty(name); !ok {
		err = NotFound(s, name)
	} else {
		ret = p.GetGeneric()
	}
	return
}

func (s *ObjectScope) ScopePath() []string {
	return []string{"object", s.inst.GetId().String()}
}

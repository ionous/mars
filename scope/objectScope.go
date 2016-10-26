package scope

import (
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

// has to be a struct because instance is interface.
// as an internal object it cant be serialized accidentally.
type ObjectScope struct {
	meta.Instance
}

func (obj ObjectScope) String() string {
	return obj.GetId().String()
}

// FindValue implements Scope
func (obj ObjectScope) FindValue(name string) (ret meta.Generic, err error) {
	if p, ok := obj.FindProperty(name); !ok {
		err = errutil.New("object property unknown", obj, name)
	} else {
		ret = p.GetGeneric()
	}
	return
}

func (obj ObjectScope) ScopePath() []string {
	return []string{"object", obj.Instance.GetId().String()}
}

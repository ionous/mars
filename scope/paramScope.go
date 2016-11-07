package scope

import (
	"github.com/ionous/sashimi/meta"
	"strings"
)

type ParamScope struct {
	values []meta.Generic
}

func NewParamScope(vs []meta.Generic) ParamScope {
	return ParamScope{vs}
}

func (p ParamScope) FindValue(name string) (ret meta.Generic, err error) {
	if i, ok := p.findByParamName(name); !ok {
		err = NotFound(p, name)
	} else {
		ret = p.values[i]
	}
	return
}
func (p ParamScope) ScopePath() []string {
	return []string{"scope.ParamScope"}
}

// findByParamName: source, target, or context
func (p ParamScope) findByParamName(name string) (ret int, okay bool) {
	for index, src := range []string{"action.Source", "action.Target", "action.Context"} {
		if strings.EqualFold(name, src) {
			ret, okay = index, true
			break
		}
	}
	return
}

package scope

import (
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
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
	} else if i < len(p.values) {
		ret = p.values[i]
	} else {
		err = errutil.New("ParamScope", name, "out of range", i, "of", len(p.values))
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

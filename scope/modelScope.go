package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/lang"
	"github.com/ionous/sashimi/util/sbuf"
)

// .. make sure hint only comes from listener class target: yes.
type ModelScope struct {
	model meta.Model
}

func NewModelScope(m meta.Model) ModelScope {
	return ModelScope{m}
}

func (ms ModelScope) FindValue(name string) (ret meta.Generic, err error) {
	if id := ident.MakeId(lang.StripArticle(name)); id.Empty() {
		err = errutil.New("find value bad id from name", sbuf.Q(name))
	} else if i, ok := ms.model.GetInstance(id); !ok {
		err = NotFound(ms, name)
	} else {
		ret = rt.Object{i}
	}
	return
}

func (sc ModelScope) ScopePath() rt.ScopePath {
	return []string{"scope.ModelScope"}
}

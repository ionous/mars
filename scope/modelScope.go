package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/lang"
)

// .. make sure hint only comes from listener class target: yes.

type ModelScope struct {
	model meta.Model
}

func NewModelScope(m meta.Model) ModelScope {
	return ModelScope{m}
}

func (ms ModelScope) FindValue(name string) (ret meta.Generic, err error) {
	// StripStringId
	if id := ident.MakeId(lang.StripArticle(name)); id.Empty() {
		err = NotNamed(name)
	} else if i, ok := ms.model.GetInstance(id); !ok {
		err = NotFound(ms, name)
	} else {
		ret = rt.Object{i}
	}
	return
}

func (sc ModelScope) ScopePath() []string {
	return []string{"scope.ModelScope"}
}

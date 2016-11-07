package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

// SimilarClass: when all else fails try compatible classes one by one.
func SimilarClass(model meta.Model, nouns meta.Nouns,
	values []meta.Generic) (ret rt.Scope) {
	return &ClassScope{model, &SimilarClassFinder{model, nouns, values}}
}

type SimilarClassFinder struct {
	model  meta.Model
	nouns  meta.Nouns
	values ValueFinder
}

func (cf *SimilarClassFinder) FindClass(id ident.Id) (ret meta.Generic, err error) {
	err = ClassNotFound(string(id))
	for i, nounClass := range cf.nouns {
		if similar := cf.model.AreCompatible(id, nounClass); similar {
			ret, err = cf.values.getValue(i)
			break
		}
	}
	return
}

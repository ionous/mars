package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

// matching logic follows GameEventAdapter.GetObject ( and, RuntimeAction.findByName ) from sashimi v1
func NewListener(model meta.Model, nouns meta.Nouns, hint ident.Id, values []meta.Generic) rt.FindValue {
	return NewChain(
		ModelScope{model},
		ParamScope{values},
		ClassHint(model, hint, values),
		ExactClass(model, nouns, values),
		SimilarClass(model, nouns, values))
}

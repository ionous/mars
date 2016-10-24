package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

// matching logic follows RuntimeAction from sashimi v1
func NewAction(model meta.Model, nouns meta.Nouns, values []meta.Generic) rt.FindValue {
	return NewChain(
		NewModelScope(model),
		NewParamScope(values),
		ExactClass(model, nouns, values),
		SimilarClass(model, nouns, values))
}

package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

// matching logic follows GameEventAdapter.GetObject ( and, RuntimeAction.findByName ) from sashimi v1
func NewListener(model meta.Model, nouns meta.Nouns, values []meta.Generic) *ListenerScope {
	hp := new(ident.Id)
	return &ListenerScope{hp, ScopeChain{
		NewModelScope(model),
		// action.Target, etc.
		NewParamScope(values),
		// when we handle events from callbacks, we set this to the target's class
		ClassHint(model, hp, values),
		// these are the classes originally named in the action declaration; not the sub-classes of the event target. ie. s.The("actors", Can("crawl"), not s.The("babies", When("crawling")
		ExactClass(model, nouns, values),
		SimilarClass(model, nouns, values),
	}}
}

type ListenerScope struct {
	hint  *ident.Id
	chain ScopeChain
}

func (l *ListenerScope) SetHint(id ident.Id) {
	*l.hint = id
}

func (l *ListenerScope) FindValue(s string) (meta.Generic, error) {
	return l.chain.FindValue(s)
}

func (l *ListenerScope) ScopePath() rt.ScopePath {
	return l.chain.ScopePath()
}

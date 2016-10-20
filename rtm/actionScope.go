package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/lang"
	"strings"
)

type ActionScope struct {
	model  meta.Model
	nouns  meta.Nouns
	values []meta.Generic
	chain  rt.Scope
}

func (act ActionScope) FindValue(name string) (ret meta.Generic, err error) {
	// FIX FIX FIX FIX: the hint happens from listenr
	if i, ok := act.findByName(act.model, name, ident.Empty()); !ok {
		ret, err = act.chain.FindValue(name)
	} else {
		if i < len(act.values) {
			ret = act.values[i]
		} else {
			err = errutil.New("out of range", name, i, len(act.values))
		}
	}
	return
}

// findByName:
func (act ActionScope) findByName(m meta.Model, name string, hint ident.Id) (ret int, okay bool) {
	if obj, ok := act.findByParamName(name); ok {
		okay, ret = true, obj
	} else if obj, ok := act.findByClassName(m, name, hint); ok {
		okay, ret = true, obj
	}
	return
}

// findByParamName: source, target, or context
func (act ActionScope) findByParamName(name string) (ret int, okay bool) {
	for index, src := range []string{"action.Source", "action.Target", "action.Context"} {
		if strings.EqualFold(name, src) {
			ret, okay = index, true
			break
		}
	}
	return
}

// findByClassName:
func (act ActionScope) findByClassName(m meta.Model, name string, hint ident.Id) (ret int, okay bool) {
	clsid := ident.MakeId(m.Pluralize(lang.StripArticle(name)))
	if clsid == hint {
		ret, okay = 0, true
	} else {
		ret, okay = act.findByClass(m, clsid)
	}
	return
}

// findByExactClass; true if found
func (act ActionScope) findByClass(m meta.Model, id ident.Id) (ret int, okay bool) {
	// these are the classes originally named in the action declaration; not the sub-classes of the event target. ie. s.The("actors", Can("crawl"), not s.The("babies", When("crawling")
	if obj, ok := act.findByExactClass(m, id); ok {
		ret, okay = obj, true
	} else {
		// when all else fails try compatible classes one by one.
		ret, okay = act.findBySimilarClass(m, id)
	}
	return ret, okay
}

// findByExactClass; true if found
func (act ActionScope) findByExactClass(_ meta.Model, id ident.Id) (ret int, okay bool) {
	for i, nounClass := range act.nouns {
		if same := id == nounClass; same {
			ret, okay = i, true
			break
		}
	}
	return
}

// findBySimilarClass; true if found
func (act ActionScope) findBySimilarClass(m meta.Model, id ident.Id) (ret int, okay bool) {
	for i, nounClass := range act.nouns {
		if similar := m.AreCompatible(id, nounClass); similar {
			ret, okay = i, true
			break
		}
	}
	return
}

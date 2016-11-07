package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

// ExactClass matches the classes named in the action declaration and not the sub-classes of the event target. ie. s.The("actors", Can("crawl"), not s.The("babies", When("crawling")
func ExactClass(model meta.Model, nouns meta.Nouns,
	values []meta.Generic) (ret rt.Scope) {
	return &ClassScope{model, &ExactClassFinder{nouns, values}}
}

type ClassNotFound string

func (e ClassNotFound) Error() string {
	return string(e)
}

type ExactClassFinder struct {
	nouns  meta.Nouns
	values ValueFinder
}

func (cf *ExactClassFinder) FindClass(id ident.Id) (ret meta.Generic, err error) {
	err = ClassNotFound(string(id))
	for i, nounClass := range cf.nouns {
		if same := id == nounClass; same {
			ret, err = cf.values.getValue(i)
			break
		}
	}
	return
}

package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

//FindValue(string) (meta.Generic, error)
type ScopeChain struct {
	scopes []rt.FindValue
}

func NewChain(scopes ...rt.FindValue) rt.FindValue {
	return ScopeChain{scopes}
}

func (sc ScopeChain) FindValue(name string) (ret meta.Generic, err error) {
	err = NotFound(sc, name)
	for _, s := range sc.scopes {
		if v, e := s.FindValue(name); e == nil {
			ret, err = v, nil
			break
		} else if _, notFound := e.(NotFoundError); !notFound {
			err = e
			break
		}
	}
	return
}

func (sc ScopeChain) ScopePath() (parts []string) {
	for _, s := range sc.scopes {
		p := s.ScopePath()
		parts = append(parts, p...)
	}
	return parts
}

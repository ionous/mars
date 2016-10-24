package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

func Make(run rt.Runtime, scope rt.FindValue) rt.Runtime {
	return Provider{run, scope}
}

type Provider struct {
	rt.Runtime
	scope rt.FindValue
}

func (p Provider) FindValue(name string) (meta.Generic, error) {
	return p.scope.FindValue(name)
}
func (p Provider) ScopePath() []string {
	return p.scope.ScopePath()
}

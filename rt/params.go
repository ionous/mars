package rt

import (
	"github.com/ionous/sashimi/meta"
)

type ParameterSource interface {
	// Push should call push target.
	Resolve(Runtime) (meta.Generic, error)
}

type Parameters []ParameterSource

func (ps Parameters) Resolve(run Runtime) (ret []meta.Generic, err error) {
	for _, p := range ps {
		if v, e := p.Resolve(run); e != nil {
			err = e
			break
		} else {
			ret = append(ret, v)
		}
	}
	return
}

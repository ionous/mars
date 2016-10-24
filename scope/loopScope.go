package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"strings"
)

// LoopMaker creates LoopScopes
type LoopMaker struct {
	run rt.Runtime
}

func NewLoopMaker(run rt.Runtime) *LoopMaker {
	return &LoopMaker{run: run}
}

type LoopScope struct {
	rt.Runtime
	index           int
	isFirst, isLast bool
	value           meta.Generic
}

func (sc *LoopMaker) Looper(i int, first, last bool, value meta.Generic) rt.Runtime {
	return LoopScope{sc.run, i, first, last, value}
}

func (sc LoopScope) FindValue(name string) (ret meta.Generic, err error) {
	if name == "" {
		ret = sc.value
	} else if !strings.HasPrefix(name, "@") {
		ret, err = sc.Runtime.FindValue(name)
	} else {
		switch {
		case strings.EqualFold(name, "@first"):
			ret = rt.Bool(sc.isFirst)
		case strings.EqualFold(name, "@last"):
			ret = rt.Bool(sc.isLast)
		case strings.EqualFold(name, "@index"):
			ret = rt.Number(float64(sc.index))
		default:
			err = errutil.New("LoopMaker, unknown field", name)
		}
	}
	return
}

func (sc LoopScope) ScopePath() []string {
	parts := sc.Runtime.ScopePath()
	return append(parts, "loop scope")
}

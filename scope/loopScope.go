package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"strings"
)

// Looper creates LoopScopes
type Looper struct {
	run rt.Runtime
	Stream
	index int
}

type Stream interface {
	HasNext() bool
}

func NewLooper(run rt.Runtime, stream Stream) *Looper {
	return &Looper{run, stream, 0}
}

type LoopScope struct {
	rt.Runtime
	index           int
	isFirst, isLast bool
	value           meta.Generic
}

func (l *Looper) NextScope(value meta.Generic) rt.Runtime {
	first := l.index == 0
	last := !l.HasNext()
	run := &LoopScope{l.run, l.index, first, last, value}
	l.index++
	return run
}

func (l *LoopScope) FindValue(name string) (ret meta.Generic, err error) {
	if name == "" {
		ret = l.value
	} else if !strings.HasPrefix(name, "@") {
		ret, err = l.Runtime.FindValue(name)
	} else {
		switch {
		case strings.EqualFold(name, "@first"):
			ret = rt.Bool(l.isFirst)
		case strings.EqualFold(name, "@last"):
			ret = rt.Bool(l.isLast)
		case strings.EqualFold(name, "@index"):
			ret = rt.Number(float64(l.index))
		default:
			err = errutil.New("Looper, unknown field", name)
		}
	}
	return
}

func (l *LoopScope) ScopePath() []string {
	parts := l.Runtime.ScopePath()
	return append(parts, "loop scope")
}

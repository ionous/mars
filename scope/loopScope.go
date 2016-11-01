package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"strings"
)

// Looper creates LoopScopes
type Looper struct {
	Stream
	index int
}

type Stream interface {
	HasNext() bool
}

func NewLooper(stream Stream) *Looper {
	return &Looper{stream, 0}
}

type LoopScope struct {
	index           int
	isFirst, isLast bool
	value           meta.Generic
}

func (l *Looper) NextScope(value meta.Generic) rt.FindValue {
	first := l.index == 0
	last := !l.HasNext()
	s := &LoopScope{l.index + 1, first, last, value}
	l.index++
	return s
}

func (l *LoopScope) FindValue(name string) (ret meta.Generic, err error) {
	if !strings.HasPrefix(name, "@") {
		err = NotFound(l, name)
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
	return []string{"loop scope"}
}

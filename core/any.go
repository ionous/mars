package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

func Any(any ...rt.BoolEval) rt.BoolEval {
	return AnyTrue{any}
}

type AnyTrue struct {
	Test []rt.BoolEval
}

func (a AnyTrue) GetBool(run rt.Runtime) (okay rt.Bool, err error) {
	prelim := false
	for _, b := range a.Test {
		if ok, e := b.GetBool(run); e != nil {
			err = errutil.Append(err, e)
			// fix: this is a very interesting question
			// guess it depends on whether you want errors to be continuable or not
			// break
		} else if ok {
			prelim = true
			break
		}
	}
	okay = rt.Bool(prelim)
	return
}

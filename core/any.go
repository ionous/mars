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

func (a AnyTrue) GetBool(r rt.Runtime) (okay bool, err error) {
	prelim := false
	for _, b := range a.Test {
		if ok, e := b.GetBool(r); e != nil {
			err = errutil.Append(err, e)
			// fix: this is a very interesting question
			// guess it depends on whether you want errors to be continuable or not
			// break
		} else if ok {
			prelim = true
			break
		}
	}
	okay = prelim
	return
}

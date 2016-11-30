package core

import "github.com/ionous/mars/rt"

func All(all ...rt.BoolEval) rt.BoolEval {
	return AllTrue{all}
}

type AllTrue struct {
	Test []rt.BoolEval
}

func (a AllTrue) GetBool(run rt.Runtime) (okay bool, err error) {
	prelim := true
	for _, b := range a.Test {
		if ok, e := b.GetBool(run); e != nil {
			err = e
			break // see also any.
		} else if !ok {
			prelim = false
			break
		}
	}
	okay = prelim
	return
}

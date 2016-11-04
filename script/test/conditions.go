package test

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

type Conditions []rt.BoolEval

func (cs Conditions) Test(try Trytime) (err error) {
	for i, c := range cs {
		if e := try.Test(c); e != nil {
			err = errutil.New("failed condition", i, e)
			break
		}
	}
	return
}

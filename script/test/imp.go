package test

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

// Imp might become an interface
type Imp struct {
	Input, Match string
	Execute      rt.Execute
}

func (t Imp) Run(try Trytime) (err error) {
	if t.Input != "" && t.Execute != nil {
		err = errutil.New("test implementation has both parser input and raw execute statements specified", t.Input)
	} else {
		var out string
		if t.Execute != nil {
			if res, e := try.Execute(t.Execute); e != nil {
				err = e
			} else {
				out = res
			}
		} else {
			if res, e := try.Parse(t.Input); e != nil {
				err = e
			} else {
				out = res
			}
		}
		// after running
		if err == nil && t.Match != "" && t.Match != out {
			// FIX: add quote to sbuf
			err = errutil.New("expected", "`"+t.Match+"`", "received", "`"+out+"`")
		}
	}
	return
}

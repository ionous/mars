package test

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
	"github.com/stretchr/testify/assert"
	"strings"
)

// Imp might become an interface
type Imp struct {
	Input   string
	Match   []string
	Args    []meta.Generic
	Execute rt.Execute
}

func (t Imp) Run(try Trytime) (err error) {
	if t.Input != "" && t.Execute != nil {
		err = errutil.New("test implementation has both parser input and raw execute statements specified", t.Input)
	} else {
		var out []string
		if t.Execute != nil {
			if res, e := try.Execute(t.Execute); e != nil {
				err = e
			} else {
				out = res
			}
		} else if t.Args != nil {
			if res, e := try.Run(t.Input, t.Args); e != nil {
				err = e
			} else {
				out = res
			}
		} else if t.Input != "" {
			if res, e := try.Parse(t.Input); e != nil {
				err = e
			} else {
				out = res
			}
		}
		e := len(out)
		for e > 0 {
			if len(out[e-1]) > 0 {
				break
			}
			e -= 1
		}
		out = out[:e]
		// after running
		if err == nil && (!assert.ObjectsAreEqualValues(t.Match, out)) {
			if t.Match != nil && (len(t.Match) != 0 || len(out) != 0) {
				// FIX: add quote to sbuf
				err = errutil.New("Imp.Run",
					"expected", sbuf.Q(strings.Join(t.Match, ";")),
					"received", sbuf.Q(strings.Join(out, ";")))
			}
		}
	}
	return
}

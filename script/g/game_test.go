package g

import (
	"bytes"
	c "github.com/ionous/mars/core"
	"github.com/ionous/mars/rtm"
	"github.com/stretchr/testify/assert"
	"testing"
)

//
func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	run := rtm.NewRtm(nil)
	run.PushOutput(&buf)
	x := Say("hello", "there.", "world.")
	if e := x.Execute(run.Runtime()); assert.NoError(t, e, "execute") {
		assert.Equal(t, "hello there. world.\n", buf.String(), "result")
	}
}

//
func TestForEach(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	run := rtm.NewRtm(nil)
	run.PushOutput(buf)
	ts := c.Ts("hello", "there", "world")
	lines := c.ForEachText{
		In:   ts,
		Go:   c.Say(c.GetText{}),
		Else: c.Error{c.T("should have run")},
	}
	if e := lines.Execute(run.Runtime()); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "hello\nthere\nworld\n", buf.String(), "on multiple lines") {
			t.FailNow()
		}
	}
	buf.Reset()

	x := c.PrintLine{c.ForEachText{
		In:   ts,
		Go:   c.PrintText{c.GetText{}},
		Else: c.Error{c.T("should have run")},
	}}

	if e := x.Execute(run.Runtime()); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "hello there world\n", buf.String(), "one one line") {
			t.FailNow()
		}
	}
	buf.Reset()

	index := c.ForEachText{
		In:   ts,
		Go:   Say(c.EachIndex{}),
		Else: c.Error{c.T("should have run")},
	}
	if e := index.Execute(run.Runtime()); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "1\n2\n3\n", buf.String(), "count now") {
			t.FailNow()
		}
	}

	andAlways :=
		c.ForEachText{
			Go: Say(c.ChooseText{
				If:   c.IfEach{IsFirst: true},
				True: c.T("first"),
				False: c.ChooseText{
					If:    c.IfEach{IsLast: true},
					True:  c.T("last"),
					False: c.GetText{},
				}}),
			In:   ts,
			Else: c.Error{c.T("should have run")},
		}
	buf.Reset()
	if e := andAlways.Execute(run.Runtime()); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "first\nthere\nlast\n", buf.String(), "first and last") {
			t.FailNow()
		}
	}
}

package g

import (
	"bytes"
	. "github.com/ionous/mars/core"
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
	if e := x.Execute(run); assert.NoError(t, e, "execute") {
		assert.Equal(t, "hello there. world.\n", buf.String(), "result")
	}
}

//
func TestForEach(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	run := rtm.NewRtm(nil)
	run.PushOutput(buf)
	ts := Ts("hello", "there", "world")
	lines := EachText{
		In:   ts,
		Go:   Say(GetText{}),
		Else: Error{T("should have run")},
	}
	if e := lines.Execute(run); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "hello\nthere\nworld\n", buf.String(), "on multiple lines") {
			t.FailNow()
		}
	}
	buf.Reset()

	x := PrintLine{EachText{
		In:   ts,
		Go:   PrintText{GetText{}},
		Else: Error{T("should have run")},
	}}

	if e := x.Execute(run); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "hello there world\n", buf.String(), "one one line") {
			t.FailNow()
		}
	}
	buf.Reset()

	index := EachText{
		In:   ts,
		Go:   Say(EachIndex{}),
		Else: Error{T("should have run")},
	}
	if e := index.Execute(run); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "1\n2\n3\n", buf.String(), "count now") {
			t.FailNow()
		}
	}

	andAlways :=
		EachText{
			Go: Say(ChooseText{
				If:   IfEach{IsFirst: true},
				True: T("first"),
				False: ChooseText{
					If:    IfEach{IsLast: true},
					True:  T("last"),
					False: GetText{},
				}}),
			In:   ts,
			Else: Error{T("should have run")},
		}
	buf.Reset()
	if e := andAlways.Execute(run); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "first\nthere\nlast\n", buf.String(), "first and last") {
			t.FailNow()
		}
	}
}

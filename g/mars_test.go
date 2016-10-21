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
	r := rtm.NewRtm(nil)
	r.PushOutput(&buf)
	x := Statements{
		Say("hello", "there.", "world."),
	}
	if e := x.Execute(r); assert.NoError(t, e, "execute") {
		assert.Equal(t, "hello there. world.\n", buf.String(), "result")
	}
}

//
func TestForEach(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	r := rtm.NewRtm(nil)
	r.PushOutput(buf)
	ts := Ts("hello", "there", "world")
	lines := EachText{
		Go:   Say(GetText{}),
		For:  ts,
		Else: Error{"should have run"},
	}
	if e := lines.Execute(r); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "hello\nthere\nworld\n", buf.String(), "on multiple lines") {
			t.FailNow()
		}
	}
	buf.Reset()

	x := PrintLine{Statements{EachText{
		Go:   PrintText{GetText{}},
		For:  ts,
		Else: Error{"should have run"},
	}}}

	if e := x.Execute(r); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "hello there world\n", buf.String(), "one one line") {
			t.FailNow()
		}
	}
	buf.Reset()

	index := EachText{
		Go:   Say(EachIndex{}),
		For:  ts,
		Else: Error{"should have run"},
	}
	if e := index.Execute(r); assert.NoError(t, e, "execute") {
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
			For:  ts,
			Else: Error{"should have run"},
		}
	buf.Reset()
	if e := andAlways.Execute(r); assert.NoError(t, e, "execute") {
		if !assert.Equal(t, "first\nthere\nlast\n", buf.String(), "first and last") {
			t.FailNow()
		}
	}
}

package script

import (
	"github.com/ionous/mars/rt"
	S "github.com/ionous/sashimi/source"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleScript(t *testing.T) {
	src := &S.Statements{}
	s := NewScript(
		The("kinds",
			Called("rooms"),
			Have("greeting", "text"),
		),
		The("room", Called("world"), HasText("greeting", rt.Text{"hello"})),
	)
	if e := s.GenerateScript(src); assert.NoError(t, e, "failed to build") {
		assert.Len(t, src.Asserts, 2, "one kind, one instance")
		assert.Len(t, src.Properties, 1, "room has one property")
		assert.Len(t, src.KeyValues, 1, "instance has one value")
	}
}

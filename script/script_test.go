package script

import (
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
		The("room", Called("world"), Has("greeting", "hello")),
	)
	if e := s.Generate(src); assert.NoError(t, e, "failed to build") {
		assert.Len(t, src.Asserts, 2, "one kind, one instance")
		assert.Len(t, src.Properties, 1, "room has one property")
		assert.Len(t, src.KeyValues, 1, "instance has one value")
	}
}

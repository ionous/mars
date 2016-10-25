package script

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleScript(t *testing.T) {
	s := Script{
		The("kinds",
			Called("rooms"),
			Have("greeting", "text"),
		),
		The("room", Called("world"), Has("greeting", "hello")),
	}
	if res, e := s.BuildStatements(); assert.NoError(t, e, "failed to build") {
		assert.Len(t, res.Asserts, 2, "one kind, one instance")
		assert.Len(t, res.Properties, 1, "room has one property")
		assert.Len(t, res.KeyValues, 1, "instance has one value")
	}
}

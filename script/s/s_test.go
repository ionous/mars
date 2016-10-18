package s

import (
	"encoding/xml"
	"github.com/ionous/mars/script"
	"github.com/ionous/sashimi/source"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleRoom(t *testing.T) {
	s := script.Script{
		The("kinds",
			Called("rooms"),
			Have("greeting", "text"),
		),
		The("room", Called("world"), Has("greeting", "hello")),
	}

	if xml, e := xml.MarshalIndent(s, "", " "); assert.NoError(t, e, "serialized") {
		t.Log(string(xml))
	}

	b := source.BuildingBlocks{}
	if e := s.Build(script.Source{&b}); assert.NoError(t, e, "failed to build") {
		res := b.Statements()
		assert.Len(t, res.Asserts, 2, "one kind, one instance")
		assert.Len(t, res.Properties, 1, "room has one property")
		assert.Len(t, res.KeyValues, 1, "instance has one value")
	}
}

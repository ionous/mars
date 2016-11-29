package internal

import (
	. "github.com/ionous/mars/script/backend"
	"github.com/ionous/sashimi/meta"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

func HasPropertyValue(key types.NamedProperty, value meta.Generic) PropertyValue {
	return PropertyValue{key, value}
}

type PropertyValue struct {
	Property types.NamedProperty
	Value    meta.Generic
}

func (f PropertyValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject.String(), f.Property.String(), f.Value}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

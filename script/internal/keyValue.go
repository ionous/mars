package internal

import (
	. "github.com/ionous/mars/script/backend"
	"github.com/ionous/sashimi/meta"
	S "github.com/ionous/sashimi/source"
)

func HasPropertyValue(key string, value meta.Generic) PropertyValue {
	return PropertyValue{key, value}
}

type PropertyValue struct {
	Key   string
	Value meta.Generic
}

func (f PropertyValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{string(top.Subject), f.Key, f.Value}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

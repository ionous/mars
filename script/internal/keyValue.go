package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

func SetKeyValue(key string, value interface{}) KeyValue {
	return KeyValue{key, value}
}

type KeyValue struct {
	key   string
	value interface{}
}

func (f KeyValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject, f.key, f.value}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

package frag

import (
	"github.com/ionous/mars/script"
	S "github.com/ionous/sashimi/source"
)

type SetKeyValue struct {
	Key   string
	Value interface{}
}

func (f SetKeyValue) Build(src script.Source, top Topic) (err error) {
	fields := S.KeyValueFields{top.Subject, f.Key, f.Value}
	return src.NewKeyValue(fields, script.Unknown)
}

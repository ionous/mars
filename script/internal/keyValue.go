package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type NumberValue struct {
	Property string `mars:"has [property]"`
	Number   rt.NumberEval
}

func (f NumberValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject, f.Property, f.Number}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

type TextValue struct {
	Property string `mars:"has [property]"`
	Text     rt.TextEval
}

func (f TextValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject, f.Property, f.Text}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

type RefValue struct {
	Property string `mars:"has [property]"`
	Noun     string
}

func (f RefValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject, f.Property, f.Noun}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

type NumberValues struct {
	Property string `mars:"has [property]"`
	Numbers  rt.NumListEval
}

func (f NumberValues) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject, f.Property, f.Numbers}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

type TextValues struct {
	Property string `mars:"has [property]"`
	Strings  rt.TextListEval
}

func (f TextValues) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject, f.Property, f.Strings}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

// type RefValues struct {
// 	Property string
// 	Values   []string
// }

// func (f RefValues) GenFragment(src *S.Statements, top Topic) error {
// 	fields := S.KeyValueFields{top.Subject, f.Property, f.Values}
// 	return src.NewKeyValue(fields, S.UnknownLocation)
// }

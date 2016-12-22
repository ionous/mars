package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

type NumberValue struct {
	Property types.NamedProperty
	Value    rt.NumberEval
}

func (f NumberValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject.String(), f.Property.String(), f.Value}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

type TextValue struct {
	Property types.NamedProperty
	Value    rt.TextEval
}

func (f TextValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject.String(), f.Property.String(), f.Value}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

type RefValue struct {
	Property types.NamedProperty
	Value    string
}

func (f RefValue) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject.String(), f.Property.String(), f.Value}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

type NumberValues struct {
	Property types.NamedProperty
	Values   rt.NumListEval
}

func (f NumberValues) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject.String(), f.Property.String(), f.Values}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

type TextValues struct {
	Property types.NamedProperty
	Values   rt.TextListEval
}

func (f TextValues) GenFragment(src *S.Statements, top Topic) error {
	fields := S.KeyValueFields{top.Subject.String(), f.Property.String(), f.Values}
	return src.NewKeyValue(fields, S.UnknownLocation)
}

// type RefValues struct {
// 	Property types.NamedProperty
// 	Values   []string
// }

// func (f RefValues) GenFragment(src *S.Statements, top Topic) error {
// 	fields := S.KeyValueFields{top.Subject.String(), f.Property.String(), f.Values}
// 	return src.NewKeyValue(fields, S.UnknownLocation)
// }

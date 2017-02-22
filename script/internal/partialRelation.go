package internal

import (
	. "github.com/ionous/mars/script/backend"
)

// PartialRelation provides golang functions for creating class relation data.
type PartialRelation struct {
	fragment Fragment
	data     *RelationData
}

func NewHaveOne(name string, class string) PartialRelation {
	relative := Relative{name, class}
	r := &HaveOne{Relative: relative}
	return PartialRelation{r, (*RelationData)(r)}
}

func NewHaveMany(name string, class string) PartialRelation {
	relative := Relative{name, class}
	r := &HaveMany{Relative: relative}
	return PartialRelation{r, (*RelationData)(r)}
}

// Implying pivots to allow a reciprocal kind property relation.
func (f PartialRelation) Implying(kind string, rev PartialRelation) Fragment {
	switch rev.fragment.(type) {
	case *HaveOne:
		f.data.Implying = ImplyingOne{kind, rev.data.Relative}
	case *HaveMany:
		f.data.Implying = ImplyingMany{kind, rev.data.Relative}
	}
	return f.fragment
}

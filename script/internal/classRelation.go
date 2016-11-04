package internal

import (
	S "github.com/ionous/sashimi/source"
	"strings"
)

type ClassRelation struct {
	Src          ClassRelative
	ReverseClass string // the reverse subject from implying(reverse)
	Dst          ClassRelative
}

type ClassRelative struct {
	Name string // property,field name
	Kind string // property kind: primitive or user class
	Hint S.RelativeHint
}

// Implying pivots to allow a reciprocal kind property relation.
func (f ClassRelation) Implying(kind string, dst ClassRelation) ClassRelation {
	// NOTE: we can't test that the implied class matches the original have class
	// because the plurals might not match -- we rely on the compiler to detect mismatches
	return ClassRelation{f.Src, kind, dst.Src}
}

func (f ClassRelation) GenFragment(src *S.Statements, top Topic) (err error) {
	// uses the subject, ex. gremlins, and the field, ex. pets: gremlins-pets-relation
	via := strings.Join([]string{top.Subject, f.Src.Name, "relation"}, "-")

	srel := S.RelativeProperty{top.Subject, f.Src.Name, f.Src.Kind, via, f.Src.Hint | S.RelativeSource}
	if e := src.NewRelative(srel, S.UnknownLocation); e != nil {
		err = e
	} else if f.ReverseClass != "" {
		drel := S.RelativeProperty{f.ReverseClass, f.Dst.Name, f.Dst.Kind, via, f.Dst.Hint}
		err = src.NewRelative(drel, S.UnknownLocation)
	}
	return err
}

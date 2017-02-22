package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"strings"
)

type Relative struct {
	Property string `mars:";property"`
	Class    string `mars:"of [class];class"`
}

type RelationData struct {
	Relative
	Implying ReverseRelation
}

type HaveOne RelationData
type HaveMany RelationData

type Relation interface {
	GetRelation() (S.RelativeHint, Relative)
}

type ReverseRelation interface {
	GetReverse() (string, S.RelativeHint, Relative)
}

//
type ImplyingNothing struct{}

func (f ImplyingNothing) GetReverse() (a string, b S.RelativeHint, c Relative) {
	return a, b, c
}

type ReverseRelative struct {
	Kind string `mars:";class"`
	Relative
}

type ImplyingOne ReverseRelative

func (f ImplyingOne) GetReverse() (string, S.RelativeHint, Relative) {
	return f.Kind, S.RelativeOne, f.Relative
}

type ImplyingMany ReverseRelative

func (f ImplyingMany) GetReverse() (string, S.RelativeHint, Relative) {
	return f.Kind, S.RelativeMany, f.Relative
}

func (f PartialRelation) ImplyingNothing() Fragment {
	f.data.Implying = ImplyingNothing{}
	return f.fragment
}

func (f *HaveOne) GetRelation() (S.RelativeHint, Relative) {
	return S.RelativeOne, f.Relative
}

func (f *HaveOne) GenFragment(src *S.Statements, top Topic) (err error) {
	return f.Genifer(src, top, f, f.Implying)
}

func (f *HaveMany) GetRelation() (S.RelativeHint, Relative) {
	return S.RelativeMany, f.Relative
}

func (f *HaveMany) GenFragment(src *S.Statements, top Topic) (err error) {
	return f.Genifer(src, top, f, f.Implying)
}

func (f Relative) Genifer(s *S.Statements, top Topic, this Relation, other ReverseRelation) (err error) {
	// uses the subject, ex. gremlins, and the field, ex. pets: gremlins-pets-relation
	srcHint, srcRel := this.GetRelation()
	srcName, srcTarget := srcRel.Property, srcRel.Class
	srcClass, srcHint := top.Subject, srcHint|S.RelativeSource
	//
	via := strings.Join([]string{srcClass, srcName, "relation"}, "-")
	srel := S.RelativeProperty{srcClass, srcName, srcTarget, via, srcHint}
	if e := s.NewRelative(srel, S.UnknownLocation); e != nil {
		err = e
	} else {
		revClass, revHint, revRel := other.GetReverse()
		if revHint != 0 {
			revName, revTarget := revRel.Property, revRel.Class
			drel := S.RelativeProperty{revClass, revName, revTarget, via, revHint}
			err = s.NewRelative(drel, S.UnknownLocation)
		}
	}
	return err
}

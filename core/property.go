package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

type PropertyName string

func (p PropertyName) String() string {
	return string(p)
}

// Property refers to a field within an object.
type Property struct {
	Ref   rt.RefEval
	Field PropertyName
}

type NumProperty Property
type TextProperty Property
type RefProperty Property
type NumListProperty Property
type TextListProperty Property
type RefListProperty Property

func (p NumProperty) GetNumber(r rt.Runtime) (ret rt.Number, err error) {
	if p, g, e := Property(p).GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.NumEval); !ok {
		err = errutil.New("property not a number", p, sbuf.Type{v})
	} else {
		ret, err = v.GetNumber(r)
	}
	return
}

func (p TextProperty) GetText(r rt.Runtime) (ret rt.Text, err error) {
	if p, g, e := Property(p).GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.TextEval); !ok {
		err = errutil.New("property not text", p, sbuf.Type{v})
	} else {
		ret, err = v.GetText(r)
	}
	return
}

func (p RefProperty) GetReference(r rt.Runtime) (ret rt.Reference, err error) {
	if p, g, e := Property(p).GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.RefEval); !ok {
		err = errutil.New("property not a reference", p, sbuf.Type{v})
	} else {
		ret, err = v.GetReference(r)
	}
	return
}

func (p NumListProperty) GetNumberIdx(r rt.Runtime, i int) (ret rt.Number, err error) {
	if p, g, e := Property(p).GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.NumListEval); !ok {
		err = errutil.New("property not a number list", p, sbuf.Type{v})
	} else {
		ret, err = v.GetNumberIdx(r, i)
	}
	return
}

func (p TextListProperty) GetTextIdx(r rt.Runtime, i int) (ret rt.Text, err error) {
	if p, g, e := Property(p).GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.TextListEval); !ok {
		err = errutil.New("property not a text list", p, sbuf.Type{v})
	} else {
		ret, err = v.GetTextIdx(r, i)
	}
	return
}

func (p RefListProperty) GetReferenceIdx(r rt.Runtime, i int) (ret rt.Reference, err error) {
	if p, g, e := Property(p).GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.RefListEval); !ok {
		err = errutil.New("property not a reference list", p, sbuf.Type{v})
	} else {
		ret, err = v.GetReferenceIdx(r, i)
	}
	return
}

func (p Property) GetGeneric(r rt.Runtime) (retprop meta.Property, retvalue meta.Generic, err error) {
	if ref, e := Property(p).Ref.GetReference(r); e != nil {
		err = e
	} else if obj, e := r.GetObject(ref); e != nil {
		err = e
	} else if prop, ok := obj.FindProperty(Property(p).Field.String()); !ok {
		err = errutil.New("object property not found", obj, p)
	} else {
		retprop, retvalue = prop, prop.GetGeneric()
	}
	return
}

func (p Property) SetGeneric(r rt.Runtime, g meta.Generic) (err error) {
	if ref, e := Property(p).Ref.GetReference(r); e != nil {
		err = e
	} else if obj, e := r.GetObject(ref); e != nil {
		err = e
	} else if prop, ok := obj.FindProperty(Property(p).Field.String()); !ok {
		err = errutil.New("object property not found", obj, Property(p).Field)
	} else {
		err = prop.SetGeneric(g)
	}
	return
}

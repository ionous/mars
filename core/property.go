package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

// P provides access to the named field within the referened object.
func P(ref rt.RefEval, field string) Property {
	return Property{
		Ref: ref, Field: PropertyName(field),
	}
}

type PropertyName string

func (p PropertyName) String() string {
	return string(p)
}

// Property refers to a field within an object.
type Property struct {
	Ref   rt.RefEval
	Field PropertyName
}

func (p Property) GetNumber(r rt.Runtime) (ret rt.Number, err error) {
	if p, g, e := p.GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.NumEval); !ok {
		err = errutil.New("property not a number", p, v)
	} else {
		ret, err = v.GetNumber(r)
	}
	return
}

func (p Property) GetText(r rt.Runtime) (ret rt.Text, err error) {
	if p, g, e := p.GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.TextEval); !ok {
		err = errutil.New("property not text", p, v)
	} else {
		ret, err = v.GetText(r)
	}
	return
}

func (p Property) GetReference(r rt.Runtime) (ret rt.Reference, err error) {
	if p, g, e := p.GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.RefEval); !ok {
		err = errutil.New("property not a reference", p, v)
	} else {
		ret, err = v.GetReference(r)
	}
	return
}

func (p Property) GetNumberIdx(r rt.Runtime, i int) (ret rt.Number, err error) {
	if p, g, e := p.GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.NumListEval); !ok {
		err = errutil.New("property not a number list", p, v)
	} else {
		ret, err = v.GetNumberIdx(r, i)
	}
	return
}

func (p Property) GetTextIdx(r rt.Runtime, i int) (ret rt.Text, err error) {
	if p, g, e := p.GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.TextListEval); !ok {
		err = errutil.New("property not a text list", p, v)
	} else {
		ret, err = v.GetTextIdx(r, i)
	}
	return
}

func (p Property) GetReferenceIdx(r rt.Runtime, i int) (ret rt.Reference, err error) {
	if p, g, e := p.GetGeneric(r); e != nil {
		err = e
	} else if v, ok := g.(rt.RefListEval); !ok {
		err = errutil.New("property not a reference list", p, v)
	} else {
		ret, err = v.GetReferenceIdx(r, i)
	}
	return
}

func (p Property) GetGeneric(r rt.Runtime) (retprop meta.Property, retvalue meta.Generic, err error) {
	if ref, e := p.Ref.GetReference(r); e != nil {
		err = e
	} else if obj, e := r.GetObject(ref); e != nil {
		err = e
	} else if prop, ok := obj.FindProperty(p.Field.String()); !ok {
		err = errutil.New("object property not found", obj, p)
	} else {
		retvalue, retprop = prop.GetGeneric(), prop
	}
	return
}

func (p Property) SetGeneric(r rt.Runtime, g meta.Generic) (err error) {
	if ref, e := p.Ref.GetReference(r); e != nil {
		err = e
	} else if obj, e := r.GetObject(ref); e != nil {
		err = e
	} else if prop, ok := obj.FindProperty(p.Field.String()); !ok {
		err = errutil.New("object property not found", obj, p.Field)
	} else {
		err = prop.SetGeneric(g)
	}
	return
}

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
	Ref   rt.ObjEval
	Field PropertyName
}

type NumProperty Property
type TextProperty Property
type RefProperty Property
type NumListProperty Property
type TextListProperty Property
type RefListProperty Property

func (p NumProperty) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.NumEval); !ok {
		err = errutil.New("property", p, "is not a number", sbuf.Type{v})
	} else {
		ret, err = v.GetNumber(run)
	}
	return
}

func (p TextProperty) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.TextEval); !ok {
		err = errutil.New("property", p, "is not text", sbuf.Type{v})
	} else {
		ret, err = v.GetText(run)
	}
	return
}

func (p RefProperty) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.ObjEval); !ok {
		err = errutil.New("property", p, "is not an object", sbuf.Type{v})
	} else {
		ret, err = v.GetObject(run)
	}
	return
}

func (p NumListProperty) GetNumStream(run rt.Runtime, i int) (ret rt.NumberStream, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.NumListEval); !ok {
		err = errutil.New("property", p, "is not a number list", sbuf.Type{v})
	} else {
		ret, err = v.GetNumStream(run)
	}
	return
}

func (p TextListProperty) GetTextStream(run rt.Runtime) (ret rt.TextStream, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.TextListEval); !ok {
		err = errutil.New("property", p, "is not a text list", sbuf.Type{v})
	} else {
		ret, err = v.GetTextStream(run)
	}
	return
}

func (p RefListProperty) GetObjStream(run rt.Runtime) (ret rt.ObjectStream, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.ObjListEval); !ok {
		err = errutil.New("property", p, "is not a reference list", sbuf.Type{v})
	} else {
		ret, err = v.GetObjStream(run)
	}
	return
}

func (p Property) GetGeneric(run rt.Runtime) (retprop meta.Property, retvalue meta.Generic, err error) {
	if obj, e := p.Ref.GetObject(run); e != nil {
		err = e
	} else if prop, ok := obj.FindProperty(p.Field.String()); !ok {
		err = errutil.New("object property not found", obj, p)
	} else {
		retprop, retvalue = prop, prop.GetGeneric()
	}
	return
}

func (p Property) SetGeneric(run rt.Runtime, g meta.Generic) (err error) {
	if obj, e := p.Ref.GetObject(run); e != nil {
		err = e
	} else if prop, ok := obj.FindProperty(p.Field.String()); !ok {
		err = errutil.New("object property not found", obj, Property(p).Field)
	} else {
		err = prop.SetGeneric(g)
	}
	return
}

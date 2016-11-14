package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

// Property refers to a field within an object.
type Property struct {
	Field string
	Ref   rt.ObjEval
}

type PropertyNum Property
type PropertyText Property
type PropertyRef Property
type PropertySafeRef Property
type PropertyNumList Property
type PropertyTextList Property
type PropertyRefList Property

func (p PropertyNum) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.NumberEval); !ok {
		err = errutil.New("property", p, "is not a number", sbuf.Type{g})
	} else {
		ret, err = v.GetNumber(run)
	}
	return
}

func (p PropertyText) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.TextEval); !ok {
		err = errutil.New("property", p, "is not text", sbuf.Type{g})
	} else {
		ret, err = v.GetText(run)
	}
	return
}

func (p PropertyRef) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.ObjEval); !ok {
		err = errutil.New("property", p, "is not an object", sbuf.Type{g})
	} else {
		ret, err = v.GetObject(run)
	}
	return
}

// GetObject never returns error.
func (p PropertySafeRef) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if o, e := PropertyRef(p).GetObject(run); e == nil {
		ret = o
	}
	return
}

func (p PropertyNumList) GetNumberStream(run rt.Runtime) (ret rt.NumberStream, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.NumListEval); !ok {
		err = errutil.New("property", p, "is not a number list", sbuf.Type{g})
	} else {
		ret, err = v.GetNumberStream(run)
	}
	return
}

func (p PropertyTextList) GetTextStream(run rt.Runtime) (ret rt.TextStream, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.TextListEval); !ok {
		err = errutil.New("property", p, "is not a text list", sbuf.Type{g})
	} else {
		ret, err = v.GetTextStream(run)
	}
	return
}

func (p PropertyRefList) GetObjStream(run rt.Runtime) (ret rt.ObjectStream, err error) {
	if p, g, e := Property(p).GetGeneric(run); e != nil {
		err = e
	} else if v, ok := g.(rt.ObjListEval); !ok {
		err = errutil.New("property", p, "is not a reference list", sbuf.Type{g})
	} else {
		ret, err = v.GetObjStream(run)
	}
	return
}

func (p Property) GetGeneric(run rt.Runtime) (retprop meta.Property, retvalue meta.Generic, err error) {
	if obj, e := p.Ref.GetObject(run); e != nil {
		err = e
	} else if prop, ok := obj.FindProperty(p.Field); !ok {
		err = errutil.New("object property", sbuf.Q(p.Field), "not found in", sbuf.Q(obj))
	} else {
		retprop, retvalue = prop, prop.GetGeneric()
	}
	return
}

func (p Property) SetGeneric(run rt.Runtime, g meta.Generic) (err error) {
	if obj, e := p.Ref.GetObject(run); e != nil {
		err = e
	} else if prop, ok := obj.FindProperty(p.Field); !ok {
		err = errutil.New("object property not found", obj, Property(p).Field)
	} else {
		err = prop.SetGeneric(g)
	}
	return
}

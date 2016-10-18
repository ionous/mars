package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

// P provides access to the named field within the referened object.
func P(ref rt.RefEval, field string) Property {
	return Property{
		Ref: ref, Field: PropertyName(field),
	}
}

type PropertyId ident.Id

func (p PropertyId) Id() ident.Id {
	return ident.Id(p)
}
func (p PropertyId) String() string {
	return string(p)
}

func PropertyName(name string) PropertyId {
	return PropertyId(ident.MakeId(name))
}

// Property refers to a field within an object.
type Property struct {
	Ref   rt.RefEval
	Field PropertyId
}

func (p Property) GetNumber(r rt.Runtime) (ret rt.Number, err error) {
	if metav, e := p.value(r); e != nil {
		err = e
	} else {
		ret = N(metav.GetNum())
	}
	return
}

func (p Property) GetText(r rt.Runtime) (ret rt.Text, err error) {
	if metav, e := p.value(r); e != nil {
		err = e
	} else {
		ret = T(metav.GetText())
	}
	return
}

func (p Property) GetReference(r rt.Runtime) (ret rt.Reference, err error) {
	if metav, e := p.value(r); e != nil {
		err = e
	} else {
		ret = rt.Reference(metav.GetObject())
	}
	return
}

func (p Property) GetNumberIdx(r rt.Runtime, i int) (ret rt.Number, err error) {
	if metavs, e := p.values(r); e != nil {
		err = e
	} else {
		ret, err = NumberValueList{metavs}.GetNumberIdx(r, i)
	}
	return
}

func (p Property) GetTextIdx(r rt.Runtime, i int) (ret rt.Text, err error) {
	if metavs, e := p.values(r); e != nil {
		err = e
	} else {
		ret, err = TextValueList{metavs}.GetTextIdx(r, i)
	}
	return
}

func (p Property) GetReferenceIdx(r rt.Runtime, i int) (ret rt.Reference, err error) {
	if metavs, e := p.values(r); e != nil {
		err = e
	} else {
		ret, err = RefValueList{metavs}.GetReferenceIdx(r, i)
	}
	return
}

func (p Property) value(r rt.Runtime) (ret meta.Value, err error) {
	if metap, e := p.prop(r); e != nil {
		err = e
	} else {
		ret = metap.GetValue()
	}
	return
}

func (p Property) values(r rt.Runtime) (ret meta.Values, err error) {
	if metap, e := p.prop(r); e != nil {
		err = e
	} else {
		ret = metap.GetValues()
	}
	return
}

func (p Property) prop(r rt.Runtime) (ret meta.Property, err error) {
	if ref, e := p.Ref.GetReference(r); e != nil {
		err = e
	} else if obj, e := MakeObject(r, ref); e != nil {
		err = e
	} else if metap, ok := obj.GetProperty(p.Field.Id()); !ok {
		err = errutil.New(obj, "does not have property", p.Field)
	} else {
		ret = metap
	}
	return
}

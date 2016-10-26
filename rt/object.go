package rt

import (
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

// Object provides a dl instance representation. It is not a literal, and cannot be saved.
type Object struct {
	Instance meta.Instance
}

func (obj Object) Empty() bool {
	return obj.Instance == nil
}

// GetObject implements ObjEval for objects; allowing objects to be returned from evals.
func (obj Object) GetObject(Runtime) (Object, error) {
	return obj, nil
}

// String returns the object's ident.Id string.
func (obj Object) String() (ret string) {
	if obj.Instance == nil {
		ret = "<nil object>"
	} else {
		ret = obj.GetId().String()
	}
	return
}

// ensure that object implements the instance interface.
var _ meta.Instance = Object{}

// GetId overrides meta.Instance, returning empty if the underlying instance is null.
func (obj Object) GetId() (ret ident.Id) {
	if obj.Instance != nil {
		ret = obj.Instance.GetId()
	}
	return
}

func (obj Object) GetParentClass() (ret ident.Id) {
	if obj.Instance != nil {
		ret = obj.Instance.GetParentClass()
	}
	return

}
func (obj Object) GetOriginalName() (ret string) {
	if obj.Instance != nil {
		ret = obj.Instance.GetOriginalName()
	}
	return
}

func (obj Object) NumProperty() (ret int) {
	if obj.Instance != nil {
		ret = obj.Instance.NumProperty()
	}
	return
}
func (obj Object) PropertyNum(n int) (ret meta.Property) {
	if obj.Instance != nil {
		ret = obj.Instance.PropertyNum(n)
	}
	return
}

func (obj Object) FindProperty(n string) (ret meta.Property, okay bool) {
	if obj.Instance != nil {
		ret, okay = obj.Instance.FindProperty(n)
	}
	return
}

// GetProperty by the property unique id.
func (obj Object) GetProperty(id ident.Id) (ret meta.Property, okay bool) {
	if obj.Instance != nil {
		ret, okay = obj.Instance.GetProperty(id)
	}
	return
}

// GetPropertyByChoice evalutes all properties to find an enumeration which can store the passed choice
func (obj Object) GetPropertyByChoice(id ident.Id) (ret meta.Property, okay bool) {
	if obj.Instance != nil {
		ret, okay = obj.Instance.GetPropertyByChoice(id)
	}
	return
}

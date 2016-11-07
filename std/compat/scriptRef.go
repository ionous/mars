package compat

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/lang"
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

// ScriptRef provides (some) backwards compatibility with the old GameObject interface:
// Supported: Num, Text, Object, Go,
// Future: Id, Exists, FromClass, ParentRelation, Is, IsNow, Equals, Says,
//  SetNum, ObjectList, Set, SetText.
// Not supported:
//  * ParentRelation() parent IObject, rel string
//  * Get(string) IValue: replace with Num, Text, or Object
//  * List(string) IList: replace with NumList, TextList, or ObjectList
type ScriptRef struct {
	rt.ObjEval
}

func (h ScriptRef) Lower() lang.TheLower {
	return lang.TheLower{h}
}

func (h ScriptRef) Upper() lang.TheUpper {
	return lang.TheUpper{h}
}

// Num returns an rt.NumEval yield the property with the passed name.
func (h ScriptRef) Num(name PropertyName) rt.NumEval {
	return PropertyNum{name, h}
}

// Text returns an rt.TextEval yield the property with the passed name.
// g.The("player").Text("greeting"))
func (h ScriptRef) Text(name PropertyName) rt.TextEval {
	return PropertyText{name, h}
}

// Object yields the object property with the passed name.
// It wraps the property's rt.ObjectEval with a ScriptRef to allow chaining.
func (h ScriptRef) Object(name PropertyName) ScriptRef {
	return ScriptRef{PropertyRef{name, h}}
}

func (h ScriptRef) ObjectList(name PropertyName) ScriptRefList {
	return ScriptRefList{PropertyRefList{name, h}}
}

// Is this object in the passed state?
func (h ScriptRef) Is(state string) rt.BoolEval {
	return IsState{h, state}
}

func (h ScriptRef) Equals(other rt.ObjEval) rt.BoolEval {
	return IsObject{h, other}
}

func (h ScriptRef) Exists() rt.BoolEval {
	return IsValid{h}
}

// g.The("player").Go("test nothing"),
func (h ScriptRef) Go(run string, all ...interface{}) rt.Execute {
	parms := make([]meta.Generic, len(all)+1)
	parms[0] = h
	for i, a := range all {
		var ps meta.Generic
		switch val := a.(type) {
		case int:
			ps = I(val)
		case float64:
			ps = N(val)
		// note, rt.Number implements rt.NumEval. no need for a separate switch
		case rt.NumEval:
			ps = val
		case string:
			ps = T(val)
		// note, rt.Text  implements rt.TextEval. no need for a separate switch
		case rt.TextEval:
			ps = val
		// note, rt.Object implements rt.ObjEval, no need for a separate switch
		case rt.ObjEval:
			ps = val
		default:
			panic("go what?")
		}
		parms[i+1] = ps
	}
	return GoCall{
		Action:     MakeStringId(run),
		Parameters: parms,
	}
}

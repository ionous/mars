package core

import (
	"github.com/ionous/mars/rt"
)

// False returns a Bool which evaluates to true.
func True() rt.Bool {
	return true
}

// False returns a Bool which evaluates to false.
func False() rt.Bool {
	return false
}

// I creates a new number variant from an int const
func I(i int) rt.Number {
	return rt.Number(float64(i))
}

// N creates a new number variant from a float64 const
func N(f float64) rt.Number {
	return rt.Number(f)
}

// Zero returns a Number which evaluates to zero(0).
func Zero() rt.Number {
	return rt.Number(0)
}

// T creates a new text literal
func T(s string) rt.Text {
	return rt.Text(s)
}

// Empty return a Text which evaluates to the empty string.
func Empty() rt.Text {
	return rt.Text("")
}

// Id creates a new reference via MakeStringId
func RawId(s string) rt.Reference {
	id := MakeStringId(s)
	return rt.Reference(id)
}

// Nothing return a Ref which evaluates to the "null" object.
func Nothing() (ret rt.Reference) {
	return ret
}

func Ns(vals ...float64) rt.Numbers {
	return rt.Numbers(vals)
}

func Ts(vals ...string) rt.Texts {
	return rt.Texts(vals)
}

func RawIds(vals ...string) (ret rt.References) {
	for i := 0; i < len(vals); i++ {
		ref := RawId(vals[i])
		ret = append(ret, ref)
	}
	return ret
}

package core

import "github.com/ionous/mars/rt"

// I creates a new number variant from an int const
func I(i int) rt.Number {
	return rt.Number(float64(i))
}

// N creates a new number variant from a float64 const
func N(f float64) rt.Number {
	return rt.Number(f)
}

// T creates a new text literal
func T(s string) rt.Text {
	return rt.Text(s)
}

// R creates a new reference via StringStringId
func R(s string) rt.Reference {
	id := StripStringId(s)
	return rt.Reference(id)
}

func Ns(vals ...float64) rt.Numbers {
	return rt.Numbers(vals)
}

func Ts(vals ...string) rt.Texts {
	return rt.Texts(vals)
}

func Rs(vals ...string) rt.References {
	refs := []rt.Reference{}
	for i := 0; i < len(vals); i++ {
		refs = append(refs, R(vals[i]))
	}
	return rt.References(refs)
}

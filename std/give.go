package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Give(prop string) GivePropTo {
	return GivePropTo{prop: prop}
}

func (give GivePropTo) To(actor string) rt.Execute {
	//added the indirection of "acquire it" so we can transform props after the rules of taking/giving have run
	return GoCall{P(R(actor), "acquire it"), rt.Parameters{
		CallWithRef{R(give.prop)},
	}}
}

type GivePropTo struct {
	prop string
}

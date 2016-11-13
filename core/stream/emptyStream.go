package stream

import (
	"github.com/ionous/mars/rt"
)

type EmptyStream struct {
}

func (EmptyStream) HasNext() bool {
	return false
}

func (EmptyStream) GetNext() (rt.Object, error) {
	var none rt.Object
	return none, rt.StreamExceeded("EmptyStream")
}

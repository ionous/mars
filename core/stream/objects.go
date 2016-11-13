package stream

import (
	"github.com/ionous/mars/rt"
)

// Objects provides an array literal for object ids.
type Objects struct {
	list []rt.Object
	idx  int
}

func (it *Objects) HasNext() bool {
	return it.idx < len(it.list)
}

func (it *Objects) GetNext() (ret rt.Object, err error) {
	if !it.HasNext() {
		err = rt.StreamExceeded("Objects")
	} else {
		ret = it.list[it.idx]
		it.idx++
	}
	return
}

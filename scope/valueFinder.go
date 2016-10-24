package scope

import (
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

type ValueFinder []meta.Generic

func (vf ValueFinder) getValue(i int) (ret meta.Generic, err error) {
	if cnt := len(vf); i < 0 || i >= cnt {
		err = errutil.New("HintFinder, value", i, "out of range", cnt)
	} else {
		ret = vf[i]
	}
	return
}

package internal

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func JoinCallbacks(cb rt.Execute, cbs []rt.Execute) (ret rt.Execute) {
	if len(cbs) == 0 {
		ret = cb
	} else {
		ret = core.ExecuteList(append([]rt.Execute{cb}, cbs...))
	}
	return
}

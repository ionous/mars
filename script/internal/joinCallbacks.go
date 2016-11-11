package internal

import (
	"github.com/ionous/mars/rt"
)

func JoinCallbacks(cb rt.Execute, cbs []rt.Execute) []rt.Execute {
	return append([]rt.Execute{cb}, cbs...)
}

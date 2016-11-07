package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/util/errutil"
)

// ScopeStack implements FindValue by reflecting calls to the top of a stack of FindValue objects.
type ScopeStack struct {
	scp []rt.Scope
}

func (os *ScopeStack) Top() rt.Scope {
	return os.scp[0]
}

func (os *ScopeStack) PushScope(args ...rt.Scope) {
	os.scp = append(os.scp, scope.ScopeChain(args))
}

func (os *ScopeStack) PopScope() {
	if l := len(os.scp) - 1; l > 0 {
		os.scp = os.scp[:l]
	} else {
		panic(errutil.New("cant pop last scope"))
	}
}

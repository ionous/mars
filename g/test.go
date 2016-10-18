package g

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Test(b rt.BoolEval, message string) rt.Execute {
	return Choose{If: b,
		True:  Say(message),
		False: Error{message}}
}

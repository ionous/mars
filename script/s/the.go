package s

import (
	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/frag"
)

func The(target string, fragments ...frag.Fragment) script.BackendPhrase {
	return frag.TheOldStyle{
		target, fragments,
	}
}

var Our = The

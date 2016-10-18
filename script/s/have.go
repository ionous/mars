package s

import (
	"github.com/ionous/mars/script/frag"
)

// Have adds a property to all instances of a class.
// For relations, use HaveOne or HaveMany.
func Have(name string, kind string) frag.Fragment {
	return frag.ClassProperty{name, kind}
}

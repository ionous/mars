package g

import (
	. "github.com/ionous/mars/core"
	rt "github.com/ionous/mars/rt"
)

// Goshortcut runs a bunch of statements
func Go(all ...rt.Execute) rt.Execute {
	return Statements(all)
}

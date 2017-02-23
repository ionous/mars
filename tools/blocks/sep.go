package blocks

func BlockEndFilter(prim string) string {
	return prim + `.`
}
func QuoteFilter(prim string) string {
	return `"` + prim + `"`
}

type Separator interface {
	Sep(*State) string
}

type CommaSep struct {
	FullStop bool
}

func (f CommaSep) Sep(n *State) (ret string) {
	if n.end > 0 {
		if fini := (n.pos + 1) == n.end; !fini {
			ret = ", and"
		} else {
			if f.FullStop {
				ret = "."
			} else {
				ret = " "
			}
		}
	}
	return ret
}

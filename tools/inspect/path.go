package inspect

import (
	"strings"
)

type Path []string

func NewPath(p string) Path {
	return strings.Split(p, "/")
}

func (p Path) Empty() bool {
	return len(p) == 0
}

func (p Path) String() (ret string) {
	if len(p) > 0 {
		ret = strings.Join(p, "/")
	} else {
		ret = "(empty)"
	}
	return
}

func PathCompare(a, b Path) (ret int) {
	if d := len(b) - len(a); d != 0 {
		ret = d
	} else {
		ret = strings.Compare(a.String(), b.String())
	}
	return
}

func (p Path) ChildPath(s string) Path {
	// if you append to a slice, and it does not grow
	// the value returned by append doesnt change.
	// not sure why they dont allocate a new slice --
	// im sure its for efficency reasons, but it feel unexpected:
	// string addition always retuns a new string --
	// you dont magically alter the item on the right side of the expression.
	cnt := len(p)
	c := make([]string, cnt, cnt+1)
	copy(c, p)
	return append(c, s)
}

func (p Path) ParentPath() Path {
	cnt := len(p)
	return p[:cnt-1]
}

func (p Path) Last() string {
	return p[len(p)-1]
}

func (p Path) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

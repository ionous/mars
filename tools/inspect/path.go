package inspect

import (
	"strings"
)

type Path []string

func NewPath(p string) Path {
	return strings.Split(p, "/")
}

func (p Path) String() (ret string) {
	if len(p) > 0 {
		ret = strings.Join(p, "/")
	} else {
		ret = "(empty)"
	}
	return
}

func (p Path) ChildPath(s string) Path {
	return append(p, s)
}

func (p Path) Last() string {
	return p[len(p)-1]
}

func (p Path) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/sbuf"
	"strings"
)

func NotFound(s rt.Scope, n string) error {
	return NotFoundError{s, n}
}

type NotFoundError struct {
	scope rt.Scope
	name  string
}

func (nf NotFoundError) Error() string {
	str := strings.Join(nf.scope.ScopePath(), "/")
	return sbuf.New("NotFoundError", sbuf.Q(nf.name), str).Join(" ")
}

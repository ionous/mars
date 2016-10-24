package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/sbuf"
	"strings"
)

type NotNamed string

func (a NotNamed) Error() string {
	return sbuf.New("not named", string(a)).Join(" ")
}

func NotFound(s rt.FindValue, n string) error {
	return NotFoundError{s, n}
}

type NotFoundError struct {
	scope rt.FindValue
	name  string
}

func (nf NotFoundError) Error() string {
	str := strings.Join(nf.scope.ScopePath(), "/")
	return sbuf.New("not found", sbuf.Quote{nf.name}, str).Join(" ")
}

package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type KnownAs struct {
	Name string `mars:"is known as [name]"`
}

//
func (f KnownAs) GenFragment(src *S.Statements, top Topic) (err error) {
	alias := S.AliasFields{top.Subject, []string{f.Name}}
	return src.NewAlias(alias, S.UnknownLocation)
}

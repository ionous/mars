package internal

import (
	S "github.com/ionous/sashimi/source"
)

type KnownAs struct {
	Names []string
}

// Add additional aliases for the current subject.
func (f KnownAs) And(name string) KnownAs {
	f.Names = append(f.Names, name)
	return f
}

func (f KnownAs) BuildFragment(src Source, top Topic) (err error) {
	alias := S.AliasFields{top.Subject, f.Names}
	return src.NewAlias(alias, UnknownLocation)
}

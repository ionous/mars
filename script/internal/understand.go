package internal

import (
	S "github.com/ionous/sashimi/source"
)

type ParserPartial []string

func (p ParserPartial) And(s string) ParserPartial {
	return append(p, s)
}

// MARS: its eems this would read better with the event name than the action name...
func (p ParserPartial) As(s string) BackendPhrase {
	return ParserPhrase{s, p}
}

type ParserPhrase struct {
	What    string
	Phrases ParserPartial
}

func (p ParserPhrase) Build(src Source) error {
	alias := S.AliasFields{p.What, p.Phrases}
	return src.NewAlias(alias, UnknownLocation)
}

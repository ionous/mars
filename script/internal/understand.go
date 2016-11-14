package internal

import (
	"github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type ParserPartial []string

func (p ParserPartial) And(s string) ParserPartial {
	return append(p, s)
}

// MARS: its eems this would read better with the event name than the action name...
func (p ParserPartial) As(s string) backend.Spec {
	return ParserPhrase{s, p}
}

type ParserPhrase struct {
	What    string
	Phrases []string
}

func (p ParserPhrase) Generate(src *S.Statements) error {
	alias := S.AliasFields{p.What, p.Phrases}
	return src.NewAlias(alias, S.UnknownLocation)
}

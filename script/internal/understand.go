package internal

import (
	"github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

type ParserPartial types.PlayerInput

func (p ParserPartial) And(s string) ParserPartial {
	return append(p, s)
}

// MARS: its eems this would read better with the event name than the action name...
func (p ParserPartial) As(s types.NamedAction) backend.Spec {
	return ParserPhrase{types.PlayerInput(p), s}
}

type ParserPhrase struct {
	Input types.PlayerInput `mars:"understand [text]"`
	What  types.NamedAction `mars:"as [action]"`
}

func (p ParserPhrase) Generate(src *S.Statements) error {
	alias := S.AliasFields{p.What.String(), p.Input.Strings()}
	return src.NewAlias(alias, S.UnknownLocation)
}

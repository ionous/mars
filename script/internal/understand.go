package internal

import (
	"github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type ParserPartial []ParserInput

func (p ParserPartial) And(s string) ParserPartial {
	return append(p, MatchString{s})
}

func (p ParserPartial) As(s string) backend.Directive {
	return ParserDirective{[]ParserInput(p), s}
}

type ParserDirective struct {
	Input []ParserInput `mars:"Understand [input];input"`
	What  string        `mars:"as [event];event"`
}

type ParserInput interface {
	GetInputString() string
}

type MatchString struct {
	Input string
}

func (m MatchString) GetInputString() string {
	return m.Input
}

func (p ParserDirective) Generate(src *S.Statements) error {
	var phrases []string
	for _, in := range p.Input {
		phrases = append(phrases, in.GetInputString())
	}
	alias := S.AliasFields{p.What, phrases}
	return src.NewAlias(alias, S.UnknownLocation)
}

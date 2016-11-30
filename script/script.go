package script

import (
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/internal"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

type Script struct {
	Name  string
	Specs []backend.Spec
}

func NewScript(specs ...backend.Spec) Script {
	return Script{Specs: specs}
}

func (s *Script) Add(specs ...backend.Spec) *Script {
	s.Specs = append(s.Specs, specs...)
	return s
}

func (s *Script) The(target string, frags ...backend.Fragment) *Script {
	s.Specs = append(s.Specs, The(target, frags...))
	return s
}

func (s *Script) Understand(input ...string) (ret *ParserIt) {
	ret = &ParserIt{Input: input}
	s.Specs = append(s.Specs, ret)
	return
}

// Generate implements Spec
func (s Script) Generate(src *S.Statements) (err error) {
	for _, b := range s.Specs {
		if e := b.Generate(src); e != nil {
			err = e
			break
		}
	}
	return err
}

//
type ParserIt internal.ParserPhrase

func (p *ParserIt) And(input string) *ParserIt {
	p.Input = append(p.Input, input)
	return p
}

func (p *ParserIt) As(what types.NamedAction) {
	p.What = what
}

func (p *ParserIt) Generate(src *S.Statements) error {
	alias := S.AliasFields{p.What.String(), p.Input.Strings()}
	return src.NewAlias(alias, S.UnknownLocation)
}

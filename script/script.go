package script

import (
	"github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type Script struct {
	specs backend.SpecList
}

func NewScript(specs ...backend.Spec) Script {
	return Script{specs}
}

func (s *Script) Add(specs ...backend.Spec) *Script {
	s.specs = append(s.specs, specs...)
	return s
}

func (s *Script) The(target string, frags ...backend.Fragment) *Script {
	s.specs = append(s.specs, The(target, frags...))
	return s
}

func (s *Script) Understand(input ...string) (ret *ParserIt) {
	ret = &ParserIt{Input: input}
	s.specs = append(s.specs, ret)
	return
}

// Generate implements Spec
func (s Script) Generate(src *S.Statements) (err error) {
	return s.specs.Generate(src)
}

//
type ParserIt struct {
	What  string
	Input []string
}

func (p *ParserIt) And(input string) *ParserIt {
	p.Input = append(p.Input, input)
	return p
}

func (p *ParserIt) As(what string) {
	p.What = what
}

func (p *ParserIt) Generate(src *S.Statements) error {
	alias := S.AliasFields{p.What, p.Input}
	return src.NewAlias(alias, S.UnknownLocation)
}

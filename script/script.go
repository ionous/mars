package script

import (
	"github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type LocalScript struct {
	specs backend.Script
}

func NewScript(specs ...backend.Spec) LocalScript {
	return LocalScript{specs}
}

func (s *LocalScript) Add(specs ...backend.Spec) *LocalScript {
	s.specs = append(s.specs, specs...)
	return s
}

func (s *LocalScript) The(target string, frags ...backend.Fragment) *LocalScript {
	s.specs = append(s.specs, The(target, frags...))
	return s
}

func (s *LocalScript) Understand(input ...string) (ret *ParserIt) {
	ret = &ParserIt{Input: input}
	s.specs = append(s.specs, ret)
	return
}

// Generate implements Spec
func (s LocalScript) Generate(src *S.Statements) (err error) {
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

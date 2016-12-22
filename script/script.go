package script

import (
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/internal"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

type Script struct {
	Name  types.NamedScript
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

func (s *Script) Understand(input ...string) ParserHelper {
	p := ParserHelper{&internal.ParserPhrase{Input: input}}
	s.Specs = append(s.Specs, p.ptr)
	return p
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

// ParserHelper customize a parser phrase created via Script.Understand.
type ParserHelper struct {
	ptr *internal.ParserPhrase
}

func (p ParserHelper) And(input string) ParserHelper {
	p.ptr.Input = append(p.ptr.Input, input)
	return p
}

func (p ParserHelper) As(what types.NamedAction) {
	p.ptr.What = what
}

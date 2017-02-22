package script

import (
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/internal"
	S "github.com/ionous/sashimi/source"
)

type Script struct {
	specs []backend.Directive
}

func NewScript(specs ...backend.Directive) Script {
	return Script{specs: specs}
}

func (s *Script) Add(specs ...backend.Directive) *Script {
	s.specs = append(s.specs, specs...)
	return s
}

func (s *Script) The(target string, frags ...interface{}) *Script {
	s.specs = append(s.specs, The(target, frags...))
	return s
}

func (s *Script) Understand(input string) ParserHelper {
	inputs := []internal.ParserInput{internal.MatchString{input}}
	p := &internal.ParserDirective{Input: inputs}
	h := ParserHelper{p}
	s.specs = append(s.specs, h.ptr)
	return h
}

func (s Script) Directives() []backend.Directive {
	return s.specs
}

// // Generate implements Directive for script.
// // ( but it might be better if parents did this more manually. )
func (s Script) GenerateScript(src *S.Statements) (err error) {
	for _, b := range s.specs {
		if e := b.Generate(src); e != nil {
			err = e
			break
		}
	}
	return err
}

// ParserHelper customize a parser phrase created via Script.Understand.
type ParserHelper struct {
	ptr *internal.ParserDirective
}

func (p ParserHelper) And(input string) ParserHelper {
	p.ptr.Input = append(p.ptr.Input, internal.MatchString{input})
	return p
}

func (p ParserHelper) As(what string) {
	p.ptr.What = what
}

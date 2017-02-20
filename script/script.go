package script

import (
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/internal"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

type Script struct {
	specs []backend.Declaration
}

func NewScript(specs ...backend.Declaration) Script {
	return Script{specs: specs}
}

func (s *Script) Add(specs ...backend.Declaration) *Script {
	s.specs = append(s.specs, specs...)
	return s
}

func (s *Script) The(target string, frags ...interface{}) *Script {
	s.specs = append(s.specs, The(target, frags...))
	return s
}

func (s *Script) Understand(input ...string) ParserHelper {
	p := ParserHelper{&internal.ParserPhrase{Input: input}}
	s.specs = append(s.specs, p.ptr)
	return p
}

func (s Script) Declarations() []backend.Declaration {
	return s.specs
}

// // Generate implements Declaration for script.
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
	ptr *internal.ParserPhrase
}

func (p ParserHelper) And(input string) ParserHelper {
	p.ptr.Input = append(p.ptr.Input, input)
	return p
}

func (p ParserHelper) As(what types.NamedAction) {
	p.ptr.What = what
}

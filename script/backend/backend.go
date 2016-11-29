package backend

import S "github.com/ionous/sashimi/source"

//  Specs are used to generate script into source.
type Spec interface {
	Generate(*S.Statements) error
}

type SpecList struct {
	Specs []Spec
}

// Generate implements Spec
func (s SpecList) Generate(src *S.Statements) (err error) {
	for _, b := range s.Specs {
		if e := b.Generate(src); e != nil {
			err = e
			break
		}
	}
	return err
}

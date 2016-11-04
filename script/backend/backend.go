package backend

import S "github.com/ionous/sashimi/source"

//  Specs are used to generate script into source.
type Spec interface {
	Generate(*S.Statements) error
}

type Script []Spec

// Generate implements Spec
func (s Script) Generate(src *S.Statements) (err error) {
	for _, b := range s {
		if e := b.Generate(src); e != nil {
			err = e
			break
		}
	}
	return err
}

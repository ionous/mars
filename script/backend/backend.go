package backend

import S "github.com/ionous/sashimi/source"

//  Specs are used to generate script into source.
type Spec interface {
	Generate(*S.Statements) error
}

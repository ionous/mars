package backend

import S "github.com/ionous/sashimi/source"

//  Specs are used to generate script into source.
type Declaration interface {
	Generate(*S.Statements) error
}

package backend

import S "github.com/ionous/sashimi/source"

// Directive gets used to generate script into source.
type Directive interface {
	Generate(*S.Statements) error
}
